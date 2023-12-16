// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const GasPriceOracleStorageLayoutJSON = "{\"storage\":null,\"types\":{}}"

var GasPriceOracleStorageLayout = new(solc.StorageLayout)

var GasPriceOracleDeployedBin = "0x608060405234801561001057600080fd5b50600436106100be5760003560e01c806354fd4d5011610076578063de26c4a11161005b578063de26c4a114610157578063f45e65d81461016a578063fe173b971461015157600080fd5b806354fd4d50146101085780636ef25c3a1461015157600080fd5b8063313ce567116100a7578063313ce567146100e657806349948e0e146100ed578063519b4bd31461010057600080fd5b80630c18c162146100c35780632e0f2625146100de575b600080fd5b6100cb610172565b6040519081526020015b60405180910390f35b6100cb600681565b60066100cb565b6100cb6100fb3660046105b9565b6101fc565b6100cb61025d565b6101446040518060400160405280600581526020017f312e312e3000000000000000000000000000000000000000000000000000000081525081565b6040516100d59190610688565b486100cb565b6100cb6101653660046105b9565b6102be565b6100cb610301565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101f791906106fb565b905090565b600080610208836102be565b9050600061021461025d565b61021e9083610743565b9050600061022e6006600a6108a0565b9050600061023a610301565b6102449084610743565b9050600061025283836108ac565b979650505050505050565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16635cf249696040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101d3573d6000803e3d6000fd5b6000806102ca83610362565b6102d5906010610743565b905060006102e1610172565b6102eb90836108e7565b90506102f9816104406108e7565b949350505050565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16639e8c49666040518163ffffffff1660e01b8152600401602060405180830381865afa1580156101d3573d6000803e3d6000fd5b600061044c565b600082840393505b838110156103925782810151828201511860001a1590930292600101610371565b9392505050565b601f81169060209004602102820181156103b35781016001015b92915050565b600181039050610106810460030282019150600060066101068306106103e35750600382016103b3565b505060020190565b61044482820361042861041884600081518060001a8160011a60081b178160021a60101b17915050919050565b639e3779b90260131c611fff1690565b8060021b6040510182815160e01c1860e01b8151188152505050565b600101919050565b6180003860405139600090506020820180600d8451820103826002015b81811015610579576000805b50508051604051600082901a600183901a60081b1760029290921a60101b91909117639e3779b9810260111c617ffc16909101805160e081811c878603811890911b909118909152840190818303908484106104d1575061050c565b600184019350611fff8211610506578251600081901a600182901a60081b1760029190911a60101b178103610506575061050c565b50610475565b5082821061051a5750610579565b600182039150848211156105375761053486868403610399565b95505b61054b600984016003840160038401610369565b905061055786826103b9565b955061056e84610569868486016103eb565b6103eb565b915050809350610469565b50506102f983838651840103610399565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156105cb57600080fd5b813567ffffffffffffffff808211156105e357600080fd5b818401915084601f8301126105f757600080fd5b8135818111156106095761060961058a565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561064f5761064f61058a565b8160405282815287602084870101111561066857600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156106b557858101830151858201604001528201610699565b818111156106c7576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b60006020828403121561070d57600080fd5b5051919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048311821515161561077b5761077b610714565b500290565b600181815b808511156107d957817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156107bf576107bf610714565b808516156107cc57918102915b93841c9390800290610785565b509250929050565b6000826107f0575060016103b3565b816107fd575060006103b3565b8160018114610813576002811461081d57610839565b60019150506103b3565b60ff84111561082e5761082e610714565b50506001821b6103b3565b5060208310610133831016604e8410600b841016171561085c575081810a6103b3565b6108668383610780565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0482111561089857610898610714565b029392505050565b600061039283836107e1565b6000826108e2577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600082198211156108fa576108fa610714565b50019056fea164736f6c634300080f000a"


func init() {
	if err := json.Unmarshal([]byte(GasPriceOracleStorageLayoutJSON), GasPriceOracleStorageLayout); err != nil {
		panic(err)
	}

	layouts["GasPriceOracle"] = GasPriceOracleStorageLayout
	deployedBytecodes["GasPriceOracle"] = GasPriceOracleDeployedBin
	immutableReferences["GasPriceOracle"] = false
}
