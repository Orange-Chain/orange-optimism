/* Imports: External */
import * as path from 'path'
import * as dotenv from 'dotenv'
import * as assert from 'assert'
import { ethers } from 'ethers'
import { KECCAK256_RLP_S, KECCAK256_NULL_S } from 'ethereumjs-util'

/* Imports: Uniswap */
import {
  computePoolAddress,
  POOL_INIT_CODE_HASH,
  POOL_INIT_CODE_HASH_OPTIMISM,
  POOL_INIT_CODE_HASH_OPTIMISM_KOVAN,
} from '@uniswap/v3-sdk'
import { Token } from '@uniswap/sdk-core'
import { abi as UNISWAP_FACTORY_ABI } from '@uniswap/v3-core/artifacts/contracts/UniswapV3Factory.sol/UniswapV3Factory.json'

/* Imports: Internal */
import { StateDump, ChainState, PoolData } from './types'
import {
  reqenv,
  readDumpFile,
  toHex32,
  getMappingKey,
  transferStorageSlot,
} from './utils'

const main = async () => {
  // Load required enviorment variables
  dotenv.config()
  const STATE_DUMP_FILE = reqenv('REGEN__STATE_DUMP_FILE')
  const L2_PROVIDER_URL = reqenv('REGEN__L2_PROVIDER_URL')
  const L2_NETWORK_NAME = reqenv('REGEN__L2_NETWORK_NAME')
  const TESTNET_PROVIDER_URL = reqenv('REGEN__TESTNET_PROVIDER_URL')
  const TESTNET_PRIVATE_KEY = reqenv('REGEN__TESTNET_PRIVATE_KEY')
  const UNISWAP_FACTORY_ADDRESS = reqenv('REGEN__UNISWAP_FACTORY_ADDRESS')
  const UNISWAP_NFPM_ADDRESS = reqenv('REGEN__UNISWAP_NFPM_ADDRESS')

  // Input assertions
  assert.ok(
    ['mainnet', 'kovan'].includes(L2_NETWORK_NAME),
    `L2_NETWORK_NAME must be one of "mainnet" or "kovan"`
  )

  // Load the state dump from the JSON file.
  const dump: StateDump = await readDumpFile(
    path.resolve(__dirname, `../dumps/${STATE_DUMP_FILE}`)
  )

  // Store a list of all addresses.
  const allAddresses = Object.keys(dump.accounts)

  // Set up the L2 provider.
  const l2Provider = new ethers.providers.JsonRpcProvider(L2_PROVIDER_URL)

  // Set up the testnet wallet.
  const wallet = new ethers.Wallet(
    TESTNET_PRIVATE_KEY,
    new ethers.providers.JsonRpcProvider(TESTNET_PROVIDER_URL)
  )

  // Create an empty object that represents the new genesis state
  // We're going to move items from the dump into this genesis state
  const genesis: ChainState = {}

  // Sanity check to guarantee that all addresses in dump.accounts are lower case.
  // TODO: do more sanity checking
  console.log(`verifying that all contract addresses are lower case`)
  for (const [address, account] of Object.entries(dump.accounts)) {
    assert.equal(
      address.toLowerCase(),
      address,
      `unexpected upper case character in state dump address: ${address}`
    )

    // All accounts other than precompiles should have a balance of zero.
    if (!address.startsWith('0x00000000000000000000000000000000000000')) {
      assert.equal(
        account.balance,
        '0',
        `unexpected non-zero balance in state dump address: ${address}`
      )
    }
  }

  // Step 1. Transfer the state of each precompiled contract.
  console.log(`moving all precompile contract states to new genesis`)
  for (const [address, account] of Object.entries(dump.accounts)) {
    if (address.startsWith('0x00000000000000000000000000000000000000')) {
      genesis[address] = account
      delete dump.accounts[address]
    }
  }

  // Step 2. Transfer over each EOA address and turn it into a normal EOA.
  // TODO: Verify these are the correct and only EOA code hashes.
  console.log(`removing code from all EOA addresses`)
  const EOA_CODE_HASHES = [
    '0xa73df79c90ba2496f3440188807022bed5c7e2e826b596d22bcb4e127378835a',
    '0xef2ab076db773ffc554c9f287134123439a5228e92f5b3194a28fec0a0afafe3',
  ]
  for (const [address, account] of Object.entries(dump.accounts)) {
    if (EOA_CODE_HASHES.includes(account.codeHash.toLowerCase())) {
      genesis[address] = {
        balance: account.balance,
        nonce: account.nonce,
        root: KECCAK256_RLP_S,
        codeHash: KECCAK256_NULL_S,
      }
      delete dump.accounts[address]
    }
  }

  /* --- BEGIN UNISWAP SURGERY SECTION --- */

  // Set up the uniswap factory contract reference
  const UniswapV3Factory = new ethers.Contract(
    UNISWAP_FACTORY_ADDRESS,
    UNISWAP_FACTORY_ABI,
    l2Provider
  )

  // TODO: use compiler output to confirm the list of variables with different storage slots

  // Step X. (UNISWAP) Fix the UniswapV3Factory `owner` address.
  console.log(`fixing UniswapV3Factory owner address`)
  transferStorageSlot({
    dump,
    address: UNISWAP_FACTORY_ADDRESS,
    oldSlot: toHex32(0),
    newSlot: toHex32(3),
  })

  // Step X. (UNISWAP) Fix the UniswapV3Factory `feeAmountTickSpacing` mapping.
  console.log(`fixing UniswapV3Factory feeAmountTickSpacing mapping`)
  for (const fee of [500, 3000, 10000]) {
    transferStorageSlot({
      dump,
      address: UNISWAP_FACTORY_ADDRESS,
      oldSlot: getMappingKey([fee], 1),
      newSlot: getMappingKey([fee], 4),
    })
  }

  // Step X. (UNISWAP) Figure out the old and new pool addresses.
  console.log(`finding all UniswapV3Factory pool addresses`)
  const pools: PoolData[] = []
  const poolEvents = await UniswapV3Factory.queryFilter('PoolCreated' as any)
  for (const event of poolEvents) {
    // Compute the old pool address using the OVM init code hash.
    const oldPoolAddress = computePoolAddress({
      factoryAddress: UNISWAP_FACTORY_ADDRESS,
      tokenA: new Token(0, event.args.token0, 18),
      tokenB: new Token(0, event.args.token1, 18),
      fee: event.args.fee,
      initCodeHashManualOverride:
        L2_NETWORK_NAME === 'mainnet'
          ? POOL_INIT_CODE_HASH_OPTIMISM
          : POOL_INIT_CODE_HASH_OPTIMISM_KOVAN,
    }).toLowerCase()

    // Compute the new pool address using the EVM init code hash.
    const newPoolAddress = computePoolAddress({
      factoryAddress: UNISWAP_FACTORY_ADDRESS,
      tokenA: new Token(0, event.args.token0, 18),
      tokenB: new Token(0, event.args.token1, 18),
      fee: event.args.fee,
      initCodeHashManualOverride: POOL_INIT_CODE_HASH,
    }).toLowerCase()

    if (oldPoolAddress in dump.accounts) {
      pools.push({
        oldAddress: oldPoolAddress,
        newAddress: newPoolAddress,
        token0: event.args.token0,
        token1: event.args.token1,
        fee: event.args.fee,
      })
    } else {
      // throw new Error(
      //   `found pool event but contract not in state: ${oldPoolAddress}`
      // )
      console.log(
        `found pool event but contract not in state: ${oldPoolAddress}`
      )
    }
  }

  // Step X. (UNISWAP) Fix the UniswapV3Factory `getPool` mapping.
  console.log(`fixing UniswapV3Factory getPool mapping`)
  for (const pool of pools) {
    // Fix the token0 => token1 => fee mapping
    transferStorageSlot({
      dump,
      address: UNISWAP_FACTORY_ADDRESS,
      oldSlot: getMappingKey([pool.token0, pool.token1, pool.fee], 2),
      newSlot: getMappingKey([pool.token0, pool.token1, pool.fee], 5),
      newValue: pool.newAddress,
    })

    // Fix the token1 => token0 => fee mapping
    transferStorageSlot({
      dump,
      address: UNISWAP_FACTORY_ADDRESS,
      oldSlot: getMappingKey([pool.token1, pool.token0, pool.fee], 2),
      newSlot: getMappingKey([pool.token1, pool.token0, pool.fee], 5),
      newValue: pool.newAddress,
    })
  }

  // Step X. (UNISWAP) Fix the NonfungiblePositionManager `_poolIds` mapping.
  console.log(`fixing NonfungiblePositionManager _poolIds mapping`)
  for (const pool of pools) {
    try {
      transferStorageSlot({
        dump,
        address: UNISWAP_NFPM_ADDRESS,
        oldSlot: getMappingKey([pool.oldAddress], 10),
        newSlot: getMappingKey([pool.newAddress], 10),
      })
    } catch (err) {
      if (err.message.includes('old slot not found in state dump')) {
        // It's OK for this to happen because some pools may not have any position NFTs.
        console.log(
          `pool not found in NonfungiblePositionManager _poolIds mapping: ${pool.oldAddress}`
        )
      } else {
        throw err
      }
    }
  }

  // Step X. (UNISWAP) Check for any references to the pool addresses in storage *values*.
  console.log(`checking for any references to pool addresses in storage values`)
  for (const pool of pools) {
    for (const [address, account] of Object.entries(dump.accounts)) {
      if (account.storage === undefined) {
        continue
      }

      // Check for any references to the pool address in storage values.
      for (const [slotKey, slotValue] of Object.entries(account.storage)) {
        if (slotValue.includes(pool.oldAddress.slice(2))) {
          // TODO: Figure out what to do here.
          throw new Error(`found unexpected reference to pool address`)
        }
      }
    }
  }

  // Step X. (UNISWAP) Fix every balance mapping where the pool has a balance.
  console.log(`checking for balance mappings where the pool is referenced`)
  for (const pool of pools) {
    for (const [address, account] of Object.entries(dump.accounts)) {
      if (account.storage === undefined) {
        continue
      }

      // TODO: Choose an appropriate ceiling for the storage slots here
      // Check for single-level nested keys (i.e., address => xxxx).
      for (let i = 0; i < 1000; i++) {
        const oldSlotKey = getMappingKey([pool.oldAddress], i)
        if (account.storage[oldSlotKey] !== undefined) {
          console.log(
            `fixing single-level mapping in contract`,
            `address=${address}`,
            `pool=${pool.oldAddress}`,
            `slot=${oldSlotKey}`
          )
          transferStorageSlot({
            dump,
            address,
            oldSlot: oldSlotKey,
            newSlot: getMappingKey([pool.newAddress], i),
          })
        }
      }
    }
  }

  // Step X. (UNISWAP) Fix every allowance mapping where the pool is referenced.
  console.log(`checking for allowance mappings where the pool is referenced`)
  for (const pool of pools) {
    for (const [address, account] of Object.entries(dump.accounts)) {
      if (account.storage === undefined) {
        continue
      }

      // Check for double-level nested keys (i.e., address => address => xxxx).
      for (let i = 0; i < 1000; i++) {
        for (const otherAddress of allAddresses) {
          // otherAddress => poolAddress => xxxx
          const oldSlotKey1 = getMappingKey([otherAddress, pool.oldAddress], i)
          if (account.storage[oldSlotKey1] !== undefined) {
            console.log(
              `fixing double-level mapping in contract (other => pool => xxxx)`,
              `address=${address}`,
              `pool=${pool.oldAddress}`,
              `slot=${oldSlotKey1}`
            )
            transferStorageSlot({
              dump,
              address,
              oldSlot: oldSlotKey1,
              newSlot: getMappingKey([otherAddress, pool.newAddress], i),
            })
          }

          // poolAddress => otherAddress => xxxx
          const oldSlotKey2 = getMappingKey([pool.oldAddress, otherAddress], i)
          if (account.storage[oldSlotKey2] !== undefined) {
            console.log(
              `fixing double-level mapping in contract (pool => other => xxxx)`,
              `address=${address}`,
              `pool=${pool.oldAddress}`,
              `slot=${oldSlotKey2}`
            )
            transferStorageSlot({
              dump,
              address,
              oldSlot: oldSlotKey2,
              newSlot: getMappingKey([pool.newAddress, otherAddress], i),
            })
          }
        }
      }
    }
  }

  // Step X. (UNISWAP) Deploy pool code for all pools that don't already have code.
  console.log('deploying pool code')
  for (const pool of pools) {
    const poolCode = await wallet.provider.getCode(pool.newAddress)
    if (poolCode === '0x') {
      console.log(`pool has no code, deploying it: ${pool.newAddress}`)
      const poolCreationTx = await UniswapV3Factory.connect(wallet).createPool(
        pool.token0,
        pool.token1,
        pool.fee
      )
      await poolCreationTx.wait()
    }
  }

  // Step X. (UNISWAP) Transfer code from testnet to state dump.
  for (const pool of pools) {
    const poolCode = await wallet.provider.getCode(pool.newAddress)
    assert.notEqual(poolCode, '0x', `pool has no code: ${pool.newAddress}`)

    dump.accounts[pool.newAddress] = dump.accounts[pool.oldAddress]
    dump.accounts[pool.newAddress].code = poolCode
    delete dump.accounts[pool.oldAddress]
  }

  /* --- END UNISWAP SURGERY SECTION --- */

  // Step 10. Remove any remaining unverified contracts from the state dump.

  // Step 11. Recompile every remaining contract with the standard compiler.
}

main()
