package trie

import (
	"crypto/rand"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethdb/memorydb"
	"github.com/ethereum/go-ethereum/rlp"
	"github.com/ethereum/go-ethereum/trie"
)

// Generate an abi-encoded `trieTestCase` of a specified variant
func FuzzTrie(variant string) {
	if len(variant) < 1 {
		log.Fatal("Must pass a variant to the trie fuzzer!")
	}

	var testCase trieTestCase
	switch variant {
	case "valid":
		testCase = genValidTrieTestCase()
		break
	case "extra_proof_elems":
		testCase = genValidTrieTestCase()
		// Duplicate the last element of the proof
		testCase.Proof = append(testCase.Proof, [][]byte{testCase.Proof[len(testCase.Proof)-1]}...)
		break
	case "corrupted_proof":
		testCase = genValidTrieTestCase()

		// Re-encode a random element within the proof
		idx := randRange(0, int64(len(testCase.Proof)))
		encoded, _ := rlp.EncodeToBytes(testCase.Proof[idx])
		testCase.Proof[idx] = encoded
		break
	case "invalid_data_remainder":
		testCase = genValidTrieTestCase()

		// Alter true length of random proof element by appending random bytes
		// Do not update the encoded length
		idx := randRange(0, int64(len(testCase.Proof)))
		bytes := make([]byte, randRange(1, 512))
		rand.Read(bytes)
		testCase.Proof[idx] = append(testCase.Proof[idx], bytes...)
		break
	case "invalid_large_internal_hash":
		testCase = genValidTrieTestCase()

		// Clobber 10 bytes at a random location within a random proof element
		idx := randRange(1, int64(len(testCase.Proof)))
		b := make([]byte, 10)
		rand.Read(b)
		st := randRange(10, int64(len(testCase.Proof[idx])-10))
		testCase.Proof[idx] = append(testCase.Proof[idx][0:st], append(b, testCase.Proof[idx][st+10:]...)...)
		break
	case "invalid_internal_node_hash":
		testCase = genValidTrieTestCase()
		// Assign the last proof element to an encoded list containing a
		// random 29 byte value
		b := make([]byte, 29)
		rand.Read(b)
		e, _ := rlp.EncodeToBytes(b)
		testCase.Proof[len(testCase.Proof)-1] = append([]byte{0xc0 + 30}, e...)
	default:
		log.Fatal("Invalid variant passed to trie fuzzer!")
	}

	// Print encoded test case with no newline so that foundry's FFI can read the output
	fmt.Print(testCase.AbiEncode())
}

// Generate a random test case for Bedrock's MerkleTrie verifier.
func genValidTrieTestCase() trieTestCase {
	// Create an empty merkle trie
	memdb := memorydb.New()
	randTrie := trie.NewEmpty(trie.NewDatabase(memdb))

	// Get a random number of elements to put into the trie
	randN := randRange(2, 1024)
	// Get a random key/value pair to generate a proof of inclusion for
	randSelect := randRange(0, randN)

	// Create a fixed-length key as well as a randomly-sized value
	// We create these out of the loop to reduce mem allocations.
	randKey := make([]byte, 32)
	randValue := make([]byte, randRange(2, 1024))

	// Randomly selected key/value for proof generation
	var key []byte
	var value []byte

	// Add `randN` elements to the trie
	var i int64
	for ; i < randN; i++ {
		// Randomize the contents of `randKey` and `randValue`
		rand.Read(randKey)
		rand.Read(randValue)

		// Insert the random k/v pair into the trie
		if err := randTrie.TryUpdate(randKey, randValue); err != nil {
			log.Fatal("Error adding key-value pair to trie")
		}

		// If this is our randomly selected k/v pair, store it in `key` & `value`
		if i == randSelect {
			key = randKey
			value = randValue
		}
	}

	// Generate proof for `key`'s inclusion in our trie
	var proof proofList
	if err := randTrie.Prove(key, 0, &proof); err != nil {
		log.Fatal("Error creating proof for randomly selected key's inclusion in generated trie")
	}

	// Create our test case with the data collected
	testCase := trieTestCase{
		Root:  (*[32]byte)(randTrie.Hash().Bytes()),
		Key:   key,
		Value: value,
		Proof: proof,
	}

	return testCase
}

// Represents a test case for bedrock's `MerkleTrie.sol`
type trieTestCase struct {
	Root  *[32]byte
	Key   []byte
	Value []byte
	Proof [][]byte
}

// Tuple type to encode `TrieTestCase`
var (
	trieTestCaseTuple, _ = abi.NewType("tuple", "TrieTestCase", []abi.ArgumentMarshaling{
		{Name: "root", Type: "bytes32"},
		{Name: "key", Type: "bytes"},
		{Name: "value", Type: "bytes"},
		{Name: "proof", Type: "bytes[]"},
	})

	encoder = abi.Arguments{
		{Type: trieTestCaseTuple},
	}
)

// Encodes the trieTestCase as the `trieTestCaseTuple`.
func (t *trieTestCase) AbiEncode() string {
	// Encode the contents of the struct as a tuple
	packed, err := encoder.Pack(&t)
	if err != nil {
		log.Fatalf("Error packing TrieTestCase: %v", err)
	}

	// Remove the pointer and encode the packed bytes as a hex string
	return hexutil.Encode(packed[32:])
}

// Helper that generates a cryptographically secure random 64-bit integer
// between the range [min, max]
func randRange(min int64, max int64) int64 {
	r, err := rand.Int(rand.Reader, new(big.Int).Sub(new(big.Int).SetInt64(max), new(big.Int).SetInt64(min)))
	if err != nil {
		log.Fatal("Failed to generate random number within bounds")
	}

	return (new(big.Int).Add(r, new(big.Int).SetInt64(min))).Int64()
}

// Represents a range between two 64-bit integers
type intRange struct {
	min int64
	max int64
}

// Weird golang type coercion wizardry
type proofList [][]byte

func (n *proofList) Put(key []byte, value []byte) error {
	*n = append(*n, value)
	return nil
}

func (n *proofList) Delete(key []byte) error {
	panic("not supported")
}
