/* External Imports */
import { getLogger } from '@eth-optimism/core-utils'
import { exit } from 'process'
import { getContractFactory } from '@eth-optimism/contracts'
import { getContractInterface } from '@eth-optimism/contracts'
import { Signer, Wallet } from 'ethers'
import { Provider, JsonRpcProvider } from '@ethersproject/providers'

/* Internal Imports */
import { BatchSubmitter, CanonicalTransactionChainContract } from '..'

/* Logger */
const log = getLogger('oe:batch-submitter:init')

interface RequiredEnvVars {
    ADDRESS_MANAGER_ADDRESS: 'ADDRESS_MANAGER_ADDRESS'
    SEQUENCER_PRIVATE_KEY: 'SEQUENCER_PRIVATE_KEY'
    INFURA_NETWORK: 'INFURA_NETWORK'
    INFURA_PROJECT_ID: 'INFURA_PROJECT_ID'
    L2_WEB3_URL: 'L2_WEB3_URL'
    L2_CHAIN_ID: 'L2_CHAIN_ID'
    MAX_TX_SIZE: 'MAX_TX_SIZE'
    POLL_INTERVAL: 'POLL_INTERVAL'
}
const requiredEnvVars: RequiredEnvVars = {
    ADDRESS_MANAGER_ADDRESS: 'ADDRESS_MANAGER_ADDRESS',
    SEQUENCER_PRIVATE_KEY: 'SEQUENCER_PRIVATE_KEY',
    INFURA_NETWORK: 'INFURA_NETWORK',
    INFURA_PROJECT_ID: 'INFURA_PROJECT_ID',
    L2_WEB3_URL: 'L2_WEB3_URL',
    L2_CHAIN_ID: 'L2_CHAIN_ID',
    MAX_TX_SIZE: 'MAX_TX_SIZE',
    POLL_INTERVAL: 'POLL_INTERVAL',
}

export const run = async () => {
    log.info('Starting batch submitter...')

    for (const val in requiredEnvVars) {
        if (!process.env[val]) {
            log.error(`Missing enviornment variable: ${val}`)
            exit(1)
        }
    }
    Object.assign(requiredEnvVars, process.env)

    const l1Provider: Provider = new JsonRpcProvider(
        `https://${requiredEnvVars.INFURA_NETWORK}.infura.io/v3/${requiredEnvVars.INFURA_PROJECT_ID}`
    )
    const l2Provider: Provider = new JsonRpcProvider(requiredEnvVars.L2_WEB3_URL)
    log.info(requiredEnvVars.SEQUENCER_PRIVATE_KEY)
    const sequencerSigner: Signer = new Wallet(requiredEnvVars.SEQUENCER_PRIVATE_KEY, l1Provider)

    const Factory__OVM_CanonicalTransactionChain = await getContractFactory(
      'OVM_CanonicalTransactionChain',
      sequencerSigner
    )

    const unwrapped_OVM_CanonicalTransactionChain = await Factory__OVM_CanonicalTransactionChain.attach(
        requiredEnvVars.ADDRESS_MANAGER_ADDRESS
    )
    const OVM_CanonicalTransactionChain = new CanonicalTransactionChainContract(
      unwrapped_OVM_CanonicalTransactionChain.address,
      getContractInterface('OVM_CanonicalTransactionChain'),
      sequencerSigner
    )

    const batchSubmitter = new BatchSubmitter(
      OVM_CanonicalTransactionChain,
      sequencerSigner,
      l2Provider,
      parseInt(requiredEnvVars.L2_CHAIN_ID),
      parseInt(requiredEnvVars.MAX_TX_SIZE)
    )

    // Run batch submitter!
    while (true) {
        await batchSubmitter.submitNextBatch()
        // Sleep
        await new Promise(r => setTimeout(r, parseInt(requiredEnvVars.POLL_INTERVAL)));
    }
}