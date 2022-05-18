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
	err := c.client.Call("POST", u, nil, nil, &resp)
	return resp, err
}

// func (c *IssuingService) GetCardTransactions(customer_id string) (*models.GetCardResponse, error) {
// 	u := fmt.Sprintf("/issuing/%s", customer_id)
// 	resp := &models.GetCardResponse{}
// 	err := c.client.Call("GET", u, nil, nil, &resp)
// 	return resp, err
// }
