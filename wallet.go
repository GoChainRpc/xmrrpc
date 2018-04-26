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
func (c *Client) GetTransactionAsync(in, out, pool, failed, filter_by_hieght bool,
	min_height, max_height int64) FutureGetTransfersResult {
	cmd := xmrjson.NewGetTransfersCmd(in, out, pool, failed, filter_by_hieght, min_height, max_height)
	return c.sendCmd(cmd)
}

// GetTransaction returns detailed information about a wallet transaction.
//
// See GetRawTransaction to return the raw transaction instead.
func (c *Client) GetTransaction(in, out, pool, failed, filter_by_hieght bool,
	min_height, max_height int64) (*xmrjson.GetTransfersResult, error) {

	return c.GetTransactionAsync(in, out, pool, failed, filter_by_hieght, min_height, max_height).Receive()
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

func (c *Client) TransferAsync(transferDestination []xmrjson.TransferDestination, payment_id string) FutureTransferResult {
	cmd := xmrjson.NewTransferCmd(transferDestination, payment_id)
	return c.sendCmd(cmd)
}

func (c *Client) Transfer(transferDestination []xmrjson.TransferDestination, payment_id string) (*xmrjson.TransferResult, error) {
	return c.TransferAsync(transferDestination, payment_id).Receive()
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
