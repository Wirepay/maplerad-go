package maplerad

import (
	"github.com/wirepay/maplerad-go/models"
)

type MiscService service

func (c *MiscService) GetCurrencies() (*models.GetCurrenciesResponse, error) {
	u := "/currencies"
	resp := &models.GetCurrenciesResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}
