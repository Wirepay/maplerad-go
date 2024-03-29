package models

type GetCardResponse struct {
	Generic
	Data struct {
		Name        string `json:"name"`
		MaskedPan   string `json:"masked_pan"`
		Expiry      string `json:"expiry"`
		Cvv         string `json:"cvv"`
		Type        string `json:"type"`
		Issuer      string `json:"issuer"`
		Currency    string `json:"currency"`
		Balance     int    `json:"balance"`
		AutoApprove bool   `json:"auto_approve"`
	} `json:"data"`
}

type GetCardTransactionsResponse struct {
	Generic
	Data []struct {
		ID          string `json:"id"`
		Amount      int    `json:"amount"`
		Currency    string `json:"currency"`
		Description string `json:"description"`
		Status      string `json:"status"`
		CreatedAt   string `json:"created_at"`
	} `json:"data"`
	Meta `json:"meta"`
}

type GetAllCardsResponse []struct {
	Generic
	Data struct {
		Name        string `json:"name"`
		MaskedPan   string `json:"masked_pan"`
		Expiry      string `json:"expiry"`
		Cvv         string `json:"cvv"`
		Type        string `json:"type"`
		Issuer      string `json:"issuer"`
		Currency    string `json:"currency"`
		Balance     int    `json:"balance"`
		AutoApprove bool   `json:"auto_approve"`
	} `json:"data"`
}

type CreateCardResponse struct {
	Generic
	Data struct {
		Id          string `json:"id"`
		Name        string `json:"name"`
		MaskedPan   string `json:"masked_pan"`
		Expiry      string `json:"expiry"`
		Cvv         string `json:"cvv"`
		Type        string `json:"type"`
		Issuer      string `json:"issuer"`
		Currency    string `json:"currency"`
		Balance     int    `json:"balance"`
		AutoApprove bool   `json:"auto_approve"`
	} `json:"data"`
}
