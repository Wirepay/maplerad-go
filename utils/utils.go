package utils

import (
	"encoding/json"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"net/url"
)

type CurrencySymbol string
type CardBrand string

const (
	USD CurrencySymbol = "USD"
	NGN CurrencySymbol = "NGN"

	VISA       CardBrand = "VISA"
	MASTERCARD CardBrand = "MASTERCARD"
)

// PrettyPrint Convert string or interface{} to JSON
func PrettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

// APIError includes the response from the Maplerad API and some HTTP request info
type APIError struct {
	Message        string        `json:"message,omitempty"`
	HTTPStatusCode int           `json:"code,omitempty"`
	Details        ErrorResponse `json:"details,omitempty"`
	URL            *url.URL      `json:"url,omitempty"`
	Header         http.Header   `json:"header,omitempty"`
}

// APIError supports the error interface
func (err *APIError) Error() string {
	ret, _ := json.Marshal(err)
	return string(ret)
}

// ErrorResponse represents an error response from the Maplerad API server
type ErrorResponse struct {
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	//Errors  map[string]interface{} `json:"errors,omitempty"`
}

func NewAPIError(resp *http.Response) *APIError {
	p, _ := ioutil.ReadAll(resp.Body)

	var mapleradErrorResp ErrorResponse
	_ = json.Unmarshal(p, &mapleradErrorResp)
	return &APIError{
		HTTPStatusCode: resp.StatusCode,
		Details:        mapleradErrorResp,
		Header:         resp.Header,
	}

}

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
