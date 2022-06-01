package maplerad

import (
	"maplerad/models"
)

type GetMarketQuoteRequest struct {
	source_currency string
	target_currency string
	amount          int
}

func (c *MiscService) GetMarketQuote(source, target string, amount int) (*models.GetMarketQuoteResponse, error) {
	u := "/fx"
	var body *GetMarketQuoteRequest
	body.source_currency = source
	body.target_currency = target
	body.amount = amount
	resp := &models.GetMarketQuoteResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}
