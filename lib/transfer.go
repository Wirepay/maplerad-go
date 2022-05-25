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

func (c *TransferService) NigerianBankTransfer(bank_code, account_number, amount string) (*models.NigerianBankTransferResponse, error) {
	u := "/transfers"
	var body *BankTransferRequest
	body.amount = amount
	body.currency = "NGN"
	body.account_number = account_number
	body.bank_code = bank_code
	resp := &models.NigerianBankTransferResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

type USDRequest struct {
	Bank_code      string
	Account_number string
	Amount         int
	Reason         string
	Currency       string
	Reference      string
	Meta           struct {
		Scheme string
		Sender struct {
			First_name   string
			Last_name    string
			Phone_number string
			Address      string
			Country      string
		}
		Counterparty struct {
			First_name    string
			Last_name     string
			Address       string
			Phone_number  string
			Identity_type string
			Country       string
		}
	}
}

// Set Scheme to CASHPICKUP
func (c *TransferService) USDCashPickUp(body *USDRequest) (*models.USDCashPickUpResponse, error) {
	u := "/transfers"
	resp := &models.USDCashPickUpResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

// Set Scheme to DOM
func (c *TransferService) USDDOM(body *USDRequest) (*models.USDDOMResponse, error) {
	u := "/transfers"
	resp := &models.USDDOMResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *TransferService) GetTransfer(transfer_id string) (*models.GetTransferResponse, error) {
	u := fmt.Sprintf("/transfers/%s", transfer_id)
	resp := &models.GetTransferResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}

// func (c *TransferService) GetTransferTransactions() (*models.GetTransferResponse, error) {
// 	u := "/transactions"
// 	resp := &models.GetTransferResponse{}
// 	err := c.client.Call("GET", u, nil, nil, &resp)
// 	return resp, err
// }
