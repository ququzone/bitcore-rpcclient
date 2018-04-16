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

// NewGetAddressUTXOsCmd ...
func NewGetAddressUTXOsCmd(address *AddressesRequest) *GetAddressUTXOsCmd {
	return &GetAddressUTXOsCmd{
		Address: address,
	}
}

func init() {
	flags := btcjson.UFWalletOnly
	btcjson.MustRegisterCmd("getaddressbalance", (*GetAddressBalanceCmd)(nil), flags)
	btcjson.MustRegisterCmd("getaddressutxos", (*GetAddressUTXOsCmd)(nil), flags)
}
