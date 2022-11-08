// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const ProxyAdminStorageLayoutJSON = "{\"storage\":[{\"astId\":37808,\"contract\":\"contracts/universal/ProxyAdmin.sol:ProxyAdmin\",\"label\":\"owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":28370,\"contract\":\"contracts/universal/ProxyAdmin.sol:ProxyAdmin\",\"label\":\"proxyType\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_enum(ProxyType)28364)\"},{\"astId\":28375,\"contract\":\"contracts/universal/ProxyAdmin.sol:ProxyAdmin\",\"label\":\"implementationName\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_mapping(t_address,t_string_storage)\"},{\"astId\":28379,\"contract\":\"contracts/universal/ProxyAdmin.sol:ProxyAdmin\",\"label\":\"addressManager\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_contract(AddressManager)4425\"},{\"astId\":28383,\"contract\":\"contracts/universal/ProxyAdmin.sol:ProxyAdmin\",\"label\":\"upgrading\",\"offset\":20,\"slot\":\"3\",\"type\":\"t_bool\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(AddressManager)4425\":{\"encoding\":\"inplace\",\"label\":\"contract AddressManager\",\"numberOfBytes\":\"20\"},\"t_enum(ProxyType)28364\":{\"encoding\":\"inplace\",\"label\":\"enum ProxyAdmin.ProxyType\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_enum(ProxyType)28364)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e enum ProxyAdmin.ProxyType)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_enum(ProxyType)28364\"},\"t_mapping(t_address,t_string_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e string)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_string_storage\"},\"t_string_storage\":{\"encoding\":\"bytes\",\"label\":\"string\",\"numberOfBytes\":\"32\"}}}"

var ProxyAdminStorageLayout = new(solc.StorageLayout)

