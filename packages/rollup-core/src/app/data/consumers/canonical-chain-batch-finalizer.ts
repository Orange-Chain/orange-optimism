/* External Imports */
import { getLogger, logError, ScheduledTask } from '@eth-optimism/core-utils'
import { Contract } from 'ethers'
/* Internal Imports */
import {
  BatchSubmission,
  BatchSubmissionStatus,
  L2DataService,
} from '../../../types/data'
import { TransactionReceipt } from 'ethers/providers'

const log = getLogger('canonical-chain-batch-finalizer')

/**
 * Polls the DB for L2 batches ready to send to L1 and submits them.
 */
export class CanonicalChainBatchFinalizer extends ScheduledTask {
  constructor(
    private readonly dataService: L2DataService,
    private readonly canonicalTransactionChain: Contract,
    private readonly confirmationsUntilFinal: number = 1,
    periodMilliseconds = 10_000
  ) {
    super(periodMilliseconds)
  }

  /**
   * @inheritDoc
   *
   * Marks submitted L2 Tx batches as final when they have been sent and finalized.
   */
  public async runTask(): Promise<boolean> {
    let batchToFinalize: BatchSubmission
    try {
      batchToFinalize = await this.dataService.getNextCanonicalChainTransactionBatchToFinalize()
    } catch (e) {
      logError(
        log,
        `Error fetching tx batch that is sent but not finalized! Continuing...`,
        e
      )
      return false
    }

    if (
      !batchToFinalize ||
      batchToFinalize.batchNumber === null ||
      batchToFinalize.batchNumber === undefined ||
      !batchToFinalize.submissionTxHash
    ) {
      log.debug(`No tx batches found to finalize.`)
      return false
    }

    if (batchToFinalize.status !== BatchSubmissionStatus.SENT) {
      const msg = `Received tx batch to finalize in ${
        batchToFinalize.status
      } instead of ${
        BatchSubmissionStatus.SENT
      }. Batch Submission: ${JSON.stringify(batchToFinalize)}.`
      log.error(msg)
      throw msg
    }

    await this.waitForTxBatchConfirms(
      batchToFinalize.submissionTxHash,
      batchToFinalize.batchNumber
    )
    return true
  }

  /**
   * Waits for the configured number of confirms for the provided rollup tx transaction hash and
   * marks the tx as
   *
   * @param txHash The tx hash to wait for.
   * @param batchNumber The rollup batch number in question.
   */
  private async waitForTxBatchConfirms(
    txHash: string,
    batchNumber: number
  ): Promise<void> {
    if (this.confirmationsUntilFinal > 1) {
      try {
        log.debug(
          `Waiting for ${this.confirmationsUntilFinal} confirmations before treating tx batch ${batchNumber} submission as final.`
        )
        const receipt: TransactionReceipt = await this.canonicalTransactionChain.provider.waitForTransaction(
          txHash,
          this.confirmationsUntilFinal
        )
        log.debug(`Batch submission finalized for tx batch ${batchNumber}!`)
      } catch (e) {
        logError(
          log,
          `Error waiting for necessary block confirmations until final!`,
          e
        )
        // TODO: Should we return here? Don't want to resubmit, so I think we should update the DB
      }
    }

    try {
      log.debug(`Marking tx batch ${batchNumber} confirmed!`)
      await this.dataService.markTransactionBatchFinalOnL1(batchNumber, txHash)
      log.debug(`Tx batch ${batchNumber} marked confirmed!`)
    } catch (e) {
      logError(log, `Error marking tx batch ${batchNumber} as confirmed!`, e)
    }
  }
}
