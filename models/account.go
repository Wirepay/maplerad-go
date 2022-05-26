package models

type CreateAccountResponse struct {
	Generic
	Data struct {
		Id             string `json:"id"`
		Bank_name      string `json:"bank_name"`
		Account_number string `json:"account_number"`
		Account_name   string `json:"account_name"`
		Currency       string `json:"currency"`
		Created_at     string `json:"created_at"`
	} `json:"data"`
}
