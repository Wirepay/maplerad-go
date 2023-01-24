package maplerad

import (
	"fmt"
	"github.com/wirepay/maplerad-go/models"
	"net/url"
)

type WalletService service

func (c *WalletService) GetWallets() (*models.GetWalletsResponse, error) {
	u := "/wallets"
	resp := &models.GetWalletsResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}

// GetWalletByCurrency CurrencyCode is either NGN,USD etc
func (c *WalletService) GetWalletByCurrency(currencyCode string) (*models.GetWalletByCurrencyResponse, error) {
	u := fmt.Sprintf("/wallets/%s", currencyCode)

	resp := &models.GetWalletByCurrencyResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}
func (c *WalletService) GetWalletHistory(query *QueryParams) (*models.GetWalletHistoryResponse, error) {
	u := "/wallets/history"
	resp := &models.GetWalletHistoryResponse{}
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

	err := c.client.Call("GET", u, q, nil, &resp)
	return resp, err
}

// GetWalletHistoryByCurrency CurrencyCode is either NGN,USD etc
func (c *WalletService) GetWalletHistoryByCurrency(currencyCode string, query *QueryParams) (*models.GetWalletHistoryByCurrencyResponse, error) {
	u := fmt.Sprintf("/wallets/%s/history", currencyCode)
	resp := &models.GetWalletHistoryByCurrencyResponse{}
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

	err := c.client.Call("GET", u, q, nil, &resp)
	return resp, err
}
