package maplerad

import (
	"gihtub.com/wirepay/maplerad-go/models"
)

type CollectionsService service

type CreateAccountRequest struct {
	customerId string
	currency   string
}

func (c *CollectionsService) CreateAccount(customerId, currency string) (*models.CreateAccountResponse, error) {
	u := "/collections"
	var body *CreateAccountRequest
	body.customerId = customerId
	body.currency = currency
	resp := &models.CreateAccountResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}
