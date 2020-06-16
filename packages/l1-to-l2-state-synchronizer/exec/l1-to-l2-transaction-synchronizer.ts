/* External Imports */
import {
  BaseDB,
  DB,
  EthereumBlockProcessor,
  getLevelInstance,
  newInMemoryDB,
} from '@eth-optimism/core-db'
import { add0x, getLogger, logError } from '@eth-optimism/core-utils'
import {
  Environment,
  initializeL1Node,
  L1NodeContext,
  CHAIN_ID,
  L1TransactionBatchProcessor,
} from '@eth-optimism/rollup-core'

import { JsonRpcProvider, Provider } from 'ethers/providers'
import * as fs from 'fs'
import * as rimraf from 'rimraf'
import { Wallet } from 'ethers'
import { getWallets } from 'ethereum-waffle'

const log = getLogger('l1-to-l2-transaction-batch-processor')

export const runTest = async (
  l1Provider?: Provider,
  l2Provider?: JsonRpcProvider
): Promise<L1TransactionBatchProcessor> => {
  return run(true, l1Provider, l2Provider)
}

export const run = async (
  testMode: boolean = false,
  l1Provider?: Provider,
  l2Provider?: JsonRpcProvider
): Promise<L1TransactionBatchProcessor> => {
  initializeDBPaths(testMode)

  let l1NodeContext: L1NodeContext
  log.info(`Attempting to connect to L1 Node.`)
  try {
    l1NodeContext = await initializeL1Node(true, l1Provider)
  } catch (e) {
    logError(log, 'Error connecting to L1 Node', e)
    throw e
  }

  let provider: JsonRpcProvider = l2Provider
  if (!provider && !!Environment.l2NodeWeb3Url()) {
    log.info(`Connecting to L2 web3 URL: ${Environment.l2NodeWeb3Url()}`)
    provider = new JsonRpcProvider(Environment.l2NodeWeb3Url(), CHAIN_ID)
  }

  return getL1ToL2TransactionBatchProcessor(testMode, l1NodeContext, provider)
}

/**
 * Gets an L1ToL2TransactionSynchronizer based on configuration and the provided arguments.
 *
 * @param testMode Whether or not this is running as a test
 * @param l1NodeContext The L1 node context.
 * @param l2Provider The L2 JSON RPC Provider to use to communicate with the L2 node.
 * @returns The L1ToL2TransactionSynchronizer.
 */
const getL1ToL2TransactionBatchProcessor = async (
  testMode: boolean,
  l1NodeContext: L1NodeContext,
  l2Provider: JsonRpcProvider
): Promise<L1TransactionBatchProcessor> => {
  const db: DB = getDB(testMode)

  const stateSynchronizer = await L1TransactionBatchProcessor.create(
    db,
    l1NodeContext.provider,
    [], // TODO: fill this in
    [] // TODO: Fill this in
  )

  const earliestBlock = Environment.l1EarliestBlock()

  const blockProcessor = new EthereumBlockProcessor(
    db,
    earliestBlock,
    Environment.stateSynchronizerNumConfirmsRequired()
  )
  await blockProcessor.subscribe(
    l1NodeContext.provider,
    stateSynchronizer,
    true
  )

  return stateSynchronizer
}

/**
 * Gets the appropriate db for this node to use based on whether or not this is run in test mode.
 *
 * @param isTestMode Whether or not it is test mode.
 * @returns The constructed DB instance.
 */
const getDB = (isTestMode: boolean = false): DB => {
  if (isTestMode) {
    return newInMemoryDB()
  } else {
    if (!Environment.stateSynchronizerPersistentDbPath()) {
      log.error(
        `No L1_TO_L2_STATE_SYNCHRONIZER_PERSISTENT_DB_PATH environment variable present. Please set one!`
      )
      process.exit(1)
    }

    return new BaseDB(
      getLevelInstance(Environment.stateSynchronizerPersistentDbPath())
    )
  }
}

/**
 * Gets the wallet to use to interact with the L2 node. This may be configured via
 * private key file specified through environment variables. If not it is assumed
 * that a local test provider is being used, from which the wallet may be fetched.
 *
 * @param provider The provider with which the wallet will be associated.
 * @returns The wallet to use with the L2 node.
 */
const getL2Wallet = (provider: JsonRpcProvider): Wallet => {
  let wallet: Wallet
  if (!!Environment.stateSynchronizerPrivateKey()) {
    wallet = new Wallet(
      add0x(Environment.stateSynchronizerPrivateKey()),
      provider
    )
    log.info(
      `Initialized State Synchronizer wallet from private key. Address: ${wallet.address}`
    )
  } else {
    wallet = getWallets(provider)[0]
    log.info(
      `Getting wallet from provider. First wallet private key: [${wallet.privateKey}`
    )
  }

  if (!wallet) {
    const msg: string = `Wallet not created! Specify the L1_TO_L2_STATE_SYNCHRONIZER_PRIVATE_KEY environment variable to set one!`
    log.error(msg)
    throw Error(msg)
  } else {
    log.info(`State Synchronizer wallet created. Address: ${wallet.address}`)
  }

  return wallet
}

/**
 * Initializes filesystem DB paths. This will also purge all data if the `CLEAR_DATA_KEY` has changed.
 */
const initializeDBPaths = (isTestMode: boolean) => {
  if (isTestMode) {
    return
  }

  if (!fs.existsSync(Environment.l2RpcServerPersistentDbPath())) {
    makeDataDirectory()
  } else {
    if (Environment.clearDataKey() && !fs.existsSync(getClearDataFilePath())) {
      log.info(`Detected change in CLEAR_DATA_KEY. Purging data...`)
      rimraf.sync(`${Environment.stateSynchronizerPersistentDbPath()}/{*,.*}`)
      log.info(
        `L2 RPC Server data purged from '${Environment.stateSynchronizerPersistentDbPath()}/{*,.*}'`
      )
      makeDataDirectory()
    }
  }
}

/**
 * Makes the data directory for this full node and adds a clear data key file if it is configured to use one.
 */
const makeDataDirectory = () => {
  fs.mkdirSync(Environment.stateSynchronizerPersistentDbPath(), {
    recursive: true,
  })
  if (Environment.clearDataKey()) {
    fs.writeFileSync(getClearDataFilePath(), '')
  }
}

const getClearDataFilePath = () => {
  return `${Environment.stateSynchronizerPersistentDbPath()}/.clear_data_key_${Environment.clearDataKey()}`
}

if (typeof require !== 'undefined' && require.main === module) {
  run()
}
