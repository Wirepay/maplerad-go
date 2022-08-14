package maplerad

import (
	"fmt"

	"github.com/wirepay/maplerad-go/models"
)

type CustomerService service

type CreateCustomerRequestTier0 struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Country   string `json:"country" binding:"required"`
}
type CreateCustomerRequestTier1 struct {
	CustomerId           string `json:"customer_id"`
	PhoneNumber          string `json:"phone_number" binding:"required"`
	Dob                  string `json:"dob" binding:"required"`
	IdentificationNumber string `json:"identification_number"`
	Address              struct {
		Street     string `json:"street" binding:"required"`
		Street2    string `json:"street2"`
		City       string `json:"city" binding:"required"`
		State      string `json:"state" binding:"required"`
		Country    string `json:"country" binding:"required"`
		PostalCode string `json:"postalCode"`
	} `json:"address" binding:"required"`
}

type CreateCustomerRequestTier2 struct {
	CustomerId string `json:"customer_id"`
	Identity   struct {
		Type    string `json:"type" binding:"required"`
		Number  string `json:"number"`
		Url     string `json:"url" binding:"required"`
		Country string `json:"country" binding:"required"`
	} `json:"identity" binding:"required"`
}

func (c *CustomerService) CreateCustomer(body *CreateCustomerRequestTier0) (*models.CreateCustomerResponse, error) {
	u := "/customers"
	resp := &models.CreateCustomerResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *CustomerService) UpgradeCustomerTierOne(body *CreateCustomerRequestTier1) (*models.CreateCustomerResponse, error) {
	u := "/customers/tier1"
	resp := &models.CreateCustomerResponse{}
	err := c.client.Call("PATCH", u, nil, body, &resp)
	return resp, err
}

func (c *CustomerService) UpgradeCustomerTierTwo(body *CreateCustomerRequestTier2) (*models.CreateCustomerResponse, error) {
	u := "/customers/tier2"
	resp := &models.CreateCustomerResponse{}
	err := c.client.Call("PATCH", u, nil, body, &resp)
	return resp, err
}

func (c *CustomerService) GetCustomer(customerId string) (*models.GetCustomerResponse, error) {
	u := fmt.Sprintf("/customers/%s", customerId)
	resp := &models.GetCustomerResponse{}
	err := c.client.Call("GET", u, nil, nil, &resp)
	return resp, err
}

func (c *CustomerService) GetAllCustomers() (*models.GetCustomersResponse, error) {
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
