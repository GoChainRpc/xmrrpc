package xmrrpc

import (
	"encoding/json"
	"github.com/GoChainRpc/xmrrpc/xmrjson"
)

// FutureGetHeightResult is a future promise to deliver the result of a
// GetHeightAsync RPC invocation (or an applicable error).
type FutureGetHeightResult chan *response

// Receive waits for the response promised by the future and returns a new
// address.
func (r FutureGetHeightResult) Receive() (int64, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return 0, err
	}

	// Unmarshal result as a string.
	var get_height_result xmrjson.GetHeightResult
	err = json.Unmarshal(res, &get_height_result)
	if err != nil {
		return 0, err
	}

	return get_height_result.Height, nil
}

// GetHeightAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetHeight for the blocking version and more details.
func (c *Client) GetHeightAsync() FutureGetHeightResult {
	cmd := xmrjson.NewGetHeightCmd()
	return c.sendCmd(cmd)
}

// GetHeight returns wallet height.
func (c *Client) GetHeight() (int64, error) {
	return c.GetHeightAsync().Receive()
}

// FutureGetTransactionResult is a future promise to deliver the result
// of a GetTransactionAsync RPC invocation (or an applicable error).
type FutureGetTransfersResult chan *response

// Receive waits for the response promised by the future and returns detailed
// information about a wallet transaction.
func (r FutureGetTransfersResult) Receive() (*xmrjson.GetTransfersResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}
	// Unmarshal result as a gettransaction result object
	var getTx xmrjson.GetTransfersResult
	err = json.Unmarshal(res, &getTx)
	if err != nil {
		return nil, err
	}

	return &getTx, nil
}

// GetTransactionAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on
// the returned instance.
//
// See GetTransaction for the blocking version and more details.
func (c *Client) GetTransfersAsync(in, out, pool, failed, filter_by_hieght bool,
	min_height, max_height int64) FutureGetTransfersResult {
	cmd := xmrjson.NewGetTransfersCmd(in, out, pool, failed, filter_by_hieght, min_height, max_height)
	return c.sendCmd(cmd)
}

// GetTransaction returns detailed information about a wallet transaction.
//
// See GetRawTransaction to return the raw transaction instead.
func (c *Client) GetTransfers(in, out, pool, failed, filter_by_hieght bool,
	min_height, max_height int64) (*xmrjson.GetTransfersResult, error) {

	return c.GetTransfersAsync(in, out, pool, failed, filter_by_hieght, min_height, max_height).Receive()
}

type FutureTransferResult chan *response

func (r FutureTransferResult) Receive() (*xmrjson.TransferResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string.
	var transferResult xmrjson.TransferResult
	err = json.Unmarshal(res, &transferResult)
	if err != nil {
		return nil, err
	}

	return &transferResult, nil
}

func (c *Client) TransferAsync(transferDestination []xmrjson.TransferDestination, fee int64, payment_id string) FutureTransferResult {
	cmd := xmrjson.NewTransferCmd(transferDestination, fee, payment_id)
	return c.sendCmd(cmd)
}

func (c *Client) Transfer(transferDestination []xmrjson.TransferDestination, fee int64, payment_id string) (*xmrjson.TransferResult, error) {
	return c.TransferAsync(transferDestination, fee, payment_id).Receive()
}

// FutureGetBalanceResult is a future promise to deliver the result of a
// GetBalanceAsync or GetBalanceMinConfAsync RPC invocation (or an applicable
// error).
type FutureGetBalanceResult chan *response

// Receive waits for the response promised by the future and returns the
// available balance from the server for the specified account.
func (r FutureGetBalanceResult) Receive() (*xmrjson.GetBalanceResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}

	// Unmarshal result as a string
	var balanceResult xmrjson.GetBalanceResult
	err = json.Unmarshal(res, &balanceResult)
	if err != nil {
		return nil, err
	}

	return &balanceResult, nil
}

// GetBalanceAsync returns an instance of a type that can be used to get the
// result of the RPC at some future time by invoking the Receive function on the
// returned instance.
//
// See GetBalance for the blocking version and more details.
func (c *Client) GetBalanceAsync() FutureGetBalanceResult {
	cmd := xmrjson.NewGetBalanceCmd()
	return c.sendCmd(cmd)
}

// GetBalance returns the available balance from the server for the specified
// account using the default number of minimum confirmations.  The account may
// be "*" for all accounts.
//
// See GetBalanceMinConf to override the minimum number of confirmations.
func (c *Client) GetBalance() (*xmrjson.GetBalanceResult, error) {
	return c.GetBalanceAsync().Receive()
}

type FutureGetAddressResult chan *response

func (r FutureGetAddressResult) Receive() (*xmrjson.GetAddressResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}
	var getAddressResult xmrjson.GetAddressResult
	err = json.Unmarshal(res, &getAddressResult)
	if err != nil {
		return nil, err
	}
	return &getAddressResult, nil
}

func (c *Client) GetAddressAsync() FutureGetAddressResult {
	cmd := xmrjson.NewGetAddressCmd()
	return c.sendCmd(cmd)
}

func (c *Client) GetAddress() (*xmrjson.GetAddressResult,error) {
	return c.GetAddressAsync().Receive()
}

// async get transfer by txid
type FutureGetTransferByTxidResult chan *response

func (r FutureGetTransferByTxidResult) Receive() (*xmrjson.GetTransferByTxidResult, error) {
	res, err := receiveFuture(r)
	if err != nil {
		return nil, err
	}
	var getTransferByTxid xmrjson.GetTransferByTxidResult
	err = json.Unmarshal(res, &getTransferByTxid)
	if err != nil {
		return nil, err
	}
	return &getTransferByTxid, nil
}
func (c *Client) GetTransferByTxidAsync(txid string) FutureGetTransferByTxidResult {
	cmd := xmrjson.NewGetTransferByTxidCmd(txid)
	return c.sendCmd(cmd)
}
func (c *Client) GetTransferByTxid(txid string) (*xmrjson.GetTransferByTxidResult, error) {
	return c.GetTransferByTxidAsync(txid).Receive()
}

// async open wallet
type FutureOpenWalletResult chan *response

func (r FutureOpenWalletResult) Receive() error {
	_, err := receiveFuture(r)
	if err != nil {
		return err
	}
	return nil
}
func (c *Client) OpenWalletAsync(filename, password string) FutureOpenWalletResult {
	cmd := xmrjson.NewOpenWalletCmd(filename, password)
	return c.sendCmd(cmd)
}
func (c *Client) OpenWallet(filename, password string) (error) {
	return c.OpenWalletAsync(filename, password).Receive()
}

// async create wallet
type FutureCreateWalletResult chan *response

func (r FutureCreateWalletResult) Receive() error {
	_, err := receiveFuture(r)
	if err != nil {
		return err
	}
	return nil
}
func (c *Client) CreateWalletAsync(filename, password string) FutureCreateWalletResult {
	cmd := xmrjson.NewCreateWalletCmd(filename, password)
	return c.sendCmd(cmd)
}
func (c *Client) CreateWallet(filename, password string) (error) {
	return c.CreateWalletAsync(filename, password).Receive()
}
