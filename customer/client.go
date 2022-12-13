package customer

import (
	"context"
	"crypto/rand"
	"fmt"
	"net/http"

	"github.com/ianeinser/xendit-go"
	"github.com/ianeinser/xendit-go/utils/validator"
)

const (
	APIVersion = "2020-10-31"
)

// Client is the client used to invoke customer API.
type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

func GenerateIdempotencyKey() string {
	b := make([]byte, 8)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

/*
// CreateCustomer creates new customer
func (c *Client) CreateCustomer(data *CreateCustomerParams) (*xendit.Customer, *xendit.Error) {
	return c.CreateCustomerWithContext(context.Background(), data)
}

// CreateCustomerWithContext creates new customer
func (c *Client) CreateCustomerWithContext(ctx context.Context, data *CreateCustomerParams) (*xendit.Customer, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.Customer{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}
	if data.APIVersion != "" {
		header.Add("api-version", data.APIVersion)
	}
	if data.IdempotencyKey != "" {
		header.Add("idempotency-key", data.IdempotencyKey)
	}

	var d interface{}

	if strings.ToUpper(data.Type) == "INDIVIDUAL" {
		params := CreateCustomerIndividualParams{}
		mapstructure.Decode(data, &params)
		d = params
	}

	if strings.ToUpper(data.Type) == "BUSINESS" {
		params := CreateCustomerBusinessParams{}
		mapstructure.Decode(data, &params)
		d = params
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/customers", c.Opt.XenditURL),
		c.Opt.SecretKey,
		header,
		d,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
*/

// CreateCustomer creates new customer
func (c *Client) CreateCustomer(data *map[string]interface{}) (*xendit.Customer, *xendit.Error) {
	return c.CreateCustomerWithContext(context.Background(), data)
}

// CreateCustomerWithContext creates new customer
func (c *Client) CreateCustomerWithContext(ctx context.Context, data *map[string]interface{}) (*xendit.Customer, *xendit.Error) {

	response := &xendit.Customer{}
	header := http.Header{}

	idk := GenerateIdempotencyKey()

	if idk != "" {
		header.Add("idempotency-key", idk)
	}

	header.Add("api-version", APIVersion)

	if (*data)["ForUserID"] != nil {
		header.Add("for-user-id", fmt.Sprintf("%v", (*data)["ForUserID"]))
	}

	fmt.Printf("Header %+v\n", header)

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/customers", c.Opt.XenditURL),
		c.Opt.SecretKey,
		header,
		*data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetCustomerByReferenceID gets customer by reference ID
func (c *Client) GetCustomerByCustomerID(data *GetCustomerByCustomerIDParams) (*[]xendit.Customer, *xendit.Error) {
	return c.GetCustomerByCustomerIDWithContext(context.Background(), data)
}

// GetCustomerByReferenceIDWithContext gets customer by reference ID
func (c *Client) GetCustomerByCustomerIDWithContext(ctx context.Context, data *GetCustomerByCustomerIDParams) (*[]xendit.Customer, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := []xendit.Customer{}
	var queryString string
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}
	if data.APIVersion != "" {
		header.Add("api-version", data.APIVersion)
	}

	if data != nil {
		queryString = data.QueryString()
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/customers/%s", c.Opt.XenditURL, queryString),
		c.Opt.SecretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

// GetCustomerByReferenceID gets customer by reference ID
func (c *Client) GetCustomerByReferenceID(data *GetCustomerByReferenceIDParams) (*[]xendit.Customer, *xendit.Error) {
	return c.GetCustomerByReferenceIDWithContext(context.Background(), data)
}

// GetCustomerByReferenceIDWithContext gets customer by reference ID
func (c *Client) GetCustomerByReferenceIDWithContext(ctx context.Context, data *GetCustomerByReferenceIDParams) (*[]xendit.Customer, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := []xendit.Customer{}
	var queryString string
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}
	if data.APIVersion != "" {
		header.Add("api-version", data.APIVersion)
	}

	if data != nil {
		queryString = data.QueryString()
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/customers?%s", c.Opt.XenditURL, queryString),
		c.Opt.SecretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return &response, nil
}

/*
// UpdateCustomer updates the details on a customer
func (c *Client) UpdateCustomer(data *UpdateCustomerParams) (*xendit.Customer, *xendit.Error) {
	return c.UpdateCustomerWithContext(context.Background(), data)
}

// UpdateCustomerWithContext updates the details on a customer
func (c *Client) UpdateCustomerWithContext(ctx context.Context, data *UpdateCustomerParams) (*xendit.Customer, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.Customer{}
	header := http.Header{}
	var queryString string

	if data != nil {
		queryString = data.QueryString()
	}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}
	if data.APIVersion != "" {
		header.Add("api-version", data.APIVersion)
	}

	err := c.APIRequester.Call(
		ctx,
		"PATCH",
		fmt.Sprintf("%s/customers/%s", c.Opt.XenditURL, queryString),
		c.Opt.SecretKey,
		header,
		data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
*/

// UpdateCustomer updates the details on a customer
func (c *Client) UpdateCustomer(data *map[string]interface{}) (*xendit.Customer, *xendit.Error) {
	return c.UpdateCustomerWithContext(context.Background(), data)
}

// UpdateCustomerWithContext updates the details on a customer
func (c *Client) UpdateCustomerWithContext(ctx context.Context, data *map[string]interface{}) (*xendit.Customer, *xendit.Error) {

	response := &xendit.Customer{}
	header := http.Header{}

	queryString := ""

	if data != nil {
		queryString = fmt.Sprintf("%v", (*data)["CustomerID"])
	}

	header.Add("api-version", APIVersion)

	if (*data)["ForUserID"] != nil {
		header.Add("for-user-id", fmt.Sprintf("%v", (*data)["ForUserID"]))
	}

	fmt.Printf("Header %+v\n", header)

	err := c.APIRequester.Call(
		ctx,
		"PATCH",
		fmt.Sprintf("%s/customers/%s", c.Opt.XenditURL, queryString),
		c.Opt.SecretKey,
		header,
		*data,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
