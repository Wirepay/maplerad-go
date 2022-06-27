package maplerad

import (
	"fmt"
	"github.com/wirepay/maplerad-go/models"
)

type CounterpartyService service

func (c *CounterpartyService) BlacklistCounterparty(counterpartyId string, status bool) (*models.CounterpartyResponse, error) {
	u := fmt.Sprintf("/counterparties/blacklist/%s", counterpartyId)
	resp := &models.CounterpartyResponse{}
	var body map[string]bool
	body["blacklist"] = status
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *CounterpartyService) GetCounterparty(counterpartyId string) (*models.CounterpartyResponse, error) {
	u := fmt.Sprintf("/counterparties/%s", counterpartyId)
	resp := &models.CounterpartyResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}

func (c *CounterpartyService) GetAllCounterparties() (*models.CounterpartyResponse, error) {
	u := "/counterparties"
	resp := &models.CounterpartyResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}
