// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const SequencerFeeVaultStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/SequencerFeeVault.sol:SequencerFeeVault\",\"label\":\"totalProcessed\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_uint256\"}],\"types\":{\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var SequencerFeeVaultStorageLayout = new(solc.StorageLayout)

var SequencerFeeVaultDeployedBin = "0x6080604052600436106100695760003560e01c806384411d651161004357806384411d651461010c578063d3e5792b14610130578063d4ff92181461016457600080fd5b80630d9019e1146100755780633ccfd60b146100d357806354fd4d50146100ea57600080fd5b3661007057005b600080fd5b34801561008157600080fd5b506100a97f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156100df57600080fd5b506100e8610197565b005b3480156100f657600080fd5b506100ff6103b8565b6040516100ca9190610612565b34801561011857600080fd5b5061012260005481565b6040519081526020016100ca565b34801561013c57600080fd5b506101227f000000000000000000000000000000000000000000000000000000000000000081565b34801561017057600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006100a9565b7f0000000000000000000000000000000000000000000000000000000000000000471015610271576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f4665655661756c743a207769746864726177616c20616d6f756e74206d75737460448201527f2062652067726561746572207468616e206d696e696d756d207769746864726160648201527f77616c20616d6f756e7400000000000000000000000000000000000000000000608482015260a40160405180910390fd5b600047905080600080828254610287919061065b565b9091555050604080518281527f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff166020820152338183015290517fc8a211cc64b6ed1b50595a9fcb1932b6d1e5a6e8ef15b60e5b1f988ea9086bba9181900360600190a1604080516020810182526000815290517fe11013dd0000000000000000000000000000000000000000000000000000000081527342000000000000000000000000000000000000109163e11013dd918491610383917f0000000000000000000000000000000000000000000000000000000000000000916188b891600401610673565b6000604051808303818588803b15801561039c57600080fd5b505af11580156103b0573d6000803e3d6000fd5b505050505050565b60606103e37f000000000000000000000000000000000000000000000000000000000000000061045b565b61040c7f000000000000000000000000000000000000000000000000000000000000000061045b565b6104357f000000000000000000000000000000000000000000000000000000000000000061045b565b604051602001610447939291906106b7565b604051602081830303815290604052905090565b60608160000361049e57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156104c857806104b28161072d565b91506104c19050600a83610794565b91506104a2565b60008167ffffffffffffffff8111156104e3576104e36107a8565b6040519080825280601f01601f19166020018201604052801561050d576020820181803683370190505b5090505b8415610590576105226001836107d7565b915061052f600a866107ee565b61053a90603061065b565b60f81b81838151811061054f5761054f610802565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610589600a86610794565b9450610511565b949350505050565b60005b838110156105b357818101518382015260200161059b565b838111156105c2576000848401525b50505050565b600081518084526105e0816020860160208601610598565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061062560208301846105c8565b9392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561066e5761066e61062c565b500190565b73ffffffffffffffffffffffffffffffffffffffff8416815263ffffffff831660208201526060604082015260006106ae60608301846105c8565b95945050505050565b600084516106c9818460208901610598565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551610705816001850160208a01610598565b60019201918201528351610720816002840160208801610598565b0160020195945050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361075e5761075e61062c565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826107a3576107a3610765565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000828210156107e9576107e961062c565b500390565b6000826107fd576107fd610765565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(SequencerFeeVaultStorageLayoutJSON), SequencerFeeVaultStorageLayout); err != nil {
		panic(err)
	}

	layouts["SequencerFeeVault"] = SequencerFeeVaultStorageLayout
	deployedBytecodes["SequencerFeeVault"] = SequencerFeeVaultDeployedBin
}
