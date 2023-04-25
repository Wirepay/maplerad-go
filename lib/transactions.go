package maplerad

import (
	"fmt"
	"github.com/wirepay/maplerad-go/models"
)

type TransactionsService service

func (c *TransactionsService) GetAllTransactions() (*models.GetAllTransactionsResponse, error) {
	u := "/transactions"
	resp := &models.GetAllTransactionsResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}

func (c *TransactionsService) VerifyCollectionTransaction(transactionId string) (*models.GetTransactionResponse, error) {
	u := fmt.Sprintf("/transactions/verify/%s", transactionId)
	resp := &models.GetTransactionResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}

func (c *TransactionsService) FetchTransaction(transactionId string) (*models.GetTransactionResponse, error) {
	u := fmt.Sprintf("/transactions/%s", transactionId)
	resp := &models.GetTransactionResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}
