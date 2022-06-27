package maplerad

import "github.com/wirepay/maplerad-go/models"

type SettlementService service

type SettlementRequest struct {
	accountNumber string
	bankCode      string
	accountName   string
	currency      string
}

func (c *SettlementService) GetAllSettlements() (*models.SettlementResponse, error) {
	u := "/settlements"
	resp := &models.SettlementResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}

func (c *SettlementService) MakeSettlement(body *SettlementRequest) (*models.SettlementResponse, error) {
	u := "/settlements"
	resp := &models.SettlementResponse{}
	err := c.client.Call("POST", u, nil, nil, &resp)
	return resp, err
}

func (c *SettlementService) DeleteSettlement() (Response, error) {
	u := "/settlements"
	resp := Response{}
	err := c.client.Call("DELETE", u, nil, nil, &resp)
	return resp, err
}
