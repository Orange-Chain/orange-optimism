// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const L1CrossDomainMessengerStorageLayoutJSON = "{\"storage\":[{\"astId\":26547,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":29735,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initialized\",\"offset\":20,\"slot\":\"0\",\"type\":\"t_uint8\"},{\"astId\":29738,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_initializing\",\"offset\":21,\"slot\":\"0\",\"type\":\"t_bool\"},{\"astId\":30349,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_array(t_uint256)50_storage\"},{\"astId\":29607,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_owner\",\"offset\":0,\"slot\":\"51\",\"type\":\"t_address\"},{\"astId\":29727,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"52\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":29900,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_paused\",\"offset\":0,\"slot\":\"101\",\"type\":\"t_bool\"},{\"astId\":30005,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"102\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":30020,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"_status\",\"offset\":0,\"slot\":\"151\",\"type\":\"t_uint256\"},{\"astId\":30064,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"152\",\"type\":\"t_array(t_uint256)49_storage\"},{\"astId\":26599,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_201_0_32\",\"offset\":0,\"slot\":\"201\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":26604,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"spacer_202_0_32\",\"offset\":0,\"slot\":\"202\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":26609,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"successfulMessages\",\"offset\":0,\"slot\":\"203\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":26612,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"xDomainMsgSender\",\"offset\":0,\"slot\":\"204\",\"type\":\"t_address\"},{\"astId\":26615,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"205\",\"type\":\"t_uint240\"},{\"astId\":26620,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"receivedMessages\",\"offset\":0,\"slot\":\"206\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":26625,\"contract\":\"contracts/L1/L1CrossDomainMessenger.sol:L1CrossDomainMessenger\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"207\",\"type\":\"t_array(t_uint256)42_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)42_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[42]\",\"numberOfBytes\":\"1344\"},\"t_array(t_uint256)49_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[49]\",\"numberOfBytes\":\"1568\"},\"t_array(t_uint256)50_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[50]\",\"numberOfBytes\":\"1600\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint8\":{\"encoding\":\"inplace\",\"label\":\"uint8\",\"numberOfBytes\":\"1\"}}}"

var L1CrossDomainMessengerStorageLayout = new(solc.StorageLayout)

