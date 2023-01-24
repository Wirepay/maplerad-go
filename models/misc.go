package models

type GetMarketQuoteResponse struct {
	Generic
	Data struct {
		Reference string `json:"reference"`
		Source    struct {
			Currency            string `json:"currency"`
			Amount              int    `json:"amount"`
			HumanReadableAmount int    `json:"human_readable_amount"`
		} `json:"source"`
		Target struct {
			Currency            string `json:"currency"`
			Amount              int    `json:"amount"`
			HumanReadableAmount int    `json:"human_readable_amount"`
		} `json:"target"`
		Rate int `json:"rate"`
	} `json:"data"`
}
type GetCurrenciesResponse struct {
	Generic
	Data []struct {
		Name     string `json:"name"`
		Currency string `json:"currency"`
		Symbol   string `json:"symbol"`
	} `json:"data"`
}

type GetCountriesResponse struct {
	Generic
	Data []struct {
		Name        string `json:"name"`
		Code        string `json:"code"`
		CallingCode string `json:"calling_code"`
	}
}

type BVNResponse struct {
	FirstName          string `json:"first_name"`
	LastName           string `json:"last_name"`
	MiddleName         string `json:"middle_name"`
	Title              string `json:"title"`
	Email              string `json:"email"`
	Gender             string `json:"gender"`
	Dob                string `json:"dob"`
	PhoneNumber        string `json:"phone_number"`
	ResidentialAddress string `json:"residential_address"`
	EnrollmentBank     string `json:"enrollment_bank"`
	EnrollmentBranch   string `json:"enrollment_branch"`
	Tier               string `json:"tier"`
	LgaOfOrigin        string `json:"lga_of_origin"`
	StateOfOrigin      string `json:"state_of_origin"`
	LgaOfResidence     string `json:"lga_of_residence"`
	StateOfResidence   string `json:"state_of_residence"`
	MaritalStatus      string `json:"marital_status"`
	Nationality        string `json:"nationality"`
	Image              string `json:"image"`
}

type GetInstitutionsResponse struct {
	Generic
	Data []struct {
		Name string `json:"name"`
		Code string `json:"code"`
	} `json:"data"`
	Page     int `json:"page"`
	PageSize int `json:"page_size"`
	Total    int `json:"total"`
}

type ResolveInstitutionsResponse struct {
	Generic
	Data struct {
		AccountName   string `json:"account_name"`
		AccountNumber string `json:"account_number"`
	} `json:"data"`
}
