package maplerad

import (
	"github.com/wirepay/maplerad-go/models"
)

type BillsService service

type BuyAirtimeRequest struct {
	Phone_number string
	Amount       int
	Identifier   string
}

func (b *BillsService) BuyAirtime(body *BuyAirtimeRequest) (*models.BuyAirtimeResponse, error) {
	u := "/bills/airtime"
	resp := &models.BuyAirtimeResponse{}
	err := b.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (b *BillsService) GetAirtimeBillers() (*models.GetAirtimeBillersResponse, error) {
	u := "bills/airtime/billers/NG"
	resp := &models.GetAirtimeBillersResponse{}
	err := b.client.Call("POST", u, nil, nil, &resp)
	return resp, err
}
