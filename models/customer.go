package models

import "time"

type CreateCustomerResponseTier0 struct {
	Generic
	Data struct {
		ID        string `json:"id"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Country   string `json:"country"`
		Status    string `json:"status"`
		Tier      int    `json:"tier"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	} `json:"data"`
}

type CreateCustomerResponseTier1 struct {
	Generic
}

type CreateCustomerResponseTier2 struct {
	Generic
	Data struct {
		Id     string `json:"id"`
		Status string `json:"status"`
	} `json:"data"`
}

type GetCustomerResponse struct {
	Generic
	Data struct {
		Id          string `json:"id"`
		FirstName   string `json:"first_name"`
		LastName    string `json:"last_name"`
		MiddleName  string `json:"middle_name"`
		Email       string `json:"email"`
		PhoneNumber string `json:"phone_number"`
		Dob         string `json:"dob"`
		Active      bool   `json:"active"`
		Type        string `json:"type"`
		Disabled    bool   `json:"disabled"`
		Identity    struct {
			Type    string `json:"type"`
			Number  string `json:"number"`
			Image   string `json:"image"`
			Country string `json:"country"`
		} `json:"identity"`
	}
	Address struct {
		Street     string `json:"street"`
		Street2    string `json:"street2"`
		City       string `json:"city"`
		State      string `json:"state"`
		Country    string `json:"country"`
		PostalCode string `json:"postal_code"`
	} `json:"address"`
	Status string `json:"status"`
}

type GetCustomersResponse struct {
	Generic
	Data []struct {
		ID          string      `json:"id"`
		FirstName   string      `json:"first_name"`
		LastName    string      `json:"last_name"`
		MiddleName  interface{} `json:"middle_name"`
		Email       string      `json:"email"`
		PhoneNumber string      `json:"phone_number"`
		Dob         string      `json:"dob"`
		Type        string      `json:"type"`
		Active      bool        `json:"active"`
		Disabled    bool        `json:"disabled"`
		Identity    struct {
			Type    string      `json:"type"`
			Number  string      `json:"number"`
			Image   interface{} `json:"image"`
			Country string      `json:"country"`
		} `json:"identity"`
		Address struct {
			Street     string      `json:"street"`
			Street2    interface{} `json:"street2"`
			City       string      `json:"city"`
			State      string      `json:"state"`
			PostalCode string      `json:"postal_code"`
			Country    string      `json:"country"`
		} `json:"address"`
		Status string `json:"status"`
	} `json:"data"`
}

type GetCustomerAccountResponse struct {
	Generic
	Data []struct {
		Id            string `json:"id"`
		BankName      string `json:"bank_name"`
		AccountNumber string `json:"account_number"`
		AccountName   string `json:"account_name"`
		Currency      string `json:"currency"`
		CreatedAt     string `json:"created_at"`
	} `json:"data"`
}

type FullEnrollResponse struct {
	Generic
	Data struct {
		ID        string    `json:"id"`
		FirstName string    `json:"first_name"`
		LastName  string    `json:"last_name"`
		Email     string    `json:"email"`
		Country   string    `json:"country"`
		Status    string    `json:"status"`
		Tier      int       `json:"tier"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	} `json:"data"`
}
type GetCustomerCardsResponse struct {
}

type GetCustomerTransactionsResponse struct {
}

type CardEnrollResponse struct {
	Generic
}
