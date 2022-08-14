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
