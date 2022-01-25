package bn

import (
	"context"

	"github.com/libsv/go-bk/wif"
	imodels "github.com/libsv/go-bn/internal/models"
	"github.com/libsv/go-bn/internal/util"
	"github.com/libsv/go-bn/models"
	"github.com/libsv/go-bt/v2"
)

// WalletClient interfaces interaction with the wallet sub commands on a bitcoin node.
type WalletClient interface {
	AbandonTransaction(ctx context.Context, txID string) error
	AddMultiSigAddress(ctx context.Context, n int, keys ...string) (string, error)
	BackupWallet(ctx context.Context, dest string) error
	DumpPrivateKey(ctx context.Context, address string) (*wif.WIF, error)
	DumpWallet(ctx context.Context, dest string) (*models.DumpWallet, error)
	Account(ctx context.Context, address string) (string, error)
	AccountAddress(ctx context.Context, account string) (string, error)
	AccountAddresses(ctx context.Context, account string) ([]string, error)
	Balance(ctx context.Context, opts *models.OptsBalance) (uint64, error)
	UnconfirmedBalance(ctx context.Context) (uint64, error)
	NewAddress(ctx context.Context, opts *models.OptsNewAddress) (string, error)
	RawChangeAddress(ctx context.Context) (string, error)
	ReceivedByAddress(ctx context.Context, address string) (uint64, error)
	Transaction(ctx context.Context, txID string) (*models.Transaction, error)
	ImportAddress(ctx context.Context, address string, opts *models.OptsImportAddress) error
	WalletInfo(ctx context.Context) (*models.WalletInfo, error)
	ImportMulti(ctx context.Context, reqs []models.ImportMultiRequest,
		opts *models.OptsImportMulti) ([]*models.ImportMulti, error)
	ImportPrivateKey(ctx context.Context, w *wif.WIF, opts *models.OptsImportPrivateKey) error
	ImportPrunedFunds(ctx context.Context, tx *bt.Tx, txOutProof string) error
	ImportPublicKey(ctx context.Context, publicKey string, opts *models.OptsImportPublicKey) error
	ImportWallet(ctx context.Context, filename string) error
	KeypoolRefill(ctx context.Context, opts *models.OptsKeypoolRefill) error
	ListAccounts(ctx context.Context, opts *models.OptsListAccounts) (map[string]uint64, error)
	ListLockUnspent(ctx context.Context) ([]*models.LockUnspent, error)
	ListReceivedByAccount(ctx context.Context, opts *models.OptsListReceivedBy) ([]*models.ReceivedByAccount, error)
	ListReceivedByAddress(ctx context.Context, opts *models.OptsListReceivedBy) ([]*models.ReceivedByAddress, error)
	ListSinceBlock(ctx context.Context, opts *models.OptsListSinceBlock) (*models.SinceBlock, error)
	ListTransactions(ctx context.Context, opts *models.OptsListTransactions) ([]*models.Transaction, error)
	ListUnspent(ctx context.Context, opts *models.OptsListUnspent) (bt.UTXOs, error)
	ListWallets(ctx context.Context) ([]string, error)
	LockUnspent(ctx context.Context, lock bool, opts *models.OptsLockUnspent) (bool, error)
	Move(ctx context.Context, from, to string, amount uint64, opts *models.OptsMove) (bool, error)
	RemovePrunedFunds(ctx context.Context, txID string) error
	SendFrom(ctx context.Context, from, to string, amount uint64, opts *models.OptsSendFrom) (string, error)
	SendMany(ctx context.Context, from string, amounts map[string]uint64, opts *models.OptsSendMany) (string, error)
	SendToAddress(ctx context.Context, address string, amount uint64, opts *models.OptsSendToAddress) (string, error)
	SetAccount(ctx context.Context, address, account string) error
	SetTxFee(ctx context.Context, amount uint64) (bool, error)
	SignMessage(ctx context.Context, address, message string) (string, error)
	EncryptWallet(ctx context.Context, passphrase string) error
	WalletPhassphrase(ctx context.Context, passphrase string, timeout int) error
	WalletPhassphraseChange(ctx context.Context, oldPassphrase, newPassphrase string) error
	WalletLock(ctx context.Context) error
}

// NewWalletClient returns a client only capable of interfacing with the wallet sub commands on a bitcoin node.
func NewWalletClient(oo ...BitcoinClientOptFunc) WalletClient {
	return NewNodeClient(oo...)
}

func (c *client) AbandonTransaction(ctx context.Context, txID string) error {
	return c.rpc.Do(ctx, "abandontransaction", nil, txID)
}