var L1CrossDomainMessengerDeployedBin = "0x6080604052600436106101755760003560e01c80637dea7cc3116100cb578063b28ade251161007f578063ecc7042811610059578063ecc7042814610402578063f2fde38b14610467578063f69f81511461048757600080fd5b8063b28ade25146103af578063c4d66de8146103cf578063d764ad0b146103ef57600080fd5b80638da5cb5b116100b05780638da5cb5b146103205780639fce812c1461034b578063b1b1b2091461037f57600080fd5b80637dea7cc3146102f45780638456cb591461030b57600080fd5b80633f4ba83a1161012d5780635c975abb116101075780635c975abb146102a65780636e296e45146102ca578063715018a6146102df57600080fd5b80633f4ba83a146102475780633f827a5a1461025c57806354fd4d501461028457600080fd5b80630ff754ea1161015e5780630ff754ea146101c35780632828d7e81461021c5780633dbb202b1461023257600080fd5b8063028f85f71461017a5780630c568498146101ad575b600080fd5b34801561018657600080fd5b5061018f601081565b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101b957600080fd5b5061018f6103e881565b3480156101cf57600080fd5b506101f77f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101a4565b34801561022857600080fd5b5061018f6103f881565b610245610240366004611e1f565b6104b7565b005b34801561025357600080fd5b5061024561071b565b34801561026857600080fd5b50610271600181565b60405161ffff90911681526020016101a4565b34801561029057600080fd5b5061029961072d565b6040516101a49190611f00565b3480156102b257600080fd5b5060655460ff165b60405190151581526020016101a4565b3480156102d657600080fd5b506101f76107d0565b3480156102eb57600080fd5b506102456108bc565b34801561030057600080fd5b5061018f62030d4081565b34801561031757600080fd5b506102456108ce565b34801561032c57600080fd5b5060335473ffffffffffffffffffffffffffffffffffffffff166101f7565b34801561035757600080fd5b506101f77f000000000000000000000000000000000000000000000000000000000000000081565b34801561038b57600080fd5b506102ba61039a366004611f1a565b60cb6020526000908152604090205460ff1681565b3480156103bb57600080fd5b5061018f6103ca366004611f33565b6108de565b3480156103db57600080fd5b506102456103ea366004611f87565b61092a565b6102456103fd366004611fa4565b610b31565b34801561040e57600080fd5b5061045960cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b6040519081526020016101a4565b34801561047357600080fd5b50610245610482366004611f87565b6111e1565b34801561049357600080fd5b506102ba6104a2366004611f1a565b60ce6020526000908152604090205460ff1681565b6105f07f00000000000000000000000000000000000000000000000000000000000000006104e68585856108de565b347fd764ad0b0000000000000000000000000000000000000000000000000000000061055260cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b338a34898c8c60405160240161056e9796959493929190612073565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff00000000000000000000000000000000000000000000000000000000909316929092179091526112b4565b8373ffffffffffffffffffffffffffffffffffffffff167fcb0f7ffd78f9aee47a248fae8db181db6eee833039123e026dcbff529522e52a33858561067560cd547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b866040516106879594939291906120d2565b60405180910390a260405134815233907f8ebb2ec2465bdb2a06a66fc37a0963af8a2a6a1479d81d56fdb8cbb98096d5469060200160405180910390a2505060cd80547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff808216600101167fffff0000000000000000000000000000000000000000000000000000000000009091161790555050565b610723611369565b61072b6113ea565b565b60606107587f0000000000000000000000000000000000000000000000000000000000000000611467565b6107817f0000000000000000000000000000000000000000000000000000000000000000611467565b6107aa7f0000000000000000000000000000000000000000000000000000000000000000611467565b6040516020016107bc93929190612120565b604051602081830303815290604052905090565b60cc5460009073ffffffffffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff21530161089f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603560248201527f43726f7373446f6d61696e4d657373656e6765723a2078446f6d61696e4d657360448201527f7361676553656e646572206973206e6f7420736574000000000000000000000060648201526084015b60405180910390fd5b5060cc5473ffffffffffffffffffffffffffffffffffffffff1690565b6108c4611369565b61072b600061159c565b6108d6611369565b61072b611613565b600062030d406108ef6010856121c5565b6103e86109046103f863ffffffff87166121c5565b61090e9190612224565b610918919061224b565b610922919061224b565b949350505050565b6000547501000000000000000000000000000000000000000000900460ff1615808015610975575060005460017401000000000000000000000000000000000000000090910460ff16105b806109a75750303b1580156109a7575060005474010000000000000000000000000000000000000000900460ff166001145b610a33576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201527f647920696e697469616c697a65640000000000000000000000000000000000006064820152608401610896565b600080547fffffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffff16740100000000000000000000000000000000000000001790558015610ab957600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff1675010000000000000000000000000000000000000000001790555b610ac161166e565b610aca8261159c565b8015610b2d57600080547fffffffffffffffffffff00ffffffffffffffffffffffffffffffffffffffffff169055604051600181527f7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb38474024989060200160405180910390a15b5050565b600260975403610b9d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601f60248201527f5265656e7472616e637947756172643a207265656e7472616e742063616c6c006044820152606401610896565b6002609755610baa611765565b60f087901c60018114610c65576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605560248201527f43726f7373446f6d61696e4d657373656e6765723a206f6e6c7920766572736960448201527f6f6e2031206d657373616765732061726520737570706f72746564206166746560648201527f722074686520426564726f636b20757067726164650000000000000000000000608482015260a401610896565b6000610cab898989898989898080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506117d292505050565b9050610cb56117f5565b15610ced57853414610cc957610cc9612277565b600081815260ce602052604090205460ff1615610ce857610ce8612277565b610e3f565b3415610da1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152605060248201527f43726f7373446f6d61696e4d657373656e6765723a2076616c7565206d75737460448201527f206265207a65726f20756e6c657373206d6573736167652069732066726f6d2060648201527f612073797374656d206164647265737300000000000000000000000000000000608482015260a401610896565b600081815260ce602052604090205460ff16610e3f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603060248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520636160448201527f6e6e6f74206265207265706c61796564000000000000000000000000000000006064820152608401610896565b610e4887611919565b15610efb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604360248201527f43726f7373446f6d61696e4d657373656e6765723a2063616e6e6f742073656e60448201527f64206d65737361676520746f20626c6f636b65642073797374656d206164647260648201527f6573730000000000000000000000000000000000000000000000000000000000608482015260a401610896565b600081815260cb602052604090205460ff1615610f9a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f43726f7373446f6d61696e4d657373656e6765723a206d65737361676520686160448201527f7320616c7265616479206265656e2072656c61796564000000000000000000006064820152608401610896565b610fa661afc8866122a6565b5a1015611035576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f43726f7373446f6d61696e4d657373656e6765723a20696e737566666963696560448201527f6e742067617320746f2072656c6179206d6573736167650000000000000000006064820152608401610896565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8a1617905560006110d18861108961138861afc86122be565b5a61109491906122be565b8988888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061199092505050565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead179055905080151560010361116c57600082815260cb602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f4641df4a962071e12719d8c8c8e5ac7fc4d97b927346a3d7a335b1f7517e133c91a26111cb565b600082815260ce602052604080822080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555183917f99d0e048484baa1b1540b1367cb128acd7ab2946d1ed91ec10e3c85e4bf51b8f91a25b505060016097555050505050505050565b905090565b6111e9611369565b73ffffffffffffffffffffffffffffffffffffffff811661128c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201527f64647265737300000000000000000000000000000000000000000000000000006064820152608401610896565b6112958161159c565b50565b73ffffffffffffffffffffffffffffffffffffffff163b151590565b6040517fe9e05c4200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000169063e9e05c429084906113319088908390899060009089906004016122d5565b6000604051808303818588803b15801561134a57600080fd5b505af115801561135e573d6000803e3d6000fd5b505050505050505050565b60335473ffffffffffffffffffffffffffffffffffffffff16331461072b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610896565b6113f26119aa565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001690557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390a1565b6060816000036114aa57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156114d457806114be8161232d565b91506114cd9050600a83612365565b91506114ae565b60008167ffffffffffffffff8111156114ef576114ef612379565b6040519080825280601f01601f191660200182016040528015611519576020820181803683370190505b5090505b84156109225761152e6001836122be565b915061153b600a866123a8565b6115469060306122a6565b60f81b81838151811061155b5761155b6123bc565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350611595600a86612365565b945061151d565b6033805473ffffffffffffffffffffffffffffffffffffffff8381167fffffffffffffffffffffffff0000000000000000000000000000000000000000831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b61161b611765565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25861143d3390565b6000547501000000000000000000000000000000000000000000900460ff16611719576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610896565b60cc80547fffffffffffffffffffffffff00000000000000000000000000000000000000001661dead17905561174d611a16565b611755611ac1565b61175d611b75565b61072b611c4a565b60655460ff161561072b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601060248201527f5061757361626c653a20706175736564000000000000000000000000000000006044820152606401610896565b60006117e2878787878787611cfc565b8051906020012090509695505050505050565b60003373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156111dc57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16639bf62d826040518163ffffffff1660e01b8152600401602060405180830381865afa1580156118d9573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118fd91906123eb565b73ffffffffffffffffffffffffffffffffffffffff1614905090565b600073ffffffffffffffffffffffffffffffffffffffff821630148061198a57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16145b92915050565b600080600080845160208601878a8af19695505050505050565b60655460ff1661072b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601460248201527f5061757361626c653a206e6f74207061757365640000000000000000000000006044820152606401610896565b6000547501000000000000000000000000000000000000000000900460ff1661072b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610896565b6000547501000000000000000000000000000000000000000000900460ff16611b6c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610896565b61072b3361159c565b6000547501000000000000000000000000000000000000000000900460ff16611c20576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610896565b606580547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00169055565b6000547501000000000000000000000000000000000000000000900460ff16611cf5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f496e697469616c697a61626c653a20636f6e7472616374206973206e6f74206960448201527f6e697469616c697a696e670000000000000000000000000000000000000000006064820152608401610896565b6001609755565b6060868686868686604051602401611d1996959493929190612408565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fd764ad0b0000000000000000000000000000000000000000000000000000000017905290509695505050505050565b73ffffffffffffffffffffffffffffffffffffffff8116811461129557600080fd5b60008083601f840112611dcf57600080fd5b50813567ffffffffffffffff811115611de757600080fd5b602083019150836020828501011115611dff57600080fd5b9250929050565b803563ffffffff81168114611e1a57600080fd5b919050565b60008060008060608587031215611e3557600080fd5b8435611e4081611d9b565b9350602085013567ffffffffffffffff811115611e5c57600080fd5b611e6887828801611dbd565b9094509250611e7b905060408601611e06565b905092959194509250565b60005b83811015611ea1578181015183820152602001611e89565b83811115611eb0576000848401525b50505050565b60008151808452611ece816020860160208601611e86565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000611f136020830184611eb6565b9392505050565b600060208284031215611f2c57600080fd5b5035919050565b600080600060408486031215611f4857600080fd5b833567ffffffffffffffff811115611f5f57600080fd5b611f6b86828701611dbd565b9094509250611f7e905060208501611e06565b90509250925092565b600060208284031215611f9957600080fd5b8135611f1381611d9b565b600080600080600080600060c0888a031215611fbf57600080fd5b873596506020880135611fd181611d9b565b95506040880135611fe181611d9b565b9450606088013593506080880135925060a088013567ffffffffffffffff81111561200b57600080fd5b6120178a828b01611dbd565b989b979a50959850939692959293505050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b878152600073ffffffffffffffffffffffffffffffffffffffff808916602084015280881660408401525085606083015263ffffffff8516608083015260c060a08301526120c560c08301848661202a565b9998505050505050505050565b73ffffffffffffffffffffffffffffffffffffffff8616815260806020820152600061210260808301868861202a565b905083604083015263ffffffff831660608301529695505050505050565b60008451612132818460208901611e86565b80830190507f2e00000000000000000000000000000000000000000000000000000000000000808252855161216e816001850160208a01611e86565b60019201918201528351612189816002840160208801611e86565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600067ffffffffffffffff808316818516818304811182151516156121ec576121ec612196565b02949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600067ffffffffffffffff8084168061223f5761223f6121f5565b92169190910492915050565b600067ffffffffffffffff80831681851680830382111561226e5761226e612196565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052600160045260246000fd5b600082198211156122b9576122b9612196565b500190565b6000828210156122d0576122d0612196565b500390565b73ffffffffffffffffffffffffffffffffffffffff8616815284602082015267ffffffffffffffff84166040820152821515606082015260a06080820152600061232260a0830184611eb6565b979650505050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361235e5761235e612196565b5060010190565b600082612374576123746121f5565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000826123b7576123b76121f5565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000602082840312156123fd57600080fd5b8151611f1381611d9b565b868152600073ffffffffffffffffffffffffffffffffffffffff808816602084015280871660408401525084606083015283608083015260c060a083015261245360c0830184611eb6565b9897505050505050505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L1CrossDomainMessengerStorageLayoutJSON), L1CrossDomainMessengerStorageLayout); err != nil {
		panic(err)
	}

	layouts["L1CrossDomainMessenger"] = L1CrossDomainMessengerStorageLayout
	deployedBytecodes["L1CrossDomainMessenger"] = L1CrossDomainMessengerDeployedBin
}
