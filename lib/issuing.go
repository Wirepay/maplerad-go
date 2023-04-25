package maplerad

import (
	"fmt"
	"github.com/wirepay/maplerad-go/models"
	"github.com/wirepay/maplerad-go/utils"
	"net/url"
)

type IssuingService service

type CreateCardRequest struct {
	CustomerId  string               `json:"customer_id"`
	Currency    utils.CurrencySymbol `json:"currency"`
	Type        string               `json:"type"`
	Brand       utils.CardBrand      `json:"brand"`
	Amount      int64                `json:"amount"`
	CardPin     string               `json:"card_pin"`
	AutoApprove bool                 `json:"auto_approve"`
}

type CreateBusinessCardRequest struct {
	Name        string               `json:"name"`
	Currency    utils.CurrencySymbol `json:"currency"`
	Type        string               `json:"type"`
	Brand       utils.CardBrand      `json:"brand"`
	Amount      string               `json:"amount"`
	AutoApprove bool                 `json:"auto_approve"`
}

type CardRequest struct {
	Amount string
}

type CardPinRequest struct {
	CardPin string `json:"card_pin"`
}

func (c *IssuingService) GetCard(cardId string) (*models.GetCardResponse, error) {
	u := fmt.Sprintf("/issuing/%s", cardId)
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

func (c *IssuingService) CreateCard(body *CreateCardRequest) (*models.CreateCardResponse, error) {
	u := "/issuing"
	resp := &models.CreateCardResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *IssuingService) CreateBusinessCard(body *CreateBusinessCardRequest) (*models.CreateCardResponse, error) {
	u := "/issuing"
	resp := &models.CreateCardResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *IssuingService) GetCardTransactions(cardId string, query *QueryParams) (*models.GetCardTransactionsResponse, error) {
	u := fmt.Sprintf("/issuing/%s/transactions", cardId)
	q := url.Values{}
	if query.Page != "" {
		q.Add("page", query.Page)
	}
	if query.PageSize != "" {
		q.Add("page_size", query.PageSize)
	}
	if query.StartDate != "" {
		q.Add("start_date", query.StartDate)
	}
	if query.EndDate != "" {
		q.Add("end_date", query.EndDate)
	}
	resp := &models.GetCardTransactionsResponse{}
	err := c.client.Call("GET", u, q, nil, &resp)
	return resp, err
}

func (c *IssuingService) FundCard(cardId string, amount string) (*models.Generic, error) {
	u := fmt.Sprintf("/issuing/%s/fund", cardId)
	body := CardRequest{Amount: amount}
	resp := &models.Generic{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *IssuingService) DebitCard(cardId string, amount string) (*models.Generic, error) {
	u := fmt.Sprintf("/issuing/%s/debit", cardId)
	body := CardRequest{Amount: amount}
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

func (c *IssuingService) SetCardPin(cardId string, cardPin string) (*models.Generic, error) {
	u := fmt.Sprintf("/issuing/%s/set-pin", cardId)
	body := CardPinRequest{CardPin: cardPin}
	resp := &models.Generic{}
	err := c.client.Call("PATCH", u, nil, body, &resp)
	return resp, err
}
