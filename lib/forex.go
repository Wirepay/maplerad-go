package maplerad

import (
	"github.com/wirepay/maplerad-go/models"
)

type GetMarketQuoteRequest struct {
	sourceCurrency string
	targetCurrency string
	amount         int
}

func (c *MiscService) GetMarketQuote(body *GetMarketQuoteRequest) (*models.GetMarketQuoteResponse, error) {
	u := "/fx/quote"
	resp := &models.GetMarketQuoteResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *MiscService) ExchangeCurrency(quoteReference string) (Response, error) {
	u := "/fx"
	resp := Response{}
	err := c.client.Call("POST", u, nil, quoteReference, &resp)
	return resp, err
}
