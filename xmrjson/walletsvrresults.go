package xmrjson

type GetHeightResult struct {
	Height int64 `json:"height"`
}

type TransferResult struct {
	Amount        int64  `json:"amount"`
	Fee           int64  `json:"fee"`
	MultisigTxset string `json:"multisig_txset"`
	TxBlob        string `json:"tx_blob"`
	TxHash        string `json:"tx_hash"`
	TxKey         string `json:"tx_key"`
	TxMetadata    string `json:"tx_metadata"`
}

type GetBalanceResult struct {
	Balance              int64 `json:"balance"`
	MultisigImportNeeded bool  `json:"multisig_import_needed"`
	PerSubaddress []struct {
		Address           string `json:"address"`
		AddressIndex      int64  `json:"address_index"`
		Balance           int64  `json:"balance"`
		Label             string `json:"label"`
		NumUnspentOutputs int64  `json:"num_unspent_outputs"`
		UnlockedBalance   int64  `json:"unlocked_balance"`
	} `json:"per_subaddress"`
	UnlockedBalance int64 `json:"unlocked_balance"`
}

type GetTransfersResult struct {
	In  []transferResult `json:"in,omitempty"`
	Out []transferResult `json:"out,omitempty"`
}

type transferResult struct {
	Address         string `json:"address"`
	Amount          int64  `json:"amount"`
	DoubleSpendSeen bool   `json:"double_spend_seen"`
	Fee             int64  `json:"fee"`
	Height          int64  `json:"height"`
	Note            string `json:"note"`
	PaymentID       string `json:"payment_id"`
	SubaddrIndex struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
	} `json:"subaddr_index"`
	Timestamp  int64  `json:"timestamp"`
	Txid       string `json:"txid"`
	Type       string `json:"type"`
	UnlockTime int64  `json:"unlock_time"`
}

type GetTransferByTxidResult struct {
	Transfer struct {
		Address      string `json:"address"`
		Amount       int    `json:"amount"`
		Destinations []struct {
			Address string `json:"address"`
			Amount  int    `json:"amount"`
		} `json:"destinations"`
		DoubleSpendSeen bool   `json:"double_spend_seen"`
		Fee             int    `json:"fee"`
		Height          int    `json:"height"`
		Note            string `json:"note"`
		PaymentID       string `json:"payment_id"`
		SubaddrIndex    struct {
			Major int `json:"major"`
			Minor int `json:"minor"`
		} `json:"subaddr_index"`
		Timestamp  int    `json:"timestamp"`
		Txid       string `json:"txid"`
		Type       string `json:"type"`
		UnlockTime int    `json:"unlock_time"`
	} `json:"transfer"`
}
