package bindings

import (
	"fmt"

	"github.com/bobanetwork/v3-anchorage/boba-bindings/solc"
	"github.com/ledgerwatch/erigon/common"
)

var layouts = make(map[string]*solc.StorageLayout)

var deployedBytecodes = make(map[string]string)

func GetStorageLayout(name string) (*solc.StorageLayout, error) {
	layout := layouts[name]
	if layout == nil {
		return nil, fmt.Errorf("%s: storage layout not found", name)
	}
	return layout, nil
}

func GetDeployedBytecode(name string) ([]byte, error) {
	bc := deployedBytecodes[name]
	if bc == "" {
		return nil, fmt.Errorf("%s: deployed bytecode not found", name)
	}

	return common.FromHex(bc), nil
}
