package maplerad

import (
	"fmt"
	"github.com/wirepay/maplerad-go/models"
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
func (c *WalletService) GetWalletHistory() (*models.GetWalletHistoryResponse, error) {
	u := "/wallets/history"
	resp := &models.GetWalletHistoryResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}

func (c *WalletService) GetWalletHistoryByCurrency(currencyCode string) (*models.GetWalletHistoryByCurrencyResponse, error) {
	u := fmt.Sprintf("/wallets/%s/history", currencyCode)
	resp := &models.GetWalletHistoryByCurrencyResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}
