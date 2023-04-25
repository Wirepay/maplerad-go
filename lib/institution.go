package maplerad

import "github.com/wirepay/maplerad-go/models"

type InstitutionService service

func (c *InstitutionService) GetInstitutions() (*models.GetInstitutionsResponse, error) {
	u := "/institutions"
	resp := &models.GetInstitutionsResponse{}
	err := c.client.Call("POST", u, nil, nil, &resp)
	return resp, err
}

type resolveInstitution struct {
	AccountNumber string `json:"account_number"`
	BankCode      string `json:"bank_code"`
}

func (c *InstitutionService) ResolveInstitution(accountNumber, bankCode string) (*models.ResolveInstitutionsResponse, error) {
	u := "/institutions/resolve"
	body := resolveInstitution{
		AccountNumber: accountNumber,
		BankCode:      bankCode,
	}
	resp := &models.ResolveInstitutionsResponse{}
	err := c.client.Call("POST", u, nil, body, &resp)
	return resp, err
}