func (c *client) AddMultiSigAddress(ctx context.Context, n int, keys ...string) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "addmultisigaddress", &resp, n, keys)
}

func (c *client) BackupWallet(ctx context.Context, dest string) error {
	return c.rpc.Do(ctx, "backupwallet", nil, dest)
}

func (c *client) DumpPrivateKey(ctx context.Context, address string) (*wif.WIF, error) {
	var resp imodels.InternalDumpPrivateKey
	return resp.WIF, c.rpc.Do(ctx, "dumpprivkey", &resp, address)
}

// TODO: do not cache
func (c *client) DumpWallet(ctx context.Context, dest string) (*models.DumpWallet, error) {
	var resp models.DumpWallet
	return &resp, c.rpc.Do(ctx, "dumpwallet", &resp, dest)
}

func (c *client) Account(ctx context.Context, address string) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "getaccount", &resp, address)
}

func (c *client) AccountAddress(ctx context.Context, account string) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "getaccountaddress", &resp, account)
}

func (c *client) AccountAddresses(ctx context.Context, account string) ([]string, error) {
	var resp []string
	return resp, c.rpc.Do(ctx, "getaddressesbyaccount", &resp, account)
}

// TODO: do not cache
func (c *client) Balance(ctx context.Context, opts *models.OptsBalance) (uint64, error) {
	var resp float64
	err := c.rpc.Do(ctx, "getbalance", &resp, c.argsFor(opts)...)
	return util.BSVToSatoshis(resp), err
}

// TODO: do not cache
func (c *client) UnconfirmedBalance(ctx context.Context) (uint64, error) {
	var resp float64
	err := c.rpc.Do(ctx, "getunconfirmedbalance", &resp)
	return util.BSVToSatoshis(resp), err
}

// TODO: do not cache
func (c *client) NewAddress(ctx context.Context, opts *models.OptsNewAddress) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "getnewaddress", &resp, c.argsFor(opts)...)
}

// TODO: do not cache
func (c *client) RawChangeAddress(ctx context.Context) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "getrawchangeaddress", &resp)
}

// TODO: do not cache
func (c *client) ReceivedByAddress(ctx context.Context, address string) (uint64, error) {
	var resp float64
	err := c.rpc.Do(ctx, "getreceivedbyaddress", &resp, address)
	return util.BSVToSatoshis(resp), err
}

func (c *client) Transaction(ctx context.Context, txID string) (*models.Transaction, error) {
	var resp imodels.InternalTransaction
	return resp.Transaction, c.rpc.Do(ctx, "gettransaction", &resp, txID)
}

func (c *client) ImportAddress(ctx context.Context, address string, opts *models.OptsImportAddress) error {
	return c.rpc.Do(ctx, "importaddress", nil, c.argsFor(opts)...)
}

func (c *client) WalletInfo(ctx context.Context) (*models.WalletInfo, error) {
	var resp models.WalletInfo
	return &resp, c.rpc.Do(ctx, "getwalletinfo", &resp)
}

func (c *client) ImportMulti(ctx context.Context, reqs []models.ImportMultiRequest,
	opts *models.OptsImportMulti) ([]*models.ImportMulti, error) {
	var resp []*models.ImportMulti
	return resp, c.rpc.Do(ctx, "importmulti", &resp, c.argsFor(opts, reqs)...)
}

func (c *client) ImportPrunedFunds(ctx context.Context, tx *bt.Tx, txOutProof string) error {
	return c.rpc.Do(ctx, "importprunedfunds", nil, tx.String(), txOutProof)
}

func (c *client) ImportPublicKey(ctx context.Context, publicKey string, opts *models.OptsImportPublicKey) error {
	return c.rpc.Do(ctx, "importpubkey", nil, c.argsFor(opts, publicKey)...)
}

// TODO: don't cache. test.
func (c *client) ImportPrivateKey(ctx context.Context, w *wif.WIF, opts *models.OptsImportPrivateKey) error {
	return c.rpc.Do(ctx, "importprivkey", nil, c.argsFor(opts, w.String())...)
}

func (c *client) ImportWallet(ctx context.Context, filename string) error {
	return c.rpc.Do(ctx, "importwallet", nil, filename)
}

func (c *client) KeypoolRefill(ctx context.Context, opts *models.OptsKeypoolRefill) error {
	return c.rpc.Do(ctx, "keypoolrefill", nil, c.argsFor(opts)...)
}

func (c *client) ListAccounts(ctx context.Context, opts *models.OptsListAccounts) (map[string]uint64, error) {
	var resp map[string]float64
	err := c.rpc.Do(ctx, "listaccounts", &resp, c.argsFor(opts)...)
	return util.MapBSVToSatoshis(resp), err
}

