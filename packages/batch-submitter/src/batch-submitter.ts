/* External Imports */
import { BigNumber, Signer } from 'ethers'
import { BlockWithTransactions, Provider, TransactionResponse } from '@ethersproject/abstract-provider'
import { getContractInterface } from '@eth-optimism/contracts'

/* Internal Imports */
import {
    CanonicalTransactionChainContract,
    EIP155TxData,
    CreateEOATxData,
    TxType,
    ctcCoder,
    Signature,
    Address,
    encodeAppendSequencerBatch,
    Uint8,
    Uint24,
    remove0x,
    BatchContext,
    AppendSequencerBatchParams
} from './batch-encoder'

interface L2Block {
    stateRoot: string
    isSequencerTx: boolean
    sequencerTxType: undefined | TxType
    txData: undefined | EIP155TxData | CreateEOATxData
    timestamp: number
    blockNumber: number
}

interface ExtendedBlock extends BlockWithTransactions {
    stateRoot: string
    txType: number
}

type Batch = L2Block[]

const MAX_TX_SIZE = 100_000

export class BatchSubmitter {
    txChain: CanonicalTransactionChainContract
    signer: Signer
    l2Provider: Provider
    blockCache: {
        [blockNumber: number]: L2Block
    } = {}

    constructor(canonicalTransactionChainAddress: Address, signer: Signer, l2Provider: Provider) {
        this.txChain = new CanonicalTransactionChainContract(
          canonicalTransactionChainAddress,
          getContractInterface('OVM_CanonicalTransactionChain'),
          signer
        )
        this.signer = signer
        this.l2Provider = l2Provider
    }

    async submitNextBatch():Promise<void> {
        const startBlock = parseInt((await this.txChain.getTotalElements()), 16) + 1
        const endBlock = Math.min(startBlock + 100, await this.l2Provider.getBlockNumber())

        const batchParams = await this._generateSequencerBatchParams(startBlock, endBlock)
        const txRes = await this.txChain.appendSequencerBatch(batchParams)
        console.log(txRes)
    }

    async _generateSequencerBatchParams(startBlock: number, endBlock: number):Promise<AppendSequencerBatchParams> {
        // Get all L2Blocks between the given range
        const blocks: Batch = []
        for(let i = startBlock; i < endBlock; i++) {
            if (!this.blockCache.hasOwnProperty(i)) {
                this.blockCache[i] = await this._getL2Block(i)
            }
            blocks.push(this.blockCache[i])
        }
        let sequencerBatchParams = await this._getSequencerBatchParams(startBlock, blocks)
        let encoded = encodeAppendSequencerBatch(sequencerBatchParams)
        while (encoded.length / 2 > MAX_TX_SIZE) {
            blocks.splice(Math.ceil(blocks.length * 2 / 3)) // Delete 1/3rd of all of the blocks
            sequencerBatchParams = await this._getSequencerBatchParams(startBlock, blocks)
        }
        return sequencerBatchParams
    }

    async _getSequencerBatchParams(shouldStartAtIndex: number, blocks: Batch): Promise<AppendSequencerBatchParams> {
        const totalElementsToAppend = blocks.length

        // Generate contexts
        const contexts: BatchContext[] = []
        let lastBlockIsSequencerTx = false
        const groupedBlocks: { sequenced: L2Block[], queued: L2Block[] }[] = []
        for (const block of blocks) {
            if ((lastBlockIsSequencerTx === false && block.isSequencerTx === true) || groupedBlocks.length === 0) {
                groupedBlocks.push({
                    sequenced: [],
                    queued: [],
                })
            }
            const cur = groupedBlocks.length - 1;
            (block.isSequencerTx) ? groupedBlocks[cur].sequenced.push(block) : groupedBlocks[cur].queued.push(block)
            lastBlockIsSequencerTx = block.isSequencerTx
        }
        for (const groupedBlock of groupedBlocks) {
            contexts.push({
                numSequencedTransactions: groupedBlock.sequenced.length,
                numSubsequentQueueTransactions: groupedBlock.queued.length,
                timestamp: (groupedBlock.sequenced.length > 0) ? groupedBlock.sequenced[0].timestamp: 0,
                blockNumber: (groupedBlock.sequenced.length > 0) ? groupedBlock.sequenced[0].blockNumber : 0,
            })
        }

        // Generate sequencer transactions
        const transactions: string[] = []
        for (const block of blocks) {
            if (!block.isSequencerTx) {
                continue
            }
            let encoding: string = ctcCoder.eip155TxData.encode(block.txData as EIP155TxData)
            if (block.sequencerTxType === TxType.EIP155) {
                encoding = ctcCoder.eip155TxData.encode(block.txData as EIP155TxData)
            } else if (block.sequencerTxType === TxType.createEOA) {
                encoding = ctcCoder.createEOATxData.encode(block.txData as CreateEOATxData)
            }
            transactions.push(encoding)
        }

        return {
            shouldStartAtBatch: shouldStartAtIndex - 1,
            totalElementsToAppend,
            contexts,
            transactions
        }
    }

    async _getL2Block(blockNumber: number): Promise<L2Block> {
        const block = await this.l2Provider.getBlockWithTransactions(blockNumber) as ExtendedBlock
        // TODO: Update this query to return the block with the state root!
        block.stateRoot = '0x' + '98'.repeat(32)
        // TODO: Actually pull the tx type from the tx
        block.txType = TxType.EIP155

        if (this._isSequencerTx(block)) {
            if (block.txType === TxType.EIP155) {
                return this._getEIP155L2Block(block)
            } else if (block.txType === TxType.createEOA) {
                return this._getCreateEOAL2Block(block)
            } else {
                throw new Error('Unsupported Tx Type!')
            }
        } else {
            return {
                stateRoot: block.stateRoot,
                isSequencerTx: false,
                sequencerTxType: undefined,
                txData: undefined,
                timestamp: block.timestamp,
                blockNumber: block.number
            }
        }
    }

    private _getEIP155L2Block(block: ExtendedBlock): L2Block {
        const tx: TransactionResponse = block.transactions[0]
        const txData: EIP155TxData = {
            sig: {
                // TODO: Update v value to strip the chainID
                // v: remove0x(BigNumber.from(tx.v).toHexString()).padStart(2, '0'),
                v: '01',
                r: tx.r,
                s: tx.s
            },
            gasLimit: tx.gasLimit.toNumber(),
            gasPrice: tx.gasPrice.toNumber(),
            nonce: tx.nonce,
            target: (tx.to) ? tx.to : '00'.repeat(20),
            data: tx.data,
        }
        return {
            stateRoot: block.stateRoot,
            isSequencerTx: true,
            sequencerTxType: block.txType,
            txData,
            timestamp: block.timestamp,
            blockNumber: block.number
        }
    }

    private _getCreateEOAL2Block(block: ExtendedBlock): L2Block {
        const tx: TransactionResponse = block.transactions[0]
        const txData: CreateEOATxData = {
            sig: {
                // TODO: Update v value to strip the chainID
                v: remove0x(BigNumber.from(tx.v).toHexString()).padStart(2, '0'),
                r: tx.r,
                s: tx.s
            },
            messageHash: tx.data
        }
        return {
            stateRoot: block.stateRoot,
            isSequencerTx: true,
            sequencerTxType: block.txType,
            txData,
            timestamp: block.timestamp,
            blockNumber: block.number
        }
    }

    _isSequencerTx(block: ExtendedBlock): boolean {
        // TODO: Actually check if it's a sequencer tx.
        return true
    }
}
