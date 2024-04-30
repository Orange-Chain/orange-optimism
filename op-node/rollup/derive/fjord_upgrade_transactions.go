package derive

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum-optimism/optimism/op-bindings/predeploys"
)

var (
	// Gas Price Oracle Parameters
	deployFjordGasPriceOracleSource       = UpgradeDepositSource{Intent: "Fjord: Gas Price Oracle Deployment"}
	GasPriceOracleFjordDeployerAddress    = common.HexToAddress("0x4210000000000000000000000000000000000002")
	gasPriceOracleFjordDeploymentBytecode = common.FromHex("0x608060405234801561001057600080fd5b50611930806100206000396000f3fe608060405234801561001057600080fd5b50600436106101365760003560e01c80636ef25c3a116100b2578063de26c4a111610081578063f45e65d811610066578063f45e65d81461025b578063f820614014610263578063fe173b971461020d57600080fd5b8063de26c4a114610235578063f1c7a58b1461024857600080fd5b80636ef25c3a1461020d5780638e98b10614610213578063960e3a231461021b578063c59859181461022d57600080fd5b806349948e0e11610109578063519b4bd3116100ee578063519b4bd31461019f57806354fd4d50146101a757806368d5dca6146101f057600080fd5b806349948e0e1461016f5780634ef6e2241461018257600080fd5b80630c18c1621461013b57806322b90ab3146101565780632e0f262514610160578063313ce56714610168575b600080fd5b61014361026b565b6040519081526020015b60405180910390f35b61015e61038c565b005b610143600681565b6006610143565b61014361017d3660046113a5565b6105af565b60005461018f9060ff1681565b604051901515815260200161014d565b610143610600565b6101e36040518060400160405280600581526020017f312e332e3000000000000000000000000000000000000000000000000000000081525081565b60405161014d9190611474565b6101f8610661565b60405163ffffffff909116815260200161014d565b48610143565b61015e6106e6565b60005461018f90610100900460ff1681565b6101f861097a565b6101436102433660046113a5565b6109db565b6101436102563660046114e7565b610abf565b610143610b8f565b610143610c82565b6000805460ff1615610304576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f47617350726963654f7261636c653a206f76657268656164282920697320646560448201527f707265636174656400000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa158015610363573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103879190611500565b905090565b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663e591b2826040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103eb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061040f9190611519565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104ef576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f47617350726963654f7261636c653a206f6e6c7920746865206465706f73697460448201527f6f72206163636f756e742063616e2073657420697345636f746f6e6520666c6160648201527f6700000000000000000000000000000000000000000000000000000000000000608482015260a4016102fb565b60005460ff1615610582576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f47617350726963654f7261636c653a2045636f746f6e6520616c72656164792060448201527f616374697665000000000000000000000000000000000000000000000000000060648201526084016102fb565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055565b60008054610100900460ff16156105e3576105dd6105cc83610ce3565b516105d890604461157e565b611008565b92915050565b60005460ff16156105f7576105dd826110ee565b6105dd82611192565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16635cf249696040518163ffffffff1660e01b8152600401602060405180830381865afa158015610363573d6000803e3d6000fd5b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff166368d5dca66040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106c2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103879190611596565b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663e591b2826040518163ffffffff1660e01b8152600401602060405180830381865afa158015610745573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107699190611519565b73ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610823576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603f60248201527f47617350726963654f7261636c653a206f6e6c7920746865206465706f73697460448201527f6f72206163636f756e742063616e20736574206973466a6f726420666c61670060648201526084016102fb565b60005460ff166108b5576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603960248201527f47617350726963654f7261636c653a20466a6f72642063616e206f6e6c79206260448201527f65206163746976617465642061667465722045636f746f6e650000000000000060648201526084016102fb565b600054610100900460ff161561094c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f47617350726963654f7261636c653a20466a6f726420616c726561647920616360448201527f746976650000000000000000000000000000000000000000000000000000000060648201526084016102fb565b600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff16610100179055565b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663c59859186040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106c2573d6000803e3d6000fd5b60008054610100900460ff1615610a0c576109f582610ce3565b51610a0190604461157e565b6105dd9060106115bc565b6000610a17836112e6565b60005490915060ff1615610a2b5792915050565b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a8a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aae9190611500565b610ab8908261157e565b9392505050565b60008054610100900460ff16610b57576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603660248201527f47617350726963654f7261636c653a206765744c314665655570706572426f7560448201527f6e64206f6e6c7920737570706f72747320466a6f72640000000000000000000060648201526084016102fb565b6000610b6460ff846115f9565b610b6e908461157e565b610b7990601061157e565b610b8490604461157e565b9050610ab881611008565b6000805460ff1615610c23576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f47617350726963654f7261636c653a207363616c61722829206973206465707260448201527f656361746564000000000000000000000000000000000000000000000000000060648201526084016102fb565b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16639e8c49666040518163ffffffff1660e01b8152600401602060405180830381865afa158015610363573d6000803e3d6000fd5b600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff1663f82061406040518163ffffffff1660e01b8152600401602060405180830381865afa158015610363573d6000803e3d6000fd5b6060610e7a565b818153600101919050565b600082840393505b83811015610ab85782810151828201511860001a1590930292600101610cfd565b825b60208210610d6a578251610d35601f83610cea565b52602092909201917fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe090910190602101610d20565b8115610ab8578251610d7f6001840383610cea565b520160010192915050565b60006001830392505b6101078210610dcb57610dbd8360ff16610db860fd610db88760081c60e00189610cea565b610cea565b935061010682039150610d93565b60078210610df857610df18360ff16610db860078503610db88760081c60e00189610cea565b9050610ab8565b610e118360ff16610db88560081c8560051b0187610cea565b949350505050565b610e72828203610e56610e4684600081518060001a8160011a60081b178160021a60101b17915050919050565b639e3779b90260131c611fff1690565b8060021b6040510182815160e01c1860e01b8151188152505050565b600101919050565b6180003860405139618000604051016020830180600d8551820103826002015b81811015610fad576000805b50508051604051600082901a600183901a60081b1760029290921a60101b91909117639e3779b9810260111c617ffc16909101805160e081811c878603811890911b90911890915284019081830390848410610f025750610f3d565b600184019350611fff8211610f37578251600081901a600182901a60081b1760029190911a60101b178103610f375750610f3d565b50610ea6565b838310610f4b575050610fad565b60018303925085831115610f6957610f668787888603610d1e565b96505b610f7d600985016003850160038501610cf5565b9150610f8a878284610d8a565b965050610fa284610f9d86848601610e19565b610e19565b915050809350610e9a565b5050610fbf8383848851850103610d1e565b925050506040519150618000820180820391508183526020830160005b83811015610ff4578281015182820152602001610fdc565b506000920191825250602001604052919050565b600080611013610c82565b61101b610661565b63ffffffff1661102b91906115bc565b611033610600565b61103b61097a565b611046906010611634565b63ffffffff1661105691906115bc565b611060919061157e565b9050600061107184620cc3946115bc565b61109b907ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd763200611660565b90506110ab6064620f42406116d4565b8112156110c3576110c06064620f42406116d4565b90505b6110cf600660026115bc565b6110da90600a6118b0565b6110e483836115bc565b610e1191906115f9565b6000806110fa836112e6565b90506000611106610600565b61110e61097a565b611119906010611634565b63ffffffff1661112991906115bc565b90506000611135610c82565b61113d610661565b63ffffffff1661114d91906115bc565b9050600061115b828461157e565b61116590856115bc565b90506111736006600a6118b0565b61117e9060106115bc565b61118890826115f9565b9695505050505050565b60008061119e836112e6565b9050600073420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16639e8c49666040518163ffffffff1660e01b8152600401602060405180830381865afa158015611201573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112259190611500565b61122d610600565b73420000000000000000000000000000000000001573ffffffffffffffffffffffffffffffffffffffff16638b239f736040518163ffffffff1660e01b8152600401602060405180830381865afa15801561128c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112b09190611500565b6112ba908561157e565b6112c491906115bc565b6112ce91906115bc565b90506112dc6006600a6118b0565b610e1190826115f9565b80516000908190815b8181101561136957848181518110611309576113096118bc565b01602001517fff00000000000000000000000000000000000000000000000000000000000000166000036113495761134260048461157e565b9250611357565b61135460108461157e565b92505b80611361816118eb565b9150506112ef565b50610e118261044061157e565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000602082840312156113b757600080fd5b813567ffffffffffffffff808211156113cf57600080fd5b818401915084601f8301126113e357600080fd5b8135818111156113f5576113f5611376565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f0116810190838211818310171561143b5761143b611376565b8160405282815287602084870101111561145457600080fd5b826020860160208301376000928101602001929092525095945050505050565b600060208083528351808285015260005b818110156114a157858101830151858201604001528201611485565b818111156114b3576000604083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe016929092016040019392505050565b6000602082840312156114f957600080fd5b5035919050565b60006020828403121561151257600080fd5b5051919050565b60006020828403121561152b57600080fd5b815173ffffffffffffffffffffffffffffffffffffffff81168114610ab857600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082198211156115915761159161154f565b500190565b6000602082840312156115a857600080fd5b815163ffffffff81168114610ab857600080fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff04831182151516156115f4576115f461154f565b500290565b60008261162f577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b500490565b600063ffffffff808316818516818304811182151516156116575761165761154f565b02949350505050565b6000808212827f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0384138115161561169a5761169a61154f565b827f80000000000000000000000000000000000000000000000000000000000000000384128116156116ce576116ce61154f565b50500190565b60007f7fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff6000841360008413858304851182821616156117155761171561154f565b7f800000000000000000000000000000000000000000000000000000000000000060008712868205881281841616156117505761175061154f565b6000871292508782058712848416161561176c5761176c61154f565b878505871281841616156117825761178261154f565b505050929093029392505050565b600181815b808511156117e957817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156117cf576117cf61154f565b808516156117dc57918102915b93841c9390800290611795565b509250929050565b600082611800575060016105dd565b8161180d575060006105dd565b8160018114611823576002811461182d57611849565b60019150506105dd565b60ff84111561183e5761183e61154f565b50506001821b6105dd565b5060208310610133831016604e8410600b841016171561186c575081810a6105dd565b6118768383611790565b807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff048211156118a8576118a861154f565b029392505050565b6000610ab883836117f1565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820361191c5761191c61154f565b506001019056fea164736f6c634300080f000a")

	// Update GasPricePriceOracle Proxy Parameters
	updateFjordGasPriceOracleSource = UpgradeDepositSource{Intent: "Fjord: Gas Price Oracle Proxy Update"}
	fjordGasPriceOracleAddress      = common.HexToAddress("0xa919894851548179A0750865e7974DA599C0Fac7")

	// Enable Fjord Parameters
	enableFjordSource = UpgradeDepositSource{Intent: "Fjord: Gas Price Oracle Set Fjord"}
	enableFjordInput  = crypto.Keccak256([]byte("setFjord()"))[:4]
)

