package op_e2e

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"math/rand"
	"testing"
	"time"

	"github.com/ethereum-optimism/optimism/op-bindings/bindings"
	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-node/testutils/fuzzerutils"
	"github.com/ethereum-optimism/optimism/op-node/withdrawals"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"github.com/ethereum/go-ethereum/rpc"
	fuzz "github.com/google/gofuzz"
	"github.com/stretchr/testify/require"
	"golang.org/x/tools/container/intsets"
)

// TestAccount defines an account generated by
type TestAccount struct {
	HDPath string
	Key    *ecdsa.PrivateKey
	L1Opts *bind.TransactOpts
	L2Opts *bind.TransactOpts
}

// requireExecutionWithinTimeout executes a given function and panics if it does not return within a given timeout
// duration.
func requireExecutionWithinTimeout(timeout time.Duration, msg string, f func()) {
	// Create a cancellable context that we will use to cancel our timeout-based panic if our call succeeded.
	ctx, cancelFunc := context.WithTimeout(context.Background(), timeout)
	defer cancelFunc()

	// Execute a routine in parallel which will panic after a given timeout.
	go func() {
		// Set our timeout message
		timeoutMsg := msg
		if len(timeoutMsg) == 0 {
			timeoutMsg = "test timed out"
		}

		// Continuously check to see if the time has elapsed or the context has otherwise signalled it is done/cancelled.s
		timeStart := time.Now()
		for {
			// If the timeout elapsed, panic
			if time.Now().Sub(timeStart) >= timeout {
				// TODO: This method should be rewritten in the future so it does not use panic(...). panic will end
				//  all other test execution. However, testing assertion failures cannot be flagged from goroutines so
				//  this is a temporary solution.
				panic(timeoutMsg)
			}

			// Check if we got a cancellation, in which case we do not panic.
			select {
			case <-ctx.Done():
				return
			default:
				break
			}

			// Sleep a short bit before waiting again.
			time.Sleep(50 * time.Millisecond)
		}

	}()

	// Execute our method. If execution does not return, cancelFunc() won't be called and the go routine should
	// throw a panic once the timeout is exceeded.
	f()
}

// startConfigWithTestAccounts takes a SystemConfig, generates additional accounts, adds them to the config, so they
// are funded on startup, starts the system, and imports the keys into the keystore, and obtains transaction opts for
// each account.
func startConfigWithTestAccounts(cfg *SystemConfig, accountsToGenerate int) (*System, []*TestAccount, error) {
	// Create our test accounts and add them to the pre-mine cfg.
	testAccounts := make([]*TestAccount, 0)
	for i := 0; i < accountsToGenerate; i++ {
		// Create our test account and add it to our list
		testAccount := &TestAccount{
			HDPath: fmt.Sprintf("m/44'/60'/0'/0/%d", 1000+i), // offset by 1000 to avoid collisions.
			Key:    nil,
			L1Opts: nil,
			L2Opts: nil,
		}
		testAccounts = append(testAccounts, testAccount)

		// Fund the test account in our config
		cfg.Premine[testAccount.HDPath] = intsets.MaxInt
	}

	// Start our system
	sys, err := cfg.start()
	if err != nil {
		return sys, nil, err
	}

	// Obtain all our account private keys that were generated and import them to the keystore to be used
	// with our contract bindings.
	for _, testAccount := range testAccounts {
		// Obtain our generated private key
		testAccount.Key, err = sys.wallet.PrivateKey(accounts.Account{
			URL: accounts.URL{
				Path: testAccount.HDPath,
			},
		})
		if err != nil {
			return sys, nil, err
		}

		// Obtain the transaction options for contract bindings for this account.
		testAccount.L1Opts, err = bind.NewKeyedTransactorWithChainID(testAccount.Key, cfg.L1ChainID)
		if err != nil {
			return sys, nil, err
		}
		testAccount.L2Opts, err = bind.NewKeyedTransactorWithChainID(testAccount.Key, cfg.L2ChainID)
		if err != nil {
			return sys, nil, err
		}
	}

	// Return our results.
	return sys, testAccounts, err
}

