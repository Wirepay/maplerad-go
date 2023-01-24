package maplerad

import "github.com/wirepay/maplerad-go/models"

type IdentityService service

func (c *IdentityService) VerifyBVN(bvn string) (*models.BVNResponse, error) {
	u := "/identity"
	resp := &models.BVNResponse{}
	err := c.client.Call("GET", u, nil, bvn, resp)
	return resp, err
}
