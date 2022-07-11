package maplerad

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/wirepay/maplerad-go/utils"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	SandboxUrl = "https://sandbox.api.maplerad.com"
	LiveUrl    = "https://api.maplerad.com"
)

type service struct {
	client *Client
}

// Client manages communication with the Maplerad API
type Client struct {
	common  service
	c       *http.Client
	baseURL *url.URL

	secret string

	Customer     *CustomerService
	Bills        *BillsService
	Collections  *CollectionsService
	Transfer     *TransferService
	Issuing      *IssuingService
	Transaction  *TransactionsService
	Misc         *MiscService
	Wallet       *WalletService
	Counterparty *CounterpartyService
}

func IsStringEmpty(s string) bool { return len(strings.TrimSpace(s)) == 0 }

type xAPIKeyAuthTransport struct {
	originalTransport http.RoundTripper
	secret            string
}

func (c *xAPIKeyAuthTransport) RoundTrip(r *http.Request) (*http.Response, error) {
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("Authorization", "Bearer "+c.secret)
	return c.originalTransport.RoundTrip(r)
}

// NewClient /* Environment is either live or sandbox */
func NewClient(secret, environment string) (*Client, error) {
	c := &Client{}

	if IsStringEmpty(secret) {
		return nil, errors.New("please provide your secret key")
	}
	c.secret = secret
	if environment == "live" {
		c.baseURL, _ = url.Parse(LiveUrl)
	}
	c.baseURL, _ = url.Parse(SandboxUrl)
	c.common.client = c

	c.Customer = (*CustomerService)(&c.common)
	c.Bills = (*BillsService)(&c.common)
	c.Collections = (*CollectionsService)(&c.common)
	c.Transfer = (*TransferService)(&c.common)
	c.Issuing = (*IssuingService)(&c.common)
	c.Misc = (*MiscService)(&c.common)
	c.Wallet = (*WalletService)(&c.common)
	c.Counterparty = (*CounterpartyService)(&c.common)
	c.Transaction = (*TransactionsService)(&c.common)

	if c.c == nil {
		c.c = &http.Client{
			Transport: &xAPIKeyAuthTransport{
				originalTransport: http.DefaultTransport,
				secret:            c.secret,
			},
			Timeout: time.Second * 10,
		}
	}
	return c, nil
}

// Call actually does the HTTP request to Maplerad API
func (c *Client) Call(method, path string, queryParams url.Values, body, v interface{}) error {
	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return err
		}
	}
	u, err := c.baseURL.Parse("v1" + path)
	if err != nil {
		return err
	}

	// Adding Query Param
	query := u.Query()
	for k, v := range queryParams {
		for _, iv := range v {
			query.Add(k, iv)
		}
	}

	/* TODO*/
	// Encode the parameters.
	// u.RawQuery() = query.Encode()

	req, err := http.NewRequest(method, u.String(), buf)
	log.Print(method, ": ", req.URL)

	if err != nil {
		return err
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	resp, err := c.c.Do(req)
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
		}
	}(resp.Body)
	return c.decodeResponse(resp, v)
}

// Response represents arbitrary response data
type Response map[string]interface{}

// // decodeResponse decodes the JSON response from the Maplerad API.
// // The actual response will be written to the `v` parameter
func (c *Client) decodeResponse(httpResp *http.Response, v interface{}) error {
	if httpResp.StatusCode >= 400 {
		return utils.NewAPIError(httpResp)
	}
	respBody, err := ioutil.ReadAll(httpResp.Body)
	err = json.Unmarshal(respBody, &v)
	if err != nil {
		return err
	}

	return nil
}