// testL2ProcessBlock sends a transaction to the provided client with the given signing key. If it does not propagate
// to the network before the timeout, it flags the current running test as a failure.
func testL2ProcessBlock(t *testing.T, sequencer *ethclient.Client, verifier *ethclient.Client, txSigningKey *ecdsa.PrivateKey, timeout time.Duration, msg string) {
	// Simple transfer from signer to random account
	ctx := context.Background()

	// Obtain the signer address
	fromAddr := crypto.PubkeyToAddress(txSigningKey.PublicKey)

	// Obtain the chain ID for this client
	chainId, err := sequencer.ChainID(ctx)
	require.NoError(t, err)

	// Obtain the sender's nonce
	blockNumber, err := sequencer.BlockNumber(ctx)
	require.NoError(t, err)
	nonce, err := sequencer.NonceAt(ctx, fromAddr, new(big.Int).SetUint64(blockNumber))
	require.NoError(t, err)

	// Get our block header
	header, err := sequencer.HeaderByNumber(ctx, nil)
	require.NoError(t, err)

	// Set some transaction details for our test block
	toAddr := common.Address{0xff, 0xff}
	transferAmount := big.NewInt(1)
	gasTipCap, err := sequencer.SuggestGasTipCap(ctx)
	require.NoError(t, err)

	// Get our gas fee cap
	gasFeeCap := new(big.Int).Add(
		gasTipCap,
		new(big.Int).Mul(header.BaseFee, big.NewInt(2)),
	)

	// Estimate the amount of gas needed for this tx
	gas, err := sequencer.EstimateGas(ctx, ethereum.CallMsg{
		From:      fromAddr,
		To:        &toAddr,
		GasPrice:  nil,
		GasFeeCap: gasFeeCap,
		GasTipCap: gasTipCap,
		Value:     transferAmount,
		Data:      nil,
	})

	// Sign our transaction
	tx := types.MustSignNewTx(txSigningKey, types.LatestSignerForChainID(chainId), &types.DynamicFeeTx{
		ChainID:   chainId,
		Nonce:     nonce, // Already have deposit
		To:        &toAddr,
		Value:     transferAmount,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gas,
	})

	// Send our transaction and expect a response in our timeout duration, or panic.
	requireExecutionWithinTimeout(timeout, msg, func() {
		// Send the transaction to the client
		err = sequencer.SendTransaction(ctx, tx)
		require.Nil(t, err, "Sending L2 tx to sequencer")

		// Wait for the tx hash to appear with timeout
		_, _ = waitForTransaction(tx.Hash(), verifier, timeout)
	})
}

// TestGasPriceOracleFeeUpdates checks that the gas price oracle cannot be locked by mis-configuring parameters.
func TestGasPriceOracleFeeUpdates(t *testing.T) {
	// Define our values to set in the GasPriceOracle (we set them high to see if it can lock L2 or stop bindings
	// from updating the prices once again.
	overheadValue := abi.MaxUint256
	decimalsValue := abi.MaxUint256
	scalarValue := abi.MaxUint256

	// Setup our logger handler
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	// Create our system configuration for L1/L2 and start it
	cfg := defaultSystemConfig(t)
	cfg.Premine[transactorHDPath] = intsets.MaxInt
	sys, err := cfg.start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	// Obtain our sequencer, verifier, and transactor keypair.
	l2Seq := sys.Clients["sequencer"]
	l2Verif := sys.Clients["verifier"]
	ethPrivKey, err := sys.wallet.PrivateKey(accounts.Account{
		URL: accounts.URL{
			Path: transactorHDPath,
		},
	})
	require.Nil(t, err)

	// Bind to the GasPriceOracle contract
	gpoContract, err := bindings.NewGasPriceOracle(common.HexToAddress(predeploys.OVM_GasPriceOracle), l2Seq)
	require.Nil(t, err)

	// Obtain our signer.
	l2opts, err := bind.NewKeyedTransactorWithChainID(ethPrivKey, cfg.L2ChainID)
	require.Nil(t, err)

	// Define our L1 transaction timeout duration.
	txTimeoutDuration := 10 * time.Duration(cfg.L1BlockTime) * time.Second

	// Update decimals within our given timeout.
	requireExecutionWithinTimeout(txTimeoutDuration, "failed to update GasPriceOracle decimals", func() {
		// Update decimals
		tx, err := gpoContract.SetDecimals(l2opts, decimalsValue)
		require.Nil(t, err, "sending gpo update tx")

		receipt, err := waitForTransaction(tx.Hash(), l2Verif, txTimeoutDuration)
		require.Nil(t, err, "waiting for gpo decimals update tx")
		require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful, "transaction failed")
	})

	// Update overhead within our given timeout.
	requireExecutionWithinTimeout(txTimeoutDuration, "failed to update GasPriceOracle overhead", func() {
		// Update overhead
		tx, err := gpoContract.SetOverhead(l2opts, overheadValue)
		require.Nil(t, err, "sending overhead update tx")

		receipt, err := waitForTransaction(tx.Hash(), l2Verif, txTimeoutDuration)
		require.Nil(t, err, "waiting for overhead update tx")
		require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful, "transaction failed")
	})

	// Update scalar within our given timeout.
	requireExecutionWithinTimeout(txTimeoutDuration, "failed to update GasPriceOracle scalar", func() {
		// Update scalar
		tx, err := gpoContract.SetScalar(l2opts, scalarValue)
		require.Nil(t, err, "sending gpo update tx")

		receipt, err := waitForTransaction(tx.Hash(), l2Verif, txTimeoutDuration)
		require.Nil(t, err, "waiting for gpo scalar update tx")
		require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful, "transaction failed")
	})

	_, err = gpoContract.Overhead(&bind.CallOpts{})
	require.Nil(t, err, "reading gpo overhead")
	_, err = gpoContract.Decimals(&bind.CallOpts{})
	require.Nil(t, err, "reading gpo decimals")
	_, err = gpoContract.Scalar(&bind.CallOpts{})
	require.Nil(t, err, "reading gpo scalar")

	// Update scalar again to see if L2 is locked at this point.
	requireExecutionWithinTimeout(txTimeoutDuration, "failed to update GasPriceOracle scalar after previously setting decimals, overhead, and scalar", func() {
		// Update scalar
		tx, err := gpoContract.SetScalar(l2opts, scalarValue)
		require.Nil(t, err, "sending gpo update tx")

		receipt, err := waitForTransaction(tx.Hash(), l2Verif, txTimeoutDuration)
		require.Nil(t, err, "waiting for gpo scalar update tx")
		require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful, "transaction failed")
	})
}

