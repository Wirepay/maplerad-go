package maplerad

import (
	"fmt"
	"maplerad/models"
)

type IssuingService service

func (c *IssuingService) GetCard(customer_id string) (*models.GetCardResponse, error) {
	u := fmt.Sprintf("/issuing/%s", customer_id)
	resp := &models.GetCardResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}

func (c *IssuingService) GetAllCards() (*models.GetAllCardsResponse, error) {
	u := "/issuing"
	resp := &models.GetAllCardsResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}

type CreateCardRequest struct {
	CustomerId  string
	Currency    string
	Type        string
	AutoApprove bool
}

func (c *IssuingService) CreateCard(body *CreateCardRequest) (*models.CreateCardResponse, error) {
	u := "/issuing"
	resp := &models.CreateCardResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *IssuingService) GetCardTransactions(customer_id string) (*models.GetCardTransactionsResponse, error) {
	u := fmt.Sprintf("/issuing/%s", customer_id)
	resp := &models.GetCardTransactionsResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}

type CardRequest struct {
	amount string
}

func (c *IssuingService) FundCard(card_id, amount string) (*models.Generic, error) {
	u := fmt.Sprintf("/issuing/%s/fund", card_id)
	var body *CardRequest
	body.amount = amount
	resp := &models.Generic{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *IssuingService) DebitCard(card_id, amount string) (*models.Generic, error) {
	u := fmt.Sprintf("/issuing/%s/debit", card_id)
	var body *CardRequest
	body.amount = amount
	resp := &models.Generic{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *IssuingService) EnableCard(card_id string) (*models.Generic, error) {
	u := fmt.Sprintf("/issuing/%s/enable", card_id)
	resp := &models.Generic{}
	err := c.client.Call("PATCH", u, nil, nil, &resp)
	return resp, err
}

func (c *IssuingService) DisableCard(card_id string) (*models.Generic, error) {
	u := fmt.Sprintf("/issuing/%s/disable", card_id)
	resp := &models.Generic{}
	err := c.client.Call("PATCH", u, nil, nil, &resp)
	return resp, err
}
