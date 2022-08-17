package models

import "time"

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
	Generic
	Data struct {
		ID        string    `json:"id"`
		Link      string    `json:"link"`
		Reference string    `json:"reference"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"data"`
}

type MomoResponse struct {
	Generic
	Data struct {
		ID        string    `json:"id"`
		Link      string    `json:"link"`
		Reference string    `json:"reference"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"data"`
}
