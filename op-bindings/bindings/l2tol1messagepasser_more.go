// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L2ToL1MessagePasserStorageLayoutJSON = "{\"storage\":[{\"astId\":2543,\"contract\":\"contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser\",\"label\":\"sentMessages\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":2546,\"contract\":\"contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser\",\"label\":\"nonce\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint256\"}],\"types\":{\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L2ToL1MessagePasserStorageLayout = new(solc.StorageLayout)

var L2ToL1MessagePasserDeployedBin = "0x60806040526004361061005e5760003560e01c806382e3702d1161004357806382e3702d146100c7578063affed0e014610107578063c2b3e5ac1461012b57600080fd5b806344df8e701461008757806354fd4d501461009c57600080fd5b366100825761008033620186a060405180602001604052806000815250610139565b005b600080fd5b34801561009357600080fd5b5061008061026d565b3480156100a857600080fd5b506100b16102a5565b6040516100be9190610587565b60405180910390f35b3480156100d357600080fd5b506100f76100e23660046105a1565b60006020819052908152604090205460ff1681565b60405190151581526020016100be565b34801561011357600080fd5b5061011d60015481565b6040519081526020016100be565b6100806101393660046105e9565b600061019e6040518060c0016040528060015481526020013373ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff16815260200134815260200185815260200184815250610348565b6000818152602081905260409081902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600190811790915554905191925073ffffffffffffffffffffffffffffffffffffffff8616913391907f87bf7b546c8de873abb0db5b579ec131f8d0cf5b14f39933551cf9ced23a61369061022c903490899089906106ed565b60405180910390a460405181907f2ef6ceb1668fdd882b1f89ddd53a666b0c1113d14cf90c0fbf97c7b1ad880fbb90600090a2505060018054810190555050565b4761027781610395565b60405181907f7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f90600090a250565b60606102d07f00000000000000000000000000000000000000000000000000000000000000006103c4565b6102f97f00000000000000000000000000000000000000000000000000000000000000006103c4565b6103227f00000000000000000000000000000000000000000000000000000000000000006103c4565b60405160200161033493929190610715565b604051602081830303815290604052905090565b80516020808301516040808501516060860151608087015160a0880151935160009761037897909695910161078b565b604051602081830303815290604052805190602001209050919050565b806040516103a290610501565b6040518091039082f09050801580156103bf573d6000803e3d6000fd5b505050565b60608160000361040757505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610431578061041b81610811565b915061042a9050600a83610878565b915061040b565b60008167ffffffffffffffff81111561044c5761044c6105ba565b6040519080825280601f01601f191660200182016040528015610476576020820181803683370190505b5090505b84156104f95761048b60018361088c565b9150610498600a866108a3565b6104a39060306108b7565b60f81b8183815181106104b8576104b86108cf565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506104f2600a86610878565b945061047a565b949350505050565b6008806108ff83390190565b60005b83811015610528578181015183820152602001610510565b83811115610537576000848401525b50505050565b6000815180845261055581602086016020860161050d565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061059a602083018461053d565b9392505050565b6000602082840312156105b357600080fd5b5035919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806000606084860312156105fe57600080fd5b833573ffffffffffffffffffffffffffffffffffffffff8116811461062257600080fd5b925060208401359150604084013567ffffffffffffffff8082111561064657600080fd5b818601915086601f83011261065a57600080fd5b81358181111561066c5761066c6105ba565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156106b2576106b26105ba565b816040528281528960208487010111156106cb57600080fd5b8260208601602083013760006020848301015280955050505050509250925092565b83815282602082015260606040820152600061070c606083018461053d565b95945050505050565b6000845161072781846020890161050d565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610763816001850160208a0161050d565b6001920191820152835161077e81600284016020880161050d565b0160020195945050505050565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a08301526107d660c083018461053d565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610842576108426107e2565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b60008261088757610887610849565b500490565b60008282101561089e5761089e6107e2565b500390565b6000826108b2576108b2610849565b500690565b600082198211156108ca576108ca6107e2565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfe608060405230fffea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2ToL1MessagePasserStorageLayoutJSON), L2ToL1MessagePasserStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ToL1MessagePasser"] = L2ToL1MessagePasserStorageLayout
	deployedBytecodes["L2ToL1MessagePasser"] = L2ToL1MessagePasserDeployedBin
}
