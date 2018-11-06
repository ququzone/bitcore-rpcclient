package rpcclient

import (
	"github.com/btcsuite/btcd/btcjson"
)

// GetAddressBalanceCmd getaddressbalance
type GetAddressBalanceCmd struct {
	Address *AddressesRequest
}

// NewGetAddressBalanceCmd ...
func NewGetAddressBalanceCmd(address *AddressesRequest) *GetAddressBalanceCmd {
	return &GetAddressBalanceCmd{
		Address: address,
	}
}

// GetAddressUTXOsCmd getaddressutxos
type GetAddressUTXOsCmd struct {
	Address *AddressesRequest
}

// GetBlockVerbosityCmd getblock verbosity == 2
type GetBlockVerbosityCmd struct {
	Block *BlockVerbosityRequest
}

// NewGetAddressUTXOsCmd ...
func NewGetAddressUTXOsCmd(address *AddressesRequest) *GetAddressUTXOsCmd {
	return &GetAddressUTXOsCmd{
		Address: address,
	}
}

// NewGetBlockVerbosityCmd ...
func NewGetBlockVerbosityCmd(hash string, verbosity int) *GetBlockVerbosityCmd {
	return &GetBlockVerbosityCmd{
		&BlockVerbosityRequest{
			Hash:      hash,
			Verbosity: verbosity,
		},
	}
}

func init() {
	flags := btcjson.UFWalletOnly
	btcjson.MustRegisterCmd("getaddressbalance", (*GetAddressBalanceCmd)(nil), flags)
	btcjson.MustRegisterCmd("getaddressutxos", (*GetAddressUTXOsCmd)(nil), flags)
	btcjson.MustRegisterCmd("getblock", (*GetBlockVerbosityCmd)(nil), flags)
}