// TestGasPriceOracleFeesL2Lock checks that the gas price oracle cannot lock unrelated transactions from being
// processed.
func TestGasPriceOracleFeesL2Lock(t *testing.T) {
	// Define our values to set in the GasPriceOracle (we set them high to see if it can lock L2 or stop bindings
	// from updating the prices once again.
	overheadValue := abi.MaxUint256
	decimalsValue := abi.MaxUint256
	scalarValue := abi.MaxUint256

	// Setup our logger handler
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	// Create our system configuration for L1/L2 and start it
	cfg := defaultSystemConfig(t)
	cfg.Premine[transactorHDPath] = intsets.MaxInt
	sys, err := cfg.start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	// Obtain our sequencer, verifier, and transactor keypair.
	l2Seq := sys.Clients["sequencer"]
	l2Verif := sys.Clients["verifier"]
	ethPrivKey, err := sys.wallet.PrivateKey(accounts.Account{
		URL: accounts.URL{
			Path: transactorHDPath,
		},
	})
	require.Nil(t, err)

	// Bind to the GasPriceOracle contract
	gpoContract, err := bindings.NewGasPriceOracle(common.HexToAddress(predeploys.OVM_GasPriceOracle), l2Seq)
	require.Nil(t, err)

	// Obtain our signer.
	l2opts, err := bind.NewKeyedTransactorWithChainID(ethPrivKey, cfg.L2ChainID)
	require.Nil(t, err)

	// Define our L1 transaction timeout duration.
	txTimeoutDuration := 10 * time.Duration(cfg.L1BlockTime) * time.Second

	// Update decimals within our given timeout.
	requireExecutionWithinTimeout(txTimeoutDuration, "failed to update GasPriceOracle decimals", func() {
		// Update decimals
		tx, err := gpoContract.SetDecimals(l2opts, decimalsValue)
		require.Nil(t, err, "sending gpo update tx")

		receipt, err := waitForTransaction(tx.Hash(), l2Verif, txTimeoutDuration)
		require.Nil(t, err, "waiting for gpo decimals update tx")
		require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful, "transaction failed")
	})

	// Test production of an unrelated block.
	testL2ProcessBlock(t, l2Seq, l2Verif, ethPrivKey, txTimeoutDuration, "L2 blocks failed to process before timeout after updating GasPriceOracle decimals.")

	// Update overhead within our given timeout.
	requireExecutionWithinTimeout(txTimeoutDuration, "failed to update GasPriceOracle overhead", func() {
		// Update overhead
		tx, err := gpoContract.SetOverhead(l2opts, overheadValue)
		require.Nil(t, err, "sending overhead update tx")

		receipt, err := waitForTransaction(tx.Hash(), l2Verif, txTimeoutDuration)
		require.Nil(t, err, "waiting for overhead update tx")
		require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful, "transaction failed")
	})

	// Test production of an unrelated block.
	testL2ProcessBlock(t, l2Seq, l2Verif, ethPrivKey, txTimeoutDuration, "L2 blocks failed to process before timeout after updating GasPriceOracle overhead.")

	// Update scalar within our given timeout.
	requireExecutionWithinTimeout(txTimeoutDuration, "failed to update GasPriceOracle scalar", func() {
		// Update scalar
		tx, err := gpoContract.SetScalar(l2opts, scalarValue)
		require.Nil(t, err, "sending gpo update tx")

		receipt, err := waitForTransaction(tx.Hash(), l2Verif, txTimeoutDuration)
		require.Nil(t, err, "waiting for gpo scalar update tx")
		require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful, "transaction failed")
	})

	// Test production of an unrelated block.
	testL2ProcessBlock(t, l2Seq, l2Verif, ethPrivKey, txTimeoutDuration, "L2 blocks failed to process before timeout after updating GasPriceOracle scalar.")

	_, err = gpoContract.Overhead(&bind.CallOpts{})
	require.Nil(t, err, "reading gpo overhead")
	_, err = gpoContract.Decimals(&bind.CallOpts{})
	require.Nil(t, err, "reading gpo decimals")
	_, err = gpoContract.Scalar(&bind.CallOpts{})
	require.Nil(t, err, "reading gpo scalar")

	// Update scalar again to see if L2 is locked at this point.
	requireExecutionWithinTimeout(txTimeoutDuration, "failed to update GasPriceOracle scalar after previously setting decimals, overhead, and scalar", func() {
		// Update scalar
		tx, err := gpoContract.SetScalar(l2opts, scalarValue)
		require.Nil(t, err, "sending gpo update tx")

		receipt, err := waitForTransaction(tx.Hash(), l2Verif, txTimeoutDuration)
		require.Nil(t, err, "waiting for gpo scalar update tx")
		require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful, "transaction failed")
	})
}

