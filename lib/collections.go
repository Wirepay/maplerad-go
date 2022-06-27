package maplerad

import (
	"github.com/wirepay/maplerad-go/models"
)

type CollectionsService service

type CreateAccountRequest struct {
	customerId string
	currency   string
}

type DirectDebitRequest struct {
	currency    string
	reference   string
	amount      int
	description string
}

func (c *CollectionsService) CreateAccount(customerId, currency string) (*models.CreateAccountResponse, error) {
	u := "/collections/virtual-account"
	var body *CreateAccountRequest
	body.customerId = customerId
	body.currency = currency
	resp := &models.CreateAccountResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *CollectionsService) GenerateDirectDebit(body *DirectDebitRequest) (*models.DirectDebitResponse, error) {
	u := "/collections/direct-debit"
	resp := &models.DirectDebitResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}
