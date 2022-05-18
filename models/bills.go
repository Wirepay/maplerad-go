package models

type Generic struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type BuyAirtimeResponse struct {
	Generic
	Data struct {
		Id               string `json:"id"`
		Amount           int    `json:"amount"`
		PhoneNumber      string `json:"phone_number"`
		Network          string `json:"network"`
		DebitAmount      int    `json:"debit_amount"`
		CommissionEarned int    `json:"commission_earned"`
	} `json:"data"`
}

type GetAirtimeBillersResponse struct {
	Generic
	Data []struct {
		Name       string `json:"name"`
		Identifier string `json:"Identifier"`
		Commission int    `json:"commission"`
	} `json:"data"`
}