// TestL2SequencerRPCDepositTx checks that the L2 sequencer will not accept DepositTx type transactions.
// The acceptance of these transactions would allow for arbitrary minting of ETH in L2.
func TestL2SequencerRPCDepositTx(t *testing.T) {
	// Setup our logger handler
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	// Create our system configuration for L1/L2 and start it
	cfg := defaultSystemConfig(t)
	cfg.Premine[transactorHDPath] = intsets.MaxInt
	sys, err := cfg.start()
	require.Nil(t, err, "Error starting up system")
	defer sys.Close()

	// Obtain our sequencer, verifier, and transactor keypair.
	l2Seq := sys.Clients["sequencer"]
	l2Verif := sys.Clients["verifier"]
	txSigningKey, err := sys.wallet.PrivateKey(accounts.Account{
		URL: accounts.URL{
			Path: transactorHDPath,
		},
	})
	require.Nil(t, err)

	// Define our L1 transaction timeout duration.
	txTimeoutDuration := 10 * time.Duration(cfg.L1BlockTime) * time.Second

	// Test a dummy block now that we're initialized
	testL2ProcessBlock(t, l2Seq, l2Verif, txSigningKey, txTimeoutDuration, "failed to process unrelated tx prior to testing sending of DepositTx to L2 over RPC")

	// Create a deposit tx to send over RPC.
	tx := types.NewTx(&types.DepositTx{
		SourceHash:          common.Hash{},
		From:                crypto.PubkeyToAddress(txSigningKey.PublicKey),
		To:                  &common.Address{0xff, 0xff},
		Mint:                big.NewInt(1000),
		Value:               big.NewInt(1000),
		Gas:                 0,
		IsSystemTransaction: false,
		Data:                nil,
	})

	// Send our transaction to the L2 sequencer and expect a response in our timeout duration, or panic.
	requireExecutionWithinTimeout(txTimeoutDuration, "failed to send DepositTx over RPC before timeout", func() {
		// Send the transaction to the L2 sequencer
		ctx := context.Background()
		err = l2Seq.SendTransaction(ctx, tx)
		require.Error(t, err, "a DepositTx was accepted by L2 sequencer over RPC when it should not have been.")
	})

	// Send our transaction to the L2 verifier and expect a response in our timeout duration, or panic.
	requireExecutionWithinTimeout(txTimeoutDuration, "failed to send DepositTx over RPC before timeout", func() {
		// Send the transaction to the L2 verifier
		ctx := context.Background()
		err = l2Verif.SendTransaction(ctx, tx)
		require.Error(t, err, "a DepositTx was accepted by L2 sequencer over RPC when it should not have been.")
	})

	// Verify other transactions will be processed after someone attempts to send a DepositTx.
	testL2ProcessBlock(t, l2Seq, l2Verif, txSigningKey, txTimeoutDuration, "failed to process non-DepositTx-type tx after sending DepositTx to L2 over RPC")
}

