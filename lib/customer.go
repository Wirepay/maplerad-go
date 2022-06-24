package maplerad

import (
	"fmt"

	"gihtub.com/wirepay/maplerad-go/models"
)

type CustomerService service

type CreateCustomerRequest struct {
	FirstName   string
	LastName    string
	Email       string
	PhoneNumber string
	Dob         string
	Identity    struct {
		Type    string
		Number  string
		Url     string
		Country string
	}
	Address struct {
		Street     string
		Street2    string
		City       string
		State      string
		Country    string
		PostalCode string
	}
}

func (c *CustomerService) CreateCustomer(body *CreateCustomerRequest) (*models.CreateCustomerResponse, error) {
	u := "/customers"
	resp := &models.CreateCustomerResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *CustomerService) GetCustomer(customerId string) (*models.GetCustomerResponse, error) {
	u := fmt.Sprintf("/customers/%s", customerId)
	resp := &models.GetCustomerResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}

func (c *CustomerService) GetCustomers() (*models.GetCustomersResponse, error) {
	u := "/customers"
	resp := &models.GetCustomersResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}

func (c *CustomerService) GetCustomerAccount(customerId string) (*models.GetCustomerAccountResponse, error) {
	u := fmt.Sprintf("/customers/%s/account", customerId)
	resp := &models.GetCustomerAccountResponse{}
	err := c.client.Call("POST", u, nil, nil, &resp)
	return resp, err
}

func (c *CustomerService) GetCustomerCards(customerId string) (*models.GetCustomerCardsResponse, error) {
	u := fmt.Sprintf("/customers/%s/cards", customerId)
	resp := &models.GetCustomerCardsResponse{}
	err := c.client.Call("POST", u, nil, nil, &resp)
	return resp, err
}

func (c *CustomerService) GetCustomerTransactions(customerId string) (*models.GetCustomerTransactionsResponse, error) {
	u := fmt.Sprintf("/customers/%s/transactions", customerId)
	resp := &models.GetCustomerTransactionsResponse{}
	err := c.client.Call("POST", u, nil, nil, &resp)
	return resp, err
}
