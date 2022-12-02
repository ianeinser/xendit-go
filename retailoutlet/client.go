package retailoutlet

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ianeinser/xendit-go"
	"github.com/ianeinser/xendit-go/utils/validator"
)

// Client is the client used to invoke invoice API.
type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

// CreateFixedPaymentCode creates new retail outlet fixed payment code
func (c *Client) CreateFixedPaymentCode(data *CreateFixedPaymentCodeParams) (*xendit.RetailOutlet, *xendit.Error) {
	return c.CreateFixedPaymentCodeWithContext(context.Background(), data)
}

// CreateFixedPaymentCodeWithContext creates new retail outlet fixed payment code with context
func (c *Client) CreateFixedPaymentCodeWithContext(ctx context.Context, data *CreateFixedPaymentCodeParams) (*xendit.RetailOutlet, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.RetailOutlet{}
	header := http.Header{}
	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/fixed_payment_code", c.Opt.XenditURL),
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

// GetFixedPaymentCode gets one retail outlet fixed payment code
func (c *Client) GetFixedPaymentCode(data *GetFixedPaymentCodeParams) (*xendit.RetailOutlet, *xendit.Error) {
	return c.GetFixedPaymentCodeWithContext(context.Background(), data)
}

// GetFixedPaymentCodeWithContext gets one retail outlet fixed payment code with context
func (c *Client) GetFixedPaymentCodeWithContext(ctx context.Context, data *GetFixedPaymentCodeParams) (*xendit.RetailOutlet, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.RetailOutlet{}
	header := http.Header{}
	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/fixed_payment_code/%s", c.Opt.XenditURL, data.FixedPaymentCodeID),
		c.Opt.SecretKey,
		header,
		nil,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetPaymentByFixedPaymentCode gets list retail outlet fixed payment code
func (c *Client) GetPaymentByFixedPaymentCode(data *GetPaymentByFixedPaymentCodeParams) (*xendit.RetailOutletPayments, *xendit.Error) {
	return c.GetPaymentByFixedPaymentCodeWithContext(context.Background(), data)
}

// GetPaymentByFixedPaymentCodeWithContext gets list retail outlet fixed payment code with context
func (c *Client) GetPaymentByFixedPaymentCodeWithContext(ctx context.Context, data *GetPaymentByFixedPaymentCodeParams) (*xendit.RetailOutletPayments, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.RetailOutletPayments{}
	header := http.Header{}
	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/fixed_payment_code/%s/payments", c.Opt.XenditURL, data.FixedPaymentCodeID),
		c.Opt.SecretKey,
		header,
		nil,
		response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// UpdateFixedPaymentCode updates a retail outlet fixed payment code
func (c *Client) UpdateFixedPaymentCode(data *UpdateFixedPaymentCodeParams) (*xendit.RetailOutlet, *xendit.Error) {
	return c.UpdateFixedPaymentCodeWithContext(context.Background(), data)
}

// UpdateFixedPaymentCodeWithContext updates a retail outlet fixed payment code with context
func (c *Client) UpdateFixedPaymentCodeWithContext(ctx context.Context, data *UpdateFixedPaymentCodeParams) (*xendit.RetailOutlet, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.RetailOutlet{}
	header := http.Header{}
	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	err := c.APIRequester.Call(
		ctx,
		"PATCH",
		fmt.Sprintf("%s/fixed_payment_code/%s", c.Opt.XenditURL, data.FixedPaymentCodeID),
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