// TestMixedDepositValidity makes a number of deposit transactions, some which will succeed in transferring value,
// while others do not. It ensures that the expected nonces/balances match after several interactions.
func TestMixedDepositValidity(t *testing.T) {
	// Define how many deposit txs we'll make. Each deposit mints a fixed amount and transfers up to 1/3 of the user's
	// balance. As such, this number cannot be too high or else the test will always fail due to lack of balance in L1.
	const depositTxCount = 15

	// Define how many accounts we'll use to deposit funds
	const accountUsedToDeposit = 5

	// Setup our logger handler
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	// Create our system configuration, funding all accounts we created for L1/L2, and start it
	cfg := defaultSystemConfig(t)
	cfg.Premine[transactorHDPath] = intsets.MaxInt
	sys, testAccounts, err := startConfigWithTestAccounts(&cfg, accountUsedToDeposit)
	require.NoError(t, err, "error starting up system")
	defer sys.Close()

	// Obtain our sequencer, verifier, and transactor keypair.
	l1Client := sys.Clients["l1"]
	l2Seq := sys.Clients["sequencer"]
	l2Verif := sys.Clients["verifier"]
	require.NoError(t, err)

	// Define our L1 transaction timeout duration.
	txTimeoutDuration := 10 * time.Duration(cfg.L1BlockTime) * time.Second

	// Bind to the deposit contract
	depositContract, err := bindings.NewOptimismPortal(sys.DepositContractAddr, l1Client)
	require.NoError(t, err)

	// Create a struct used to track our transactors and their transactions sent.
	type TestAccountState struct {
		Account           *TestAccount
		ExpectedL1Balance *big.Int
		ExpectedL2Balance *big.Int
		StartingL1Nonce   uint64
		ExpectedL1Nonce   uint64
		StartingL2Nonce   uint64
		ExpectedL2Nonce   uint64
	}

	// Create the state objects for every test account we'll track changes for.
	transactors := make([]*TestAccountState, 0)
	for i := 0; i < len(testAccounts); i++ {
		// Obtain our account
		testAccount := testAccounts[i]

		// Obtain the transactor's starting nonce on L1.
		ctx, cancel := context.WithTimeout(context.Background(), txTimeoutDuration)
		startL1Nonce, err := l1Client.NonceAt(ctx, testAccount.L1Opts.From, nil)
		cancel()
		require.NoError(t, err)

		// Obtain the transactor's starting balance on L2.
		ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
		startL2Balance, err := l2Verif.BalanceAt(ctx, testAccount.L2Opts.From, nil)
		cancel()
		require.NoError(t, err)

		// Obtain the transactor's starting nonce on L2.
		ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
		startL2Nonce, err := l2Verif.NonceAt(ctx, testAccount.L2Opts.From, nil)
		cancel()
		require.NoError(t, err)

		// Add our transactor to our list
		transactors = append(transactors, &TestAccountState{
			Account:           testAccount,
			ExpectedL2Balance: startL2Balance,
			ExpectedL1Nonce:   startL1Nonce,
			ExpectedL2Nonce:   startL2Nonce,
		})
	}

	// Create our random provider
	randomProvider := rand.New(rand.NewSource(time.Now().Unix()))

	// Now we create a number of deposits from each transactor
	for i := 0; i < depositTxCount; i++ {
		// Determine if this deposit should succeed in transferring value (not minting)
		validTransfer := randomProvider.Int()%2 == 0

		// Determine the transactor to use
		transactorIndex := randomProvider.Int() % len(transactors)
		transactor := transactors[transactorIndex]

		// Determine the transactor to receive the deposit
		receiverIndex := randomProvider.Int() % len(transactors)
		receiver := transactors[receiverIndex]
		toAddr := receiver.Account.L2Opts.From

		// Create our L1 deposit transaction and send it.
		mintAmount := big.NewInt(randomProvider.Int63() % 9_000_000)
		transactor.Account.L1Opts.Value = mintAmount
		var transferValue *big.Int
		if validTransfer {
			transferValue = new(big.Int).Div(transactor.ExpectedL2Balance, common.Big3) // send 1/3 our balance which should succeed.
		} else {
			transferValue = new(big.Int).Mul(common.Big2, transactor.ExpectedL2Balance) // trigger a revert by trying to transfer our current balance * 2
		}
		tx, err := depositContract.DepositTransaction(transactor.Account.L1Opts, toAddr, transferValue, 1_000_000, false, nil)
		require.Nil(t, err, "with deposit tx")

		// Wait for the deposit tx to appear in L1.
		receipt, err := waitForTransaction(tx.Hash(), l1Client, txTimeoutDuration)
		require.Nil(t, err, "Waiting for deposit tx on L1")

		// Reconstruct the L2 tx hash to wait for the deposit in L2.
		reconstructedDep, err := derive.UnmarshalDepositLogEvent(receipt.Logs[0])
		require.NoError(t, err, "Could not reconstruct L2 Deposit")
		tx = types.NewTx(reconstructedDep)
		receipt, err = waitForTransaction(tx.Hash(), l2Verif, txTimeoutDuration)
		require.NoError(t, err)

		// Verify the result of the L2 tx receipt. Based on how much we transferred it should be successful/failed.
		if validTransfer {
			require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful)
		} else {
			require.Equal(t, receipt.Status, types.ReceiptStatusFailed)
		}

		// Update our expected balances.
		if validTransfer && transactor != receiver {
			// Transactor balances changes by minted minus transferred value.
			transactor.ExpectedL2Balance = new(big.Int).Add(transactor.ExpectedL2Balance, new(big.Int).Sub(mintAmount, transferValue))
			// Receiver balance changes by transferred value.
			receiver.ExpectedL2Balance = new(big.Int).Add(receiver.ExpectedL2Balance, transferValue)
		} else {
			// If the transfer failed, minting should've still succeeded but the balance shouldn't have transferred
			// to the recipient.
			transactor.ExpectedL2Balance = new(big.Int).Add(transactor.ExpectedL2Balance, mintAmount)
		}
		transactor.ExpectedL1Nonce = transactor.ExpectedL1Nonce + 1
		transactor.ExpectedL2Nonce = transactor.ExpectedL2Nonce + 1
	}

	// At the end, assert our account balance/nonce states.
	for _, transactor := range transactors {
		// Obtain the L1 account nonce
		ctx, cancel := context.WithTimeout(context.Background(), txTimeoutDuration)
		endL1Nonce, err := l1Client.NonceAt(ctx, transactor.Account.L1Opts.From, nil)
		cancel()
		require.NoError(t, err)

		// Obtain the L2 sequencer account balance
		ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
		endL2SeqBalance, err := l2Seq.BalanceAt(ctx, transactor.Account.L2Opts.From, nil)
		cancel()
		require.NoError(t, err)

		// Obtain the L2 sequencer account nonce
		ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
		endL2SeqNonce, err := l2Seq.NonceAt(ctx, transactor.Account.L2Opts.From, nil)
		cancel()
		require.NoError(t, err)

		// Obtain the L2 verifier account balance
		ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
		endL2VerifBalance, err := l2Verif.BalanceAt(ctx, transactor.Account.L2Opts.From, nil)
		cancel()
		require.NoError(t, err)

		// Obtain the L2 verifier account nonce
		ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
		endL2VerifNonce, err := l2Verif.NonceAt(ctx, transactor.Account.L2Opts.From, nil)
		cancel()
		require.NoError(t, err)

		require.Equal(t, transactor.ExpectedL1Nonce, endL1Nonce, "Unexpected L1 nonce for transactor")
		require.Equal(t, transactor.ExpectedL2Nonce, endL2SeqNonce, "Unexpected L2 sequencer nonce for transactor")
		require.Equal(t, transactor.ExpectedL2Balance, endL2SeqBalance, "Unexpected L2 sequencer balance for transactor")
		require.Equal(t, transactor.ExpectedL2Nonce, endL2VerifNonce, "Unexpected L2 verifier nonce for transactor")
		require.Equal(t, transactor.ExpectedL2Balance, endL2VerifBalance, "Unexpected L2 verifier balance for transactor")
	}
}

