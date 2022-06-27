package models

type CreateAccountResponse struct {
	Generic
	Data struct {
		Id            string `json:"id"`
		BankName      string `json:"bank_name"`
		AccountNumber string `json:"account_number"`
		AccountName   string `json:"account_name"`
		Currency      string `json:"currency"`
		CreatedAt     string `json:"created_at"`
	} `json:"data"`
}

type DirectDebitResponse struct {
}
