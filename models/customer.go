package models

type CreateCustomerResponse struct {
	Generic
	Data struct {
		Id     string
		Status string
	}
}

type GetCustomerResponse struct {
	Generic
	Data struct {
		Id           string `json:"id"`
		First_name   string `json:"first_name"`
		Last_name    string `json:"last_name"`
		Middle_name  string `json:"middle_name"`
		Email        string `json:"email"`
		Phone_number string `json:"phone_number"`
		Dob          string `json:"dob"`
		Active       bool   `json:"active"`
		Disabled     bool   `json:"disabled"`
		Identity     struct {
			Type    string `json:"type"`
			Number  string `json:"number"`
			Image   string `json:"image"`
			Country string `json:"country"`
		} `json:"identity"`
	}
	Address struct {
		Street      string `json:"street"`
		Street2     string `json:"street2"`
		City        string `json:"city"`
		State       string `json:"state"`
		Country     string `json:"country"`
		Postal_code string `json:"postal_code"`
	}
}

type GetCustomersResponse struct {
	Generic
	Data []struct {
		Id           string `json:"id"`
		First_name   string `json:"first_name"`
		Last_name    string `json:"last_name"`
		Middle_name  string `json:"middle_name"`
		Email        string `json:"email"`
		Phone_number string `json:"phone_number"`
		Dob          string `json:"dob"`
		Active       bool   `json:"active"`
		Disabled     bool   `json:"disabled"`
		Identity     struct {
			Type    string `json:"type"`
			Number  string `json:"number"`
			Image   string `json:"image"`
			Country string `json:"country"`
		} `json:"identity"`
	}
	Address struct {
		Street      string `json:"street"`
		Street2     string `json:"street2"`
		City        string `json:"city"`
		State       string `json:"state"`
		Country     string `json:"country"`
		Postal_code string `json:"postal_code"`
	}
}
