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
		Created_at   string `json:"created_at"`
		Updated_at   string `json:"updated_at"`
		Counterparty struct {
			Id             string `json:"id"`
			Account_number string `json:"account_number"`
			Account_name   string `json:"account_name"`
			Bank_code      string `json:"bank_code"`
			Bank_name      string `json:"bank_name"`
		} `json:"counterparty"`
	}
}

type USDCashPickUpResponse struct {
}

type USDDOMResponse struct{}

type GetTransferResponse struct {
	Generic
	Data struct {
		Id             string `json:"id"`
		Account_number string
		Bank_code      string
		Currency       string `json:"currency"`
		Status         string `json:"status"`
		Entry          string `json:"entry"`
		Type           string `json:"type"`
		Amount         int    `json:"amount"`
		Summary        string `json:"summary"`
		Reason         string `json:"reason"`
		Fee            int    `json:"fee"`
	}
}
