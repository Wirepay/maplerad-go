package maplerad

import (
	"github.com/wirepay/maplerad-go/models"
)

type CollectionsService service

type CreateAccountRequest struct {
	CustomerId string `json:"customer_id"`
	Currency   string `json:"currency"`
}

type DirectDebitRequest struct {
	Currency    string `json:"currency"`
	Reference   string `json:"reference"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

type MomoRequest struct {
	Currency    string `json:"currency"`
	Reference   string `json:"reference"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
}

func (c *CollectionsService) CreateAccount(customerId, currency string) (*models.CreateAccountResponse, error) {
	u := "/collections/virtual-account"
	var body *CreateAccountRequest
	body.CustomerId = customerId
	body.Currency = currency
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

func (c *CollectionsService) MobileMoney(body *MomoRequest) (*models.MomoResponse, error) {
	u := "/collections/momo"
	resp := &models.MomoResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}
