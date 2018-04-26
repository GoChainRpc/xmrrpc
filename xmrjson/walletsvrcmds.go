package xmrjson

// GetHeightCmd defines the getheight JSON-RPC command.
type GetHeightCmd struct {
}

// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewGetHeightCmd() *GetHeightCmd {
	return &GetHeightCmd{}
}

// GetBalanceCmd defines the getheight JSON-RPC command.
type GetBalanceCmd struct {
}

// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewGetBalanceCmd() *GetBalanceCmd {
	return &GetBalanceCmd{}
}

// GetTransactionCmd defines the gettransaction JSON-RPC command.
type GetTransfersCmd struct {
	In     bool `json:"in"`
	Out    bool `json:"out"`
	Pool   bool `json:"pool"`
	Failed bool `json:"failed"`

	FilterByHeight bool  `json:"filter_by_height"`
	MinHeight      int64 `json:"min_height"`
	MaxHeight      int64 `json:"max_height"`
}

// NewGetTransfersCmd returns a new instance which can be used to issue a
// get_transfers JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewGetTransfersCmd(in, out, pool, failed, filter_by_hieght bool,
	min_height, max_height int64) *GetTransfersCmd {
	return &GetTransfersCmd{
		In:     in,
		Out:    out,
		Pool:   pool,
		Failed: failed,

		FilterByHeight: filter_by_hieght,
		MinHeight:      min_height,
		MaxHeight:      max_height,
	}
}

type TransferDestination struct {
	Address string `json:"address"`
	Amount  int    `json:"amount"`
}

// SendToAddressCmd defines the sendtoaddress JSON-RPC command.
type TransferCmd struct {
	Destinations []TransferDestination `json:"destinations"`
	PaymentID    string                `json:"payment_id"`
	GetTxKey     bool                  `json:"get_tx_key"`
	Mixin        int                   `json:"mixin"`
}

// NewSendToAddressCmd returns a new instance which can be used to issue a
// sendtoaddress JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewTransferCmd(transferDestination []TransferDestination, payment_id string) *TransferCmd {
	return &TransferCmd{
		Destinations: transferDestination,
		PaymentID:    payment_id,
		GetTxKey:     true,
		Mixin:        2,
	}
}

func init() {
	// The commands in this file are only usable with a wallet server.
	//flags := UFWalletOnly
	flags := UsageFlag(0)


	MustRegisterCmd("getheight", (*GetHeightCmd)(nil), flags)
	MustRegisterCmd("getbalance", (*GetBalanceCmd)(nil), flags)
	MustRegisterCmd("get_transfers", (*GetTransfersCmd)(nil), flags)
	MustRegisterCmd("transfer", (*TransferCmd)(nil), flags)

}