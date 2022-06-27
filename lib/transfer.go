package maplerad

import (
	"fmt"
	"github.com/wirepay/maplerad-go/models"
)

type TransferService service

type BankTransferRequest struct {
	bankCode      string
	amount        string
	accountNumber string
	currency      string
}

func (c *TransferService) NigerianBankTransfer(body *BankTransferRequest) (*models.NigerianBankTransferResponse, error) {
	u := "/transfers"
	body.currency = "NGN"
	resp := &models.NigerianBankTransferResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

type USDRequest struct {
	BankCode      string
	AccountNumber string
	Amount        int
	Reason        string
	Currency      string
	Reference     string
	Meta          struct {
		Scheme string
		Sender struct {
			FirstName   string
			LastName    string
			PhoneNumber string
			Address     string
			Country     string
		}
		Counterparty struct {
			FirstName    string
			LastName     string
			Address      string
			PhoneNumber  string
			IdentityType string
			Country      string
		}
	}
}

// USDCashPickUp Set Scheme to CASHPICKUP
func (c *TransferService) USDCashPickUp(body *USDRequest) (*models.USDCashPickUpResponse, error) {
	u := "/transfers"
	resp := &models.USDCashPickUpResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

// USDDOM Set Scheme to DOM
func (c *TransferService) USDDOM(body *USDRequest) (*models.USDDOMResponse, error) {
	u := "/transfers"
	resp := &models.USDDOMResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *TransferService) GetTransfer(transferId string) (*models.GetTransferResponse, error) {
	u := fmt.Sprintf("/transfers/%s", transferId)
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
