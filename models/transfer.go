package models

type NigerianBankTransferResponse struct {
	Generic
	Data struct {
		Id           string `json:"id"`
		Currency     string `json:"currency"`
		Status       string `json:"status"`
		Entry        string `json:"entry"`
		Type         string `json:"type"`
		Amount       int    `json:"amount"`
		Summary      string `json:"summary"`
		Reason       string `json:"reason"`
		Fee          int    `json:"fee"`
		Reference    string `json:"reference"`
		CreatedAt    string `json:"created_at"`
		UpdatedAt    string `json:"updated_at"`
		Counterparty struct {
			Id            string `json:"id"`
			AccountNumber string `json:"account_number"`
			AccountName   string `json:"account_name"`
			BankCode      string `json:"bank_code"`
			BankName      string `json:"bank_name"`
		} `json:"counterparty"`
	}
}

type USDCashPickUpResponse struct {
}

type USDDOMResponse struct{}

type GetTransferResponse struct {
	Generic
	Data struct {
		Id            string `json:"id"`
		AccountNumber string
		BankCode      string
		Currency      string `json:"currency"`
		Status        string `json:"status"`
		Entry         string `json:"entry"`
		Type          string `json:"type"`
		Amount        int    `json:"amount"`
		Summary       string `json:"summary"`
		Reason        string `json:"reason"`
		Fee           int    `json:"fee"`
	}
}
