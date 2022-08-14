package maplerad

import (
	"fmt"
	"github.com/wirepay/maplerad-go/models"
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
	CustomerId  string `json:"customer_id"`
	Currency    string `json:"currency"`
	Type        string `json:"type"`
	AutoApprove bool   `json:"auto_approve"`
}

func (c *IssuingService) CreateCard(body *CreateCardRequest) (*models.CreateCardResponse, error) {
	u := "/issuing"
	resp := &models.CreateCardResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *IssuingService) GetCardTransactions(customerId string) (*models.GetCardTransactionsResponse, error) {
	u := fmt.Sprintf("/issuing/%s", customerId)
	resp := &models.GetCardTransactionsResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}

type CardRequest struct {
	amount string
}

func (c *IssuingService) FundCard(cardId string, amount string) (*models.Generic, error) {
	u := fmt.Sprintf("/issuing/%s/fund", cardId)
	var body *CardRequest
	body.amount = amount
	resp := &models.Generic{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *IssuingService) DebitCard(cardId string, amount string) (*models.Generic, error) {
	u := fmt.Sprintf("/issuing/%s/debit", cardId)
	var body *CardRequest
	body.amount = amount
	resp := &models.Generic{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *IssuingService) UnFreezeCard(cardId string) (*models.Generic, error) {
	u := fmt.Sprintf("/issuing/%s/unfreeze", cardId)
	resp := &models.Generic{}
	err := c.client.Call("PATCH", u, nil, nil, &resp)
	return resp, err
}

func (c *IssuingService) FreezeCard(cardId string) (*models.Generic, error) {
	u := fmt.Sprintf("/issuing/%s/freeze", cardId)
	resp := &models.Generic{}
	err := c.client.Call("PATCH", u, nil, nil, &resp)
	return resp, err
}
