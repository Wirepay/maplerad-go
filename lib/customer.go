package maplerad

import (
	"fmt"
	"github.com/wirepay/maplerad-go/utils"

	"github.com/wirepay/maplerad-go/models"
)

type CustomerService service

type FullEnrollRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Phone     struct {
		PhoneCountryCode string `json:"phone_country_code"`
		PhoneNumber      string `json:"phone_number"`
	} `json:"phone"`
	Dob                  string `json:"dob"`
	IdentificationNumber string `json:"identification_number"`
	Address              struct {
		Street     string  `json:"street"`
		Street2    *string `json:"street2"`
		City       string  `json:"city"`
		State      string  `json:"state"`
		Country    string  `json:"country"`
		PostalCode string  `json:"postal_code"`
	} `json:"address"`
	Identity struct {
		Country string `json:"country"`
		Image   string `json:"image"`
		Number  string `json:"number"`
		Type    string `json:"type"`
	} `json:"identity"`
}

type CreateCustomerRequestTier0 struct {
	FirstName string `json:"first_name" binding:"required"`
	LastName  string `json:"last_name" binding:"required"`
	Email     string `json:"email" binding:"required"`
	Country   string `json:"country" binding:"required"`
}

type CreateCustomerRequestTier1 struct {
	CustomerId string `json:"customer_id" binding:"required"`
	Phone      struct {
		PhoneCountryCode string `json:"phone_country_code"`
		PhoneNumber      string `json:"phone_number"`
	} `json:"phone"`
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
	CustomerId string `json:"customer_id" binding:"required"`
	Identity   struct {
		Type    string `json:"type" binding:"required"`
		Image   string `json:"image" binding:"required"`
		Url     string `json:"url" binding:"required"`
		Country string `json:"country" binding:"required"`
	} `json:"identity" binding:"required"`
}

type CardEnrollRequest struct {
	CustomerId string          `json:"customer_id"`
	Brand      utils.CardBrand `json:"brand"`
}

type CustomerUpdateRequest struct {
	CustomerId string `json:"customer_id"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	MiddleName string `json:"middle_name"`
	Photo      string `json:"photo"`
	Phone      struct {
		PhoneCountryCode string `json:"phone_country_code"`
		PhoneNumber      string `json:"phone_number"`
	} `json:"phone"`
	Identity struct {
		Type    string `json:"type"`
		Image   string `json:"image"`
		Url     string `json:"url"`
		Country string `json:"country"`
	} `json:"identity"`
}

func (c *CustomerService) FullEnrollCustomer(body *FullEnrollRequest) (*models.FullEnrollResponse, error) {
	u := "/customers/enroll"
	resp := &models.FullEnrollResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *CustomerService) CreateCustomer(body *CreateCustomerRequestTier0) (*models.CreateCustomerResponseTier0, error) {
	u := "/customers"
	resp := &models.CreateCustomerResponseTier0{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *CustomerService) UpgradeCustomerTierOne(body *CreateCustomerRequestTier1) (*models.CreateCustomerResponseTier1, error) {
	u := "/customers/upgrade/tier1"
	resp := &models.CreateCustomerResponseTier1{}
	err := c.client.Call("PATCH", u, nil, body, &resp)
	return resp, err
}

func (c *CustomerService) UpgradeCustomerTierTwo(body *CreateCustomerRequestTier2) (*models.CreateCustomerResponseTier2, error) {
	u := "/customers/upgrade/tier2"
	resp := &models.CreateCustomerResponseTier2{}
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
	u := fmt.Sprintf("/customers/%s/accounts", customerId)
	resp := &models.GetCustomerAccountResponse{}
	err := c.client.Call("POST", u, nil, nil, &resp)
	return resp, err
}

func (c *CustomerService) GetCustomerVirtualAccounts(customerId string) (*models.GetCustomerAccountResponse, error) {
	u := fmt.Sprintf("/customers/%s/virtual-account", customerId)
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

func (c *CustomerService) CustomerCardEnroll(customerId string, brand utils.CardBrand) (*models.CardEnrollResponse, error) {
	u := "/customers/card-enroll"
	body := CardEnrollRequest{
		CustomerId: customerId,
		Brand:      brand,
	}
	resp := &models.CardEnrollResponse{}
	err := c.client.Call("PATCH", u, nil, body, &resp)
	return resp, err
}

func (c *CustomerService) UpdateCustomer(body *CustomerUpdateRequest) (*models.Generic, error) {
	u := "/customers/update"
	resp := &models.Generic{}
	err := c.client.Call("PATCH", u, nil, body, &resp)
	return resp, err
}
