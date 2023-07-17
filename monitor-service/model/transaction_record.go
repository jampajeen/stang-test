package model

type TransactionRecord struct {
	TxHash    string `json:"txHash"`
	TxMethod  uint8  `json:"txMethod"`
	TxBlock   uint64 `json:"txBlock"`
	TxAt      int64  `json:"txAt"`
	TxFrom    string `json:"txFrom"`
	TxTo      string `json:"txTo"`
	TxValue   int64  `json:"txValue"`
	CreatedAt string `json:"createdAt"`
}
