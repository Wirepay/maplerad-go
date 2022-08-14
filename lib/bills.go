package maplerad

import (
	"github.com/wirepay/maplerad-go/models"
)

type BillsService service

type BuyAirtimeRequest struct {
	PhoneNumber string `json:"phone_number"`
	Amount      int    `json:"amount"`
	Identifier  string `json:"identifier"`
}

func (b *BillsService) BuyAirtime(body *BuyAirtimeRequest) (*models.BuyAirtimeResponse, error) {
	u := "/bills/airtime"
	resp := &models.BuyAirtimeResponse{}
	err := b.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (b *BillsService) GetBillers() (*models.GetAirtimeBillersResponse, error) {
	u := "/bills/airtime/billers/NG"
	resp := &models.GetAirtimeBillersResponse{}
	err := b.client.Call("POST", u, nil, nil, &resp)
	return resp, err
}
