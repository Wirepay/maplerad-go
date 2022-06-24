package maplerad

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"maplerad/utils"
	"net/http"
	"net/url"
	"strings"
	"time"

	
	"github.com/mitchellh/mapstructure"
)

const (
	SANDBOX_URL = "https://api.sandbox.maplerad.com/v1"
	LIVE_URL    = "https://api.maplerad.com/v1"
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

	Customer *CustomerService
	Bills    *BillsService
	Account  *AccountService
	Transfer *TransferService
	Issuing  *IssuingService
	Misc     *MiscService
}

// func WithHTTPClient(cl *http.Client) Option {
// 	return func(c *Client) {
// 		c.c = cl
// 	}
// }

// func WithSecretKey(s string) Option {
// 	return func(c *Client) {
// 		c.secret = s
// 	}
// }

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

/* Environment is either live or sandbox */
func NewClient(secret, environment string) (*Client, error) {
	c := &Client{}

	if IsStringEmpty(secret) {
		return nil, errors.New("please provide your secret key")
	}
	c.secret = secret
	if environment == "live" {
		c.baseURL, _ = url.Parse(LIVE_URL)
	}
	c.baseURL, _ = url.Parse(SANDBOX_URL)
	c.common.client = c

	c.Customer = (*CustomerService)(&c.common)
	c.Bills = (*BillsService)(&c.common)
	c.Account = (*AccountService)(&c.common)
	c.Transfer = (*TransferService)(&c.common)
	c.Issuing = (*IssuingService)(&c.common)
	c.Misc = (*MiscService)(&c.common)

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

// // Call actually does the HTTP request to Maplerad API
func (c *Client) Call(method, path string, queryParams url.Values, body, v interface{}) error {
	var buf io.ReadWriter

	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return err
		}
	}
	u, err := c.baseURL.Parse(path)
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

	defer resp.Body.Close()
	return c.decodeResponse(resp, v)
}

// u, _ := c.baseURL.Parse(r.path)

func mapstruct(data interface{}, v interface{}) error {
	config := &mapstructure.DecoderConfig{
		Result:           v,
		TagName:          "json",
		WeaklyTypedInput: true,
	}
	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}
	err = decoder.Decode(data)
	return err
}

// Response represents arbitrary response data
type Response map[string]interface{}

// // decodeResponse decodes the JSON response from the Maplerad API.
// // The actual response will be written to the `v` parameter
func (c *Client) decodeResponse(httpResp *http.Response, v interface{}) error {
	var resp Response
	respBody, err := ioutil.ReadAll(httpResp.Body)
	json.Unmarshal(respBody, &resp)
	if err != nil {
		return err
	}
	if status, _ := resp["status"].(bool); !status || httpResp.StatusCode >= 400 {
		return utils.NewAPIError(httpResp)
	}

	if data, ok := resp["data"]; ok {
		switch t := resp["data"].(type) {
		case map[string]interface{}:
			return mapstruct(data, v)
		default:
			_ = t
			return mapstruct(resp, v)
		}
	}
	// if response data does not contain data key, map entire response to v
	return mapstruct(resp, v)
}
