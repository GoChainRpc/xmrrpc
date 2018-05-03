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

type GetAddressCmd struct {
}

func NewGetAddressCmd() *GetAddressCmd {
	return &GetAddressCmd{}
}

// GetTransactionCmd defines the gettransaction JSON-RPC command.
type GetTransfersCmd struct {
	In     bool `json:"in,omitempty"`
	Out    bool `json:"out,omitempty"`
	Pool   bool `json:"pool,omitempty"`
	Failed bool `json:"failed,omitempty"`

	FilterByHeight bool  `json:"filter_by_height,omitempty"`
	MinHeight      int64 `json:"min_height,omitempty"`
	MaxHeight      int64 `json:"max_height,omitempty"`
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
	Amount  int64  `json:"amount"`
}

// SendToAddressCmd defines the sendtoaddress JSON-RPC command.
type TransferCmd struct {
	Destinations []TransferDestination `json:"destinations"`
	Fee          int64                 `json:"fee"`
	PaymentID    string                `json:"payment_id"`
	GetTxKey     bool                  `json:"get_tx_key"`
	Mixin        int                   `json:"mixin"`
}

// NewSendToAddressCmd returns a new instance which can be used to issue a
// sendtoaddress JSON-RPC command.
//
// The parameters which are pointers indicate they are optional.  Passing nil
// for optional parameters will use the default value.
func NewTransferCmd(transferDestination []TransferDestination, fee int64, payment_id string) *TransferCmd {
	if fee == 0 {
		return &TransferCmd{
			Destinations: transferDestination,
			// auto set fee when 0 . Fee:          fee,
			PaymentID: payment_id,
			GetTxKey:  true,
			Mixin:     0,
		}
	} else {
		return &TransferCmd{
			Destinations: transferDestination,
			Fee:          fee,
			PaymentID:    payment_id,
			GetTxKey:     true,
			Mixin:        0,
		}
	}
}

type GetTransferByTxidCmd struct {
	Txid string `json:"txid"`
}

func NewGetTransferByTxidCmd(txid string) *GetTransferByTxidCmd {
	return &GetTransferByTxidCmd{
		Txid: txid,
	}
}

type OpenWalletCmd struct {
	Filename string `json:"filename"`
	Password string `json:"password"`
}

func NewOpenWalletCmd(filename, password string) *OpenWalletCmd {
	return &OpenWalletCmd{
		Filename: filename,
		Password: password,
	}
}

type CreateWalletCmd struct {
	Filename string `json:"filename"`
	Language string `json:"language"`
	Password string `json:"password"`
}

func NewCreateWalletCmd(filename, password string) *CreateWalletCmd {
	return &CreateWalletCmd{
		Filename: filename,
		Password: password,
		Language: "English",
	}
}

func init() {
	// The commands in this file are only usable with a wallet server.
	flags := UFWalletOnly

	MustRegisterCmd("getheight", (*GetHeightCmd)(nil), flags)
	MustRegisterCmd("getbalance", (*GetBalanceCmd)(nil), flags)
	MustRegisterCmd("getaddress", (*GetAddressCmd)(nil), flags)
	MustRegisterCmd("get_transfers", (*GetTransfersCmd)(nil), flags)
	MustRegisterCmd("transfer", (*TransferCmd)(nil), flags)
	MustRegisterCmd("get_transfer_by_txid", (*GetTransferByTxidCmd)(nil), flags)
	MustRegisterCmd("open_wallet", (*OpenWalletCmd)(nil), flags)
	MustRegisterCmd("create_wallet", (*CreateWalletCmd)(nil), flags)

}
