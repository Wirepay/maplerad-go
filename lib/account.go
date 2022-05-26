package maplerad

import (
	"maplerad/models"
)

type AccountService service

type CreateAccountRequest struct {
	customer_id string
	currency    string
}

func (c *AccountService) CreateAccount(customer_id, currency string) (*models.CreateAccountResponse, error) {
	u := "/accounts"
	var body *CreateAccountRequest
	body.customer_id = customer_id
	body.currency = currency
	resp := &models.CreateAccountResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}
