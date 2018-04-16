package rpcclient

import (
	"encoding/json"
)

// jsonRequest holds information about a json request that is used to properly
// detect, interpret, and deliver a reply to it.
type jsonRequest struct {
	id             uint64
	method         string
	cmd            interface{}
	marshalledJSON []byte
}

// Response ...
type Response struct {
	ID     int32           `json:"id"`
	Result json.RawMessage `json:"result"`
	Error  struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

// AddressesRequest ...
type AddressesRequest struct {
	Addresses []string `json:"addresses"`
}

// Balance ...
type Balance struct {
	Balance  int64 `json:"balance"`
	Received int64 `json:"received"`
}

// UTXO ...
type UTXO struct {
	Address     string `json:"address"`
	TxID        string `json:"txid"`
	OutputIndex uint32 `json:"outputIndex"`
	Scirpt      string `json:"script"`
	Satoshis    int64  `json:"satoshis"`
	Height      int64  `json:"height"`
}
