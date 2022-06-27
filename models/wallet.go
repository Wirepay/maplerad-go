package models

type GetWalletsResponse struct {
	Generic
	Data []struct {
		Id               string `json:"id"`
		Currency         string `json:"currency"`
		LedgerBalance    int    `json:"ledger_balance"`
		AvailableBalance int    `json:"available_balance"`
		HoldingBalance   int    `json:"holding_balance"`
		Active           bool   `json:"active"`
		Disabled         bool   `json:"disabled"`
		WalletType       string `json:"wallet_type"`
	} `json:"data"`
}

type GetWalletByCurrencyResponse struct {
	Generic
	Data struct {
		Id               string `json:"id"`
		Currency         string `json:"currency"`
		LedgerBalance    int    `json:"ledger_balance"`
		AvailableBalance int    `json:"available_balance"`
		HoldingBalance   int    `json:"holding_balance"`
		Active           bool   `json:"active"`
		Disabled         bool   `json:"disabled"`
		WalletType       string `json:"wallet_type"`
	} `json:"data"`
}

type GetWalletHistoryResponse struct {
	Generic
	TransactionId        string      `json:"transaction_id"`
	RelatedTransactionId interface{} `json:"related_transaction_id"`
	WalletId             string      `json:"wallet_id"`
	Debit                int         `json:"debit"`
	Credit               int         `json:"credit"`
	PreviousBalance      int         `json:"previous_balance"`
	CurrentBalance       int         `json:"current_balance"`
	BalanceType          string      `json:"balance_type"`
	Reversal             bool        `json:"reversal"`
}

type GetWalletHistoryByCurrencyResponse struct {
	Generic
	Data []struct {
		TransactionId        string      `json:"transaction_id"`
		RelatedTransactionId interface{} `json:"related_transaction_id"`
		WalletId             string      `json:"wallet_id"`
		Debit                int         `json:"debit"`
		Credit               int         `json:"credit"`
		PreviousBalance      int         `json:"previous_balance"`
		CurrentBalance       int         `json:"current_balance"`
		BalanceType          string      `json:"balance_type"`
		Reversal             bool        `json:"reversal"`
	} `json:"data"`
}
