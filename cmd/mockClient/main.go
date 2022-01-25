package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/libsv/go-bn"
	"github.com/libsv/go-bt/v2"
	"github.com/libsv/go-bt/v2/bscript"
	"github.com/nch-bowstave/paymail/data"
	"github.com/nch-bowstave/paymail/service"
)

func main() {
	c := bn.NewNodeClient(
		bn.WithHost("http://localhost:18332"),
		bn.WithCreds("bitcoin", "bitcoin"),
	)

	ht := http.Client{}
	h := data.NewClient(&ht)

	setup(c, h)
}

func setup(c bn.NodeClient, h data.HTTPClient) {
	// generate a new block.
	if _, err := c.Generate(context.Background(), 100, nil); err != nil {
		panic(err)
	}

	addr, err := c.NewAddress(context.Background(), nil)
	if err != nil {
		panic(err)
	}

	w, err := c.DumpPrivateKey(context.Background(), addr)
	if err != nil {
		panic(err)
	}

	// Spend One coinbase tx into a new utxo at a known address.
	txID, err := c.SendToAddress(context.Background(), addr, 4000, nil)
	if err != nil {
		panic(err)
	}

	// generate a new block.
	_, err = c.Generate(context.Background(), 1, nil)
	if err != nil {
		panic(err)
	}

	// grab that transaction we just made
	fundTx, err := c.RawTransaction(context.Background(), txID)
	if err != nil {
		panic(err)
	}

	ourOut := ourOutput(fundTx, addr)
	tx := bt.NewTx()
	if err := tx.FromUTXOs(&bt.UTXO{
		TxID:           fundTx.TxIDBytes(),
		LockingScript:  fundTx.Outputs[ourOut].LockingScript,
		Satoshis:       fundTx.Outputs[ourOut].Satoshis,
		SequenceNumber: 0xffffffff,
		Vout:           uint32(ourOut),
	}); err != nil {
		panic(err)
	}

	req := service.DestArgs{
		Satoshis: 2000,
	}
	res := service.DestResponse{}

	// Grab destinations via paymail endpoint
	err = h.Do(context.Background(), http.MethodPost, "http://localhost:8446/api/p2pDest/epic@nchain.com", http.StatusCreated, &req, &res)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully made p2paymail request for destinations:\n\n %+v", res)

	s, err := bscript.NewFromHexString(res.Outputs[0].Script)
	if err != nil {
		panic(err)
	}

	// PAY TO THIS
	if err := tx.PayTo(s, res.Outputs[0].Satoshis); err != nil {
		panic(err)
	}

	if err := tx.ChangeToAddress(addr, bt.NewFeeQuote()); err != nil {
		panic(err)
	}

	if err := tx.UnlockAll(context.Background(), &bt.LocalUnlockerGetter{PrivateKey: w.PrivKey}); err != nil {
		panic(err)
	}

	// send to p2pRawTx
	hexRawTx := tx.String()

	reqraw := service.TxSubmitArgs{
		RawTx:     hexRawTx,
		Reference: res.Reference,
	}
	resraw := service.TxReceipt{}

	err = h.Do(context.Background(), http.MethodPost, "http://localhost:8446/api/p2pRawTx", http.StatusCreated, &reqraw, &resraw)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully made payment to Raw p2paymail endpoint:\n\n %+v", resraw)

}

func ourOutput(tx *bt.Tx, addr string) int {
	var oIdx int
	for i := range tx.Outputs {
		pkh, err := tx.Outputs[i].LockingScript.PublicKeyHash()
		if err != nil {
			panic(err)
		}
		a, err := bscript.NewAddressFromPublicKeyHash(pkh, false)
		if err != nil {
			panic(err)
		}
		if err != nil && a.AddressString == addr {
			oIdx = i
		}
	}
	return oIdx
}