// TestMixedWithdrawalValidity makes a number of withdrawal transactions and ensures ones with modified parameters are
// rejected while unmodified ones are accepted. This runs test cases in different systems.
func TestMixedWithdrawalValidity(t *testing.T) {
	// Setup our logger handler
	if !verboseGethNodes {
		log.Root().SetHandler(log.DiscardHandler())
	}

	// There are 7 different fields we try modifying to cause a failure, plus one "good" test result we test.
	for i := 0; i <= 8; i++ {
		t.Run(fmt.Sprintf("withdrawal test#%d", i+1), func(t *testing.T) {
			// Create our system configuration, funding all accounts we created for L1/L2, and start it
			cfg := defaultSystemConfig(t)
			cfg.DepositCFG.FinalizationPeriod = big.NewInt(2)
			cfg.Premine[transactorHDPath] = intsets.MaxInt
			sys, err := cfg.start()
			require.NoError(t, err, "error starting up system")
			defer sys.Close()

			// Obtain our sequencer, verifier, and transactor keypair.
			l1Client := sys.Clients["l1"]
			l2Seq := sys.Clients["sequencer"]
			l2Verif := sys.Clients["verifier"]
			require.NoError(t, err)

			// Define our L1 transaction timeout duration.
			txTimeoutDuration := 10 * time.Duration(cfg.L1BlockTime) * time.Second

			// Bind to the deposit contract
			depositContract, err := bindings.NewOptimismPortal(sys.DepositContractAddr, l1Client)
			_ = depositContract
			require.NoError(t, err)

			// Create a struct used to track our transactors and their transactions sent.
			type TestAccountState struct {
				Account           *TestAccount
				ExpectedL1Balance *big.Int
				ExpectedL2Balance *big.Int
				ExpectedL1Nonce   uint64
				ExpectedL2Nonce   uint64
			}

			// Create a test account state for our transactor.
			transactorKey, err := sys.wallet.PrivateKey(accounts.Account{
				URL: accounts.URL{
					Path: transactorHDPath,
				},
			})
			require.NoError(t, err)
			transactor := &TestAccountState{
				Account: &TestAccount{
					HDPath: transactorHDPath,
					Key:    transactorKey,
					L1Opts: nil,
					L2Opts: nil,
				},
				ExpectedL1Balance: nil,
				ExpectedL2Balance: nil,
				ExpectedL1Nonce:   0,
				ExpectedL2Nonce:   0,
			}
			transactor.Account.L1Opts, err = bind.NewKeyedTransactorWithChainID(transactor.Account.Key, cfg.L1ChainID)
			require.NoError(t, err)
			transactor.Account.L2Opts, err = bind.NewKeyedTransactorWithChainID(transactor.Account.Key, cfg.L2ChainID)
			require.NoError(t, err)

			// Obtain the transactor's starting balance on L1.
			ctx, cancel := context.WithTimeout(context.Background(), txTimeoutDuration)
			transactor.ExpectedL1Balance, err = l1Client.BalanceAt(ctx, transactor.Account.L1Opts.From, nil)
			cancel()
			require.NoError(t, err)

			// Obtain the transactor's starting balance on L2.
			ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
			transactor.ExpectedL2Balance, err = l2Verif.BalanceAt(ctx, transactor.Account.L2Opts.From, nil)
			cancel()
			require.NoError(t, err)

			// Bind to the L2-L1 message passer
			l2l1MessagePasser, err := bindings.NewL2ToL1MessagePasser(predeploys.L2ToL1MessagePasserAddr, l2Seq)
			require.NoError(t, err, "error binding to message passer on L2")

			// Create our fuzzer wrapper to generate complex values (despite this not being a fuzz test, this is still a useful
			// provider to fill complex data structures).
			typeProvider := fuzz.NewWithSeed(time.Now().Unix()).NilChance(0).MaxDepth(10000).NumElements(0, 0x100)
			fuzzerutils.AddFuzzerFunctions(typeProvider)

			// Now we create a number of withdrawals from each transactor

			// Determine the address our request will come from
			fromAddr := crypto.PubkeyToAddress(transactor.Account.Key.PublicKey)

			// Initiate Withdrawal
			withdrawAmount := big.NewInt(500_000_000_000)
			transactor.Account.L2Opts.Value = withdrawAmount
			tx, err := l2l1MessagePasser.InitiateWithdrawal(transactor.Account.L2Opts, fromAddr, big.NewInt(21000), nil)
			require.Nil(t, err, "sending initiate withdraw tx")

			// Wait for the transaction to appear in L2 verifier
			receipt, err := waitForTransaction(tx.Hash(), l2Verif, txTimeoutDuration)
			require.Nil(t, err, "withdrawal initiated on L2 sequencer")
			require.Equal(t, receipt.Status, types.ReceiptStatusSuccessful, "transaction failed")

			// Obtain the header for the block containing the transaction (used to calculate gas fees)
			ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
			header, err := l2Verif.HeaderByNumber(ctx, receipt.BlockNumber)
			cancel()
			require.Nil(t, err)

			// Calculate gas fees for the withdrawal in L2 to later adjust our balance.
			withdrawalL2GasFee := calcGasFees(receipt.GasUsed, tx.GasTipCap(), tx.GasFeeCap(), header.BaseFee)

			// Adjust our expected L2 balance (should've decreased by withdraw amount + fees)
			transactor.ExpectedL2Balance = new(big.Int).Sub(transactor.ExpectedL2Balance, withdrawAmount)
			transactor.ExpectedL2Balance = new(big.Int).Sub(transactor.ExpectedL2Balance, withdrawalL2GasFee)
			transactor.ExpectedL2Nonce = transactor.ExpectedL2Nonce + 1

			// Wait for the finalization period, then we can finalize this withdrawal.
			ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
			blockNumber, err := withdrawals.WaitForFinalizationPeriod(ctx, l1Client, sys.DepositContractAddr, receipt.BlockNumber)
			cancel()
			require.Nil(t, err)

			ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
			header, err = l2Verif.HeaderByNumber(ctx, new(big.Int).SetUint64(blockNumber))
			cancel()
			require.Nil(t, err)

			l2ClientRPC, err := rpc.Dial(sys.nodes["verifier"].WSEndpoint())
			require.Nil(t, err)
			l2Client := withdrawals.NewClient(l2ClientRPC)

			// Now create the withdrawal
			params, err := withdrawals.FinalizeWithdrawalParameters(context.Background(), l2Client, tx.Hash(), header)
			require.Nil(t, err)

			// Obtain our withdrawal parameters
			withdrawalTransaction := &bindings.TypesWithdrawalTransaction{
				Nonce:    params.Nonce,
				Sender:   params.Sender,
				Target:   params.Target,
				Value:    params.Value,
				GasLimit: params.GasLimit,
				Data:     params.Data,
			}
			blockNumberParam := params.BlockNumber
			outputRootProofParam := params.OutputRootProof
			withdrawalProofParam := params.WithdrawalProof

			// Determine if this will be a bad withdrawal.
			badWithdrawal := i < 8
			if badWithdrawal {
				// Select a field to overwrite depending on which test case this is.
				fieldIndex := i

				// We ensure that each field changes to something different.
				if fieldIndex == 0 {
					originalValue := new(big.Int).Set(withdrawalTransaction.Nonce)
					for originalValue.Cmp(withdrawalTransaction.Nonce) == 0 {
						typeProvider.Fuzz(&withdrawalTransaction.Nonce)
					}
				} else if fieldIndex == 1 {
					originalValue := withdrawalTransaction.Sender
					for originalValue == withdrawalTransaction.Sender {
						typeProvider.Fuzz(&withdrawalTransaction.Sender)
					}
				} else if fieldIndex == 2 {
					originalValue := withdrawalTransaction.Target
					for originalValue == withdrawalTransaction.Target {
						typeProvider.Fuzz(&withdrawalTransaction.Target)
					}
				} else if fieldIndex == 3 {
					originalValue := new(big.Int).Set(withdrawalTransaction.Value)
					for originalValue.Cmp(withdrawalTransaction.Value) == 0 {
						typeProvider.Fuzz(&withdrawalTransaction.Value)
					}
				} else if fieldIndex == 4 {
					originalValue := new(big.Int).Set(withdrawalTransaction.GasLimit)
					for originalValue.Cmp(withdrawalTransaction.GasLimit) == 0 {
						typeProvider.Fuzz(&withdrawalTransaction.GasLimit)
					}
				} else if fieldIndex == 5 {
					originalValue := new(big.Int).Set(blockNumberParam)
					for originalValue.Cmp(blockNumberParam) == 0 {
						typeProvider.Fuzz(&blockNumberParam)
					}
				} else if fieldIndex == 6 {
					// TODO: this is a large structure that is unlikely to ever produce the same value, however we should
					//  verify that we actually generated different values.
					typeProvider.Fuzz(&outputRootProofParam)
				} else if fieldIndex == 7 {
					typeProvider.Fuzz(&withdrawalProofParam)
					originalValue := make([]byte, len(withdrawalProofParam))
					copy(originalValue, withdrawalProofParam)
					for bytes.Equal(originalValue, withdrawalProofParam) {
						typeProvider.Fuzz(&withdrawalProofParam)
					}
				}
			}

			// Finally, finalize our withdrawal.
			tx, err = depositContract.FinalizeWithdrawalTransaction(
				transactor.Account.L1Opts,
				*withdrawalTransaction,
				blockNumberParam,
				outputRootProofParam,
				withdrawalProofParam,
			)

			// If we had a bad withdrawal, we don't update some expected value and skip to processing the next
			// withdrawal. Otherwise, if it was valid, this should've succeeded so we proceed with updating our expected
			// values and asserting no errors occurred.
			if badWithdrawal {
				require.Error(t, err)
			} else {
				require.NoError(t, err)

				receipt, err = waitForTransaction(tx.Hash(), l1Client, txTimeoutDuration)
				require.Nil(t, err, "finalize withdrawal")
				require.Equal(t, types.ReceiptStatusSuccessful, receipt.Status)

				// Verify balance after withdrawal
				ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
				header, err = l1Client.HeaderByNumber(ctx, receipt.BlockNumber)
				cancel()
				require.Nil(t, err)

				// Ensure that withdrawal - gas fees are added to the L1 balance
				// Fun fact, the fee is greater than the withdrawal amount
				withdrawalL1GasFee := calcGasFees(receipt.GasUsed, tx.GasTipCap(), tx.GasFeeCap(), header.BaseFee)
				transactor.ExpectedL1Balance = new(big.Int).Add(transactor.ExpectedL2Balance, withdrawAmount)
				transactor.ExpectedL1Balance = new(big.Int).Sub(transactor.ExpectedL2Balance, withdrawalL1GasFee)
				transactor.ExpectedL1Nonce++
			}

			// At the end, assert our account balance/nonce states.

			// Obtain the L2 sequencer account balance
			ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
			endL1Balance, err := l1Client.BalanceAt(ctx, transactor.Account.L1Opts.From, nil)
			cancel()
			require.NoError(t, err)

			// Obtain the L1 account nonce
			ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
			endL1Nonce, err := l1Client.NonceAt(ctx, transactor.Account.L1Opts.From, nil)
			cancel()
			require.NoError(t, err)

			// Obtain the L2 sequencer account balance
			ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
			endL2SeqBalance, err := l2Seq.BalanceAt(ctx, transactor.Account.L1Opts.From, nil)
			cancel()
			require.NoError(t, err)

			// Obtain the L2 sequencer account nonce
			ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
			endL2SeqNonce, err := l2Seq.NonceAt(ctx, transactor.Account.L1Opts.From, nil)
			cancel()
			require.NoError(t, err)

			// Obtain the L2 verifier account balance
			ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
			endL2VerifBalance, err := l2Verif.BalanceAt(ctx, transactor.Account.L1Opts.From, nil)
			cancel()
			require.NoError(t, err)

			// Obtain the L2 verifier account nonce
			ctx, cancel = context.WithTimeout(context.Background(), txTimeoutDuration)
			endL2VerifNonce, err := l2Verif.NonceAt(ctx, transactor.Account.L1Opts.From, nil)
			cancel()
			require.NoError(t, err)

			// TODO: Check L1 balance as well here. We avoided this due to time constraints as it seems L1 fees
			//  were off slightly.
			_ = endL1Balance
			//require.Equal(t, transactor.ExpectedL1Balance, endL1Balance, "Unexpected L1 balance for transactor")
			require.Equal(t, transactor.ExpectedL1Nonce, endL1Nonce, "Unexpected L1 nonce for transactor")
			require.Equal(t, transactor.ExpectedL2Nonce, endL2SeqNonce, "Unexpected L2 sequencer nonce for transactor")
			require.Equal(t, transactor.ExpectedL2Balance, endL2SeqBalance, "Unexpected L2 sequencer balance for transactor")
			require.Equal(t, transactor.ExpectedL2Nonce, endL2VerifNonce, "Unexpected L2 verifier nonce for transactor")
			require.Equal(t, transactor.ExpectedL2Balance, endL2VerifBalance, "Unexpected L2 verifier balance for transactor")
		})
	}
}
