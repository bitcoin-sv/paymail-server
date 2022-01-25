package models

import (
	"github.com/libsv/go-bn/internal/util"
	"github.com/libsv/go-bn/models"
	"github.com/libsv/go-bt/v2"
)

// InternalFundRawTransaction the true to form fundrawtransaction response.
type InternalFundRawTransaction struct {
	*models.FundRawTransaction
	Hex    string  `json:"hex"`
	BsvFee float64 `json:"fee"`
}

// PostProcess an RPC response.
func (i *InternalFundRawTransaction) PostProcess() error {
	var err error
	i.Tx, err = bt.NewTxFromString(i.Hex)
	i.Fee = util.BSVToSatoshis(i.BsvFee)
	return err
}

// InternalSignRawTransaction the true to form signrawtransaction response.
type InternalSignRawTransaction struct {
	Hex string `json:"hex"`
	*models.SignedRawTransaction
}

// PostProcess an RPC response.
func (i *InternalSignRawTransaction) PostProcess() error {
	var err error
	i.Tx, err = bt.NewTxFromString(i.Hex)
	return err
}