// FjordNetworkUpgradeTransactions returns the transactions required to upgrade the Fjord network.
func FjordNetworkUpgradeTransactions() ([]hexutil.Bytes, error) {
	upgradeTxns := make([]hexutil.Bytes, 0, 3)

	// Deploy Gas Price Oracle transaction
	deployGasPriceOracle, err := types.NewTx(&types.DepositTx{
		SourceHash:          deployFjordGasPriceOracleSource.SourceHash(),
		From:                GasPriceOracleFjordDeployerAddress,
		To:                  nil,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 1_450_000,
		IsSystemTransaction: false,
		Data:                gasPriceOracleFjordDeploymentBytecode,
	}).MarshalBinary()

	if err != nil {
		return nil, err
	}

	upgradeTxns = append(upgradeTxns, deployGasPriceOracle)

	updateGasPriceOracleProxy, err := types.NewTx(&types.DepositTx{
		SourceHash:          updateFjordGasPriceOracleSource.SourceHash(),
		From:                common.Address{},
		To:                  &predeploys.GasPriceOracleAddr,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 50_000,
		IsSystemTransaction: false,
		Data:                upgradeToCalldata(fjordGasPriceOracleAddress),
	}).MarshalBinary()

	if err != nil {
		return nil, err
	}

	upgradeTxns = append(upgradeTxns, updateGasPriceOracleProxy)

	enableFjord, err := types.NewTx(&types.DepositTx{
		SourceHash:          enableFjordSource.SourceHash(),
		From:                L1InfoDepositerAddress,
		To:                  &predeploys.GasPriceOracleAddr,
		Mint:                big.NewInt(0),
		Value:               big.NewInt(0),
		Gas:                 90_000,
		IsSystemTransaction: false,
		Data:                enableFjordInput,
	}).MarshalBinary()
	if err != nil {
		return nil, err
	}
	upgradeTxns = append(upgradeTxns, enableFjord)

	return upgradeTxns, nil
}
