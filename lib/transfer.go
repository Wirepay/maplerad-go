package maplerad

import (
	"fmt"
	"maplerad/models"
)

/* This resource will return income information to the account. */
// func (c *ConnectService) GetIncome(userID string) (interface{}, interface{}) {
// 	u := fmt.Sprintf("/accounts/%s/income", userID)
// 	resp := &models.Income{}
// 	err := c.client.Call("GET", u, "", nil, &resp)
// 	return resp, err
// }

type TransferService service

type BankTransferRequest struct {
	bank_code      string
	amount         string
	account_number string
	currency       string
}

func (c *TransferService) NigerianBankTransfer(bank_code, account_number, amount string) (*models.Generic, error) {
	u := fmt.Sprintf("/transfers")
	var body *BankTransferRequest
	body.amount = amount
	body.currency = "NGN"
	body.account_number = account_number
	body.bank_code = bank_code
	resp := &models.NigerianBankTransferResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}