func (c *client) ListLockUnspent(ctx context.Context) ([]*models.LockUnspent, error) {
	var resp []*models.LockUnspent
	return resp, c.rpc.Do(ctx, "listlockunspent", &resp)
}

func (c *client) ListReceivedByAccount(ctx context.Context,
	opts *models.OptsListReceivedBy) ([]*models.ReceivedByAccount, error) {
	var resp []*models.ReceivedByAccount
	return resp, c.rpc.Do(ctx, "listreceivedbyaccount", &resp, c.argsFor(opts)...)
}

func (c *client) ListReceivedByAddress(ctx context.Context,
	opts *models.OptsListReceivedBy) ([]*models.ReceivedByAddress, error) {
	var resp []*models.ReceivedByAddress
	return resp, c.rpc.Do(ctx, "listreceivedbyaddress", &resp, c.argsFor(opts)...)
}

func (c *client) ListSinceBlock(ctx context.Context, opts *models.OptsListSinceBlock) (*models.SinceBlock, error) {
	var resp models.SinceBlock
	return &resp, c.rpc.Do(ctx, "listsinceblock", &resp, c.argsFor(opts)...)
}

func (c *client) ListTransactions(ctx context.Context,
	opts *models.OptsListTransactions) ([]*models.Transaction, error) {
	var resp []*models.Transaction
	return resp, c.rpc.Do(ctx, "listtransactions", &resp, c.argsFor(opts)...)
}

func (c *client) ListUnspent(ctx context.Context, opts *models.OptsListUnspent) (bt.UTXOs, error) {
	var resp bt.UTXOs
	return resp, c.rpc.Do(ctx, "listunspent", &resp, c.argsFor(opts)...)
}

func (c *client) ListWallets(ctx context.Context) ([]string, error) {
	var resp []string
	return resp, c.rpc.Do(ctx, "listwallets", &resp)
}

// TODO: do not cache
func (c *client) LockUnspent(ctx context.Context, lock bool, opts *models.OptsLockUnspent) (bool, error) {
	var resp bool
	return resp, c.rpc.Do(ctx, "lockunspent", &resp, c.argsFor(opts, lock)...)
}

func (c *client) Move(ctx context.Context, from, to string, amount uint64, opts *models.OptsMove) (bool, error) {
	var resp bool
	return resp, c.rpc.Do(ctx, "move", &resp, c.argsFor(opts, from, to, util.SatoshisToBSV(int64(amount)))...)
}

func (c *client) RemovePrunedFunds(ctx context.Context, txID string) error {
	return c.rpc.Do(ctx, "removeprunedfunds", nil, txID)
}

func (c *client) SendFrom(ctx context.Context, from, to string, amount uint64,
	opts *models.OptsSendFrom) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "sendfrom", &resp, c.argsFor(opts, from, to, util.SatoshisToBSV(int64(amount)))...)
}

func (c *client) SendMany(ctx context.Context, from string, amounts map[string]uint64,
	opts *models.OptsSendMany) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "sendmany", &resp, c.argsFor(opts, from, util.MapSatoshisToBSV(amounts))...)
}

func (c *client) SendToAddress(ctx context.Context, address string, amount uint64,
	opts *models.OptsSendToAddress) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "sendtoaddress", &resp, c.argsFor(opts, address, util.SatoshisToBSV(int64(amount)))...)
}

func (c *client) SetAccount(ctx context.Context, address, account string) error {
	return c.rpc.Do(ctx, "setaccount", nil, address, account)
}

func (c *client) SetTxFee(ctx context.Context, amount uint64) (bool, error) {
	var resp bool
	return resp, c.rpc.Do(ctx, "settxfee", &resp, util.SatoshisToBSV(int64(amount)))
}

func (c *client) SignMessage(ctx context.Context, address, message string) (string, error) {
	var resp string
	return resp, c.rpc.Do(ctx, "signmessage", &resp, address, message)
}

func (c *client) EncryptWallet(ctx context.Context, passphrase string) error {
	return c.rpc.Do(ctx, "encryptwallet", nil, passphrase)
}

func (c *client) WalletPhassphrase(ctx context.Context, passphrase string, timeout int) error {
	return c.rpc.Do(ctx, "walletpassphrase", nil, passphrase, timeout)
}

func (c *client) WalletPhassphraseChange(ctx context.Context, oldPassphrase, newPassphrase string) error {
	return c.rpc.Do(ctx, "walletpassphrasechange", nil, oldPassphrase, newPassphrase)
}

func (c *client) WalletLock(ctx context.Context) error {
	return c.rpc.Do(ctx, "walletlock", nil)
}
