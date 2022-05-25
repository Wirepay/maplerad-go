package maplerad

import (
	"fmt"

	"maplerad/models"
)

type CustomerService service

type CreateCustomerRequest struct {
	First_name   string
	Last_name    string
	Email        string
	Phone_number string
	Dob          string
	Identity     struct {
		Type    string
		Number  string
		Url     string
		Country string
	}
	Address struct {
		Street      string
		Street2     string
		City        string
		State       string
		Country     string
		Postal_code string
	}
}

func (c *CustomerService) CreateCustomer(body *CreateCustomerRequest) (*models.CreateCustomerResponse, error) {
	u := "/customers"
	resp := &models.CreateCustomerResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}

func (c *CustomerService) GetCustomer(customer_id string) (*models.GetCustomerResponse, error) {
	u := fmt.Sprintf("/customers/%s", customer_id)
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
