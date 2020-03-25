/* External Imports */
import { L2ToL1MessageReceiverContractDefinition } from '@eth-optimism/ovm'

import { Contract, ethers, providers, Wallet } from 'ethers'
import { createMockProvider, deployContract, getWallets } from 'ethereum-waffle'

/* Internal Imports */
import { DEFAULT_ETHNODE_GAS_LIMIT, Environment } from '../index'
import { L1NodeContext } from '../../types'

/**
 * Starts a local node on the provided port, using the provided mnemonic to
 * deploy the necessary contracts for bootstrapping.
 *
 * @param sequencerMnemonic The mnemonic to use for the Sequencer in contracts that need Sequencer ownership or reference.
 * @param port The port the node should be reachable at.
 * @returns The L1 node context with all info necessary to use the L1 node.
 */
export const startLocalL1Node = async (
  mnemonic: string,
  port: number
): Promise<L1NodeContext> => {
  const opts = {
    gasLimit: DEFAULT_ETHNODE_GAS_LIMIT,
    allowUnlimitedContractSize: true,
    locked: false,
    port,
    mnemonic,
  }
  if (!!Environment.localL1NodePersistentDbPath()) {
    opts['db_path'] = Environment.localL1NodePersistentDbPath()
  }
  const provider: providers.Web3Provider = createMockProvider(opts)

  const sequencerWallet = getWallets(provider)[0]
  const l2ToL1MessageReceiver = await deployL2ToL1MessageReceiver(
    sequencerWallet
  )

  return {
    provider,
    sequencerWallet,
    l2ToL1MessageReceiver,
  }
}

export const deployL2ToL1MessageReceiver = async (
  wallet: Wallet
): Promise<Contract> => {
  const contract = await deployContract(
    wallet,
    L2ToL1MessageReceiverContractDefinition,
    [wallet.address, Environment.l2ToL1MessageFinalityDelayInBlocks()]
  )

  process.env.L2_TO_L1_MESSAGE_RECEIVER_ADDRESS = contract.address
  return contract
}
