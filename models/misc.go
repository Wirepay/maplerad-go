package models

type GetMarketQuoteResponse struct {
}
type GetCurrenciesResponse struct {
	Generic
	Data []struct {
		Name     string `json:"name"`
		Currency string `json:"currency"`
		Symbol   string `json:"symbol"`
	}
}
