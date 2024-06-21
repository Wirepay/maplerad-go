package maplerad

import (
	"fmt"
	"github.com/wirepay/maplerad-go/models"
	"github.com/wirepay/maplerad-go/utils"
)

type TransferService service

type BankTransferRequest struct {
	BankCode      string
	Amount        string
	AccountNumber string
	Currency      utils.CurrencySymbol
}

func (c *TransferService) NigerianBankTransfer(body *BankTransferRequest) (*models.NigerianBankTransferResponse, error) {
	u := "/transfers"
	body.Currency = "NGN"
	resp := &models.NigerianBankTransferResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

type USDRequest struct {
	BankCode      string
	AccountNumber string
	Amount        int
	Reason        string
	Currency      utils.CurrencySymbol
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