var ProxyAdminDeployedBin = "0x6080604052600436106100f35760003560e01c8063860f7cda1161008a57806399a88ec41161005957806399a88ec4146102db5780639b2ea4bd146102fb578063b79472621461031b578063f3b7dead1461035657600080fd5b8063860f7cda1461025b5780638d52d4a01461027b5780638da5cb5b1461029b5780639623609d146102c857600080fd5b8063238181ae116100c6578063238181ae146101a45780633ab76e9f146101d15780636bd9f516146101fe5780637eff275e1461023b57600080fd5b80630652b57a146100f857806307c8f7b01461011a57806313af40351461013a578063204e1c7a1461015a575b600080fd5b34801561010457600080fd5b506101186101133660046114c6565b610376565b005b34801561012657600080fd5b506101186101353660046114e3565b610443565b34801561014657600080fd5b506101186101553660046114c6565b61050e565b34801561016657600080fd5b5061017a6101753660046114c6565b6105ff565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156101b057600080fd5b506101c46101bf3660046114c6565b610820565b60405161019b919061157b565b3480156101dd57600080fd5b5060035461017a9073ffffffffffffffffffffffffffffffffffffffff1681565b34801561020a57600080fd5b5061022e6102193660046114c6565b60016020526000908152604090205460ff1681565b60405161019b91906115bd565b34801561024757600080fd5b506101186102563660046115fe565b6108ba565b34801561026757600080fd5b50610118610276366004611759565b610ae6565b34801561028757600080fd5b506101186102963660046117a9565b610b96565b3480156102a757600080fd5b5060005461017a9073ffffffffffffffffffffffffffffffffffffffff1681565b6101186102d63660046117db565b610c83565b3480156102e757600080fd5b506101186102f63660046115fe565b610f13565b34801561030757600080fd5b50610118610316366004611851565b61121c565b34801561032757600080fd5b5060035474010000000000000000000000000000000000000000900460ff16604051901515815260200161019b565b34801561036257600080fd5b5061017a6103713660046114c6565b61132b565b60005473ffffffffffffffffffffffffffffffffffffffff1633146103fc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f554e415554484f52495a4544000000000000000000000000000000000000000060448201526064015b60405180910390fd5b600380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff1633146104c4576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f554e415554484f52495a4544000000000000000000000000000000000000000060448201526064016103f3565b6003805491151574010000000000000000000000000000000000000000027fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff909216919091179055565b60005473ffffffffffffffffffffffffffffffffffffffff16331461058f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f554e415554484f52495a4544000000000000000000000000000000000000000060448201526064016103f3565b600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff83169081178255604051909133917f8292fce18fa69edf4db7b94ea2e58241df0ae57f97e0a6c9b29067028bf92d769190a350565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001602052604081205460ff168181600281111561063b5761063b61158e565b036106b6578273ffffffffffffffffffffffffffffffffffffffff16635c60da1b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561068b573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106af9190611898565b9392505050565b60018160028111156106ca576106ca61158e565b0361071a578273ffffffffffffffffffffffffffffffffffffffff1663aaf10f426040518163ffffffff1660e01b8152600401602060405180830381865afa15801561068b573d6000803e3d6000fd5b600281600281111561072e5761072e61158e565b036107b85760035473ffffffffffffffffffffffffffffffffffffffff8481166000908152600260205260409081902090517fbf40fac1000000000000000000000000000000000000000000000000000000008152919092169163bf40fac19161079b9190600401611902565b602060405180830381865afa15801561068b573d6000803e3d6000fd5b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601e60248201527f50726f787941646d696e3a20756e6b6e6f776e2070726f78792074797065000060448201526064016103f3565b50919050565b60026020526000908152604090208054610839906118b5565b80601f0160208091040260200160405190810160405280929190818152602001828054610865906118b5565b80156108b25780601f10610887576101008083540402835291602001916108b2565b820191906000526020600020905b81548152906001019060200180831161089557829003601f168201915b505050505081565b60005473ffffffffffffffffffffffffffffffffffffffff16331461093b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f554e415554484f52495a4544000000000000000000000000000000000000000060448201526064016103f3565b73ffffffffffffffffffffffffffffffffffffffff821660009081526001602052604081205460ff16908160028111156109775761097761158e565b03610a03576040517f8f28397000000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152841690638f283970906024015b600060405180830381600087803b1580156109e657600080fd5b505af11580156109fa573d6000803e3d6000fd5b50505050505050565b6001816002811115610a1757610a1761158e565b03610a70576040517f13af403500000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff83811660048301528416906313af4035906024016109cc565b6002816002811115610a8457610a8461158e565b036107b8576003546040517ff2fde38b00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529091169063f2fde38b906024016109cc565b505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610b67576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f554e415554484f52495a4544000000000000000000000000000000000000000060448201526064016103f3565b73ffffffffffffffffffffffffffffffffffffffff82166000908152600260205260409020610ae182826119f1565b60005473ffffffffffffffffffffffffffffffffffffffff163314610c17576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f554e415554484f52495a4544000000000000000000000000000000000000000060448201526064016103f3565b73ffffffffffffffffffffffffffffffffffffffff82166000908152600160208190526040909120805483927fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0090911690836002811115610c7a57610c7a61158e565b02179055505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610d04576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f554e415554484f52495a4544000000000000000000000000000000000000000060448201526064016103f3565b73ffffffffffffffffffffffffffffffffffffffff831660009081526001602052604081205460ff1690816002811115610d4057610d4061158e565b03610e06576040517f4f1ef28600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff851690634f1ef286903490610d9b9087908790600401611b0b565b60006040518083038185885af1158015610db9573d6000803e3d6000fd5b50505050506040513d6000823e601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201604052610e009190810190611b42565b50610f0d565b610e108484610f13565b60008473ffffffffffffffffffffffffffffffffffffffff163484604051610e389190611bb9565b60006040518083038185875af1925050503d8060008114610e75576040519150601f19603f3d011682016040523d82523d6000602084013e610e7a565b606091505b5050905080610f0b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f50726f787941646d696e3a2063616c6c20746f2070726f78792061667465722060448201527f75706772616465206661696c656400000000000000000000000000000000000060648201526084016103f3565b505b50505050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610f94576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f554e415554484f52495a4544000000000000000000000000000000000000000060448201526064016103f3565b73ffffffffffffffffffffffffffffffffffffffff821660009081526001602052604081205460ff1690816002811115610fd057610fd061158e565b03611029576040517f3659cfe600000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8381166004830152841690633659cfe6906024016109cc565b600181600281111561103d5761103d61158e565b036110bc576040517f9b0b0fda0000000000000000000000000000000000000000000000000000000081527f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc600482015273ffffffffffffffffffffffffffffffffffffffff8381166024830152841690639b0b0fda906044016109cc565b60028160028111156110d0576110d061158e565b036112145773ffffffffffffffffffffffffffffffffffffffff831660009081526002602052604081208054611105906118b5565b80601f0160208091040260200160405190810160405280929190818152602001828054611131906118b5565b801561117e5780601f106111535761010080835404028352916020019161117e565b820191906000526020600020905b81548152906001019060200180831161116157829003601f168201915b50506003546040517f9b2ea4bd00000000000000000000000000000000000000000000000000000000815294955073ffffffffffffffffffffffffffffffffffffffff1693639b2ea4bd93506111dc92508591508790600401611bd5565b600060405180830381600087803b1580156111f657600080fd5b505af115801561120a573d6000803e3d6000fd5b5050505050505050565b610ae1611c0d565b60005473ffffffffffffffffffffffffffffffffffffffff16331461129d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152600c60248201527f554e415554484f52495a4544000000000000000000000000000000000000000060448201526064016103f3565b6003546040517f9b2ea4bd00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff90911690639b2ea4bd906112f59085908590600401611bd5565b600060405180830381600087803b15801561130f57600080fd5b505af1158015611323573d6000803e3d6000fd5b505050505050565b73ffffffffffffffffffffffffffffffffffffffff811660009081526001602052604081205460ff16818160028111156113675761136761158e565b036113b7578273ffffffffffffffffffffffffffffffffffffffff1663f851a4406040518163ffffffff1660e01b8152600401602060405180830381865afa15801561068b573d6000803e3d6000fd5b60018160028111156113cb576113cb61158e565b0361141b578273ffffffffffffffffffffffffffffffffffffffff1663893d20e86040518163ffffffff1660e01b8152600401602060405180830381865afa15801561068b573d6000803e3d6000fd5b600281600281111561142f5761142f61158e565b036107b857600360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638da5cb5b6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561068b573d6000803e3d6000fd5b73ffffffffffffffffffffffffffffffffffffffff811681146114c357600080fd5b50565b6000602082840312156114d857600080fd5b81356106af816114a1565b6000602082840312156114f557600080fd5b813580151581146106af57600080fd5b60005b83811015611520578181015183820152602001611508565b83811115610f0d5750506000910152565b60008151808452611549816020860160208601611505565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006106af6020830184611531565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60208101600383106115f8577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b6000806040838503121561161157600080fd5b823561161c816114a1565b9150602083013561162c816114a1565b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016810167ffffffffffffffff811182821017156116ad576116ad611637565b604052919050565b600067ffffffffffffffff8211156116cf576116cf611637565b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01660200190565b600061170e611709846116b5565b611666565b905082815283838301111561172257600080fd5b828260208301376000602084830101529392505050565b600082601f83011261174a57600080fd5b6106af838335602085016116fb565b6000806040838503121561176c57600080fd5b8235611777816114a1565b9150602083013567ffffffffffffffff81111561179357600080fd5b61179f85828601611739565b9150509250929050565b600080604083850312156117bc57600080fd5b82356117c7816114a1565b915060208301356003811061162c57600080fd5b6000806000606084860312156117f057600080fd5b83356117fb816114a1565b9250602084013561180b816114a1565b9150604084013567ffffffffffffffff81111561182757600080fd5b8401601f8101861361183857600080fd5b611847868235602084016116fb565b9150509250925092565b6000806040838503121561186457600080fd5b823567ffffffffffffffff81111561187b57600080fd5b61188785828601611739565b925050602083013561162c816114a1565b6000602082840312156118aa57600080fd5b81516106af816114a1565b600181811c908216806118c957607f821691505b60208210810361081a577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b6000602080835260008454611916816118b5565b80848701526040600180841660008114611937576001811461196f5761199d565b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff008516838a01528284151560051b8a0101955061199d565b896000528660002060005b858110156119955781548b820186015290830190880161197a565b8a0184019650505b509398975050505050505050565b601f821115610ae157600081815260208120601f850160051c810160208610156119d25750805b601f850160051c820191505b81811015611323578281556001016119de565b815167ffffffffffffffff811115611a0b57611a0b611637565b611a1f81611a1984546118b5565b846119ab565b602080601f831160018114611a725760008415611a3c5750858301515b7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600386901b1c1916600185901b178555611323565b6000858152602081207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08616915b82811015611abf57888601518255948401946001909101908401611aa0565b5085821015611afb57878501517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff600388901b60f8161c191681555b5050505050600190811b01905550565b73ffffffffffffffffffffffffffffffffffffffff83168152604060208201526000611b3a6040830184611531565b949350505050565b600060208284031215611b5457600080fd5b815167ffffffffffffffff811115611b6b57600080fd5b8201601f81018413611b7c57600080fd5b8051611b8a611709826116b5565b818152856020838501011115611b9f57600080fd5b611bb0826020830160208601611505565b95945050505050565b60008251611bcb818460208701611505565b9190910192915050565b604081526000611be86040830185611531565b905073ffffffffffffffffffffffffffffffffffffffff831660208301529392505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(ProxyAdminStorageLayoutJSON), ProxyAdminStorageLayout); err != nil {
		panic(err)
	}

	layouts["ProxyAdmin"] = ProxyAdminStorageLayout
	deployedBytecodes["ProxyAdmin"] = ProxyAdminDeployedBin
}
