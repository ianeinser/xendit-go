package payment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/ianeinser/xendit-go"
	"github.com/ianeinser/xendit-go/utils/validator"
)

type Client struct {
	Opt          *xendit.Option
	APIRequester xendit.APIRequester
}

// CreatePaymentMethod creates new payment method for user
func (c *Client) CreatePaymentMethod(data *CreatePaymentMethodParams) (*xendit.PaymentMethod2, *xendit.Error) {
	return c.CreatePaymentMethodWithContext(context.Background(), data)
}

// CreatePaymentMethodWithContext creates new payment method for user
func (c *Client) CreatePaymentMethodWithContext(ctx context.Context, data *CreatePaymentMethodParams) (*xendit.PaymentMethod2, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.PaymentMethod2{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}
	if data.APIVersion != "" {
		header.Add("api-version", data.APIVersion)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/v2/payment_methods", c.Opt.XenditURL),
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

// ListPaymentMethods returns a list of matching Payment Method objects based on the provided query
func (c *Client) ListPaymentMethods(data *ListPaymentMethodsParams) (*ListPaymentMethodsResponse, *xendit.Error) {
	return c.ListPaymentMethodsWithContext(context.Background(), data)
}

// ListPaymentMethodsWithContext returns a list of matching Payment Method objects based on the provided query
func (c *Client) ListPaymentMethodsWithContext(ctx context.Context, data *ListPaymentMethodsParams) (*ListPaymentMethodsResponse, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &ListPaymentMethodsResponse{}
	var queryString string

	if data != nil {
		queryString = data.QueryString()
	}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}
	if data.APIVersion != "" {
		header.Add("api-version", data.APIVersion)
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/v2/payment_methods?%s", c.Opt.XenditURL, queryString),
		c.Opt.SecretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// ListPaymentsByPayentMethodId returns a list of matching Payment objects made on a Payment Method
func (c *Client) ListPaymentsByPaymentMethodId(data *ListPaymentsByPaymentMethodIdParams) (*ListPaymentsByPaymentMethodIdResponse, *xendit.Error) {
	return c.ListPaymentsByPaymentMethodIdWithContext(context.Background(), data)
}

// ListPaymentMethodsWithContext returns a list of matching Payment objects made on a Payment Method
func (c *Client) ListPaymentsByPaymentMethodIdWithContext(ctx context.Context, data *ListPaymentsByPaymentMethodIdParams) (*ListPaymentsByPaymentMethodIdResponse, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &ListPaymentsByPaymentMethodIdResponse{}
	var queryString string

	id := ""
	if data != nil {
		id = data.ID
		data.ID = ""
		queryString = data.QueryString()
	}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	if data.APIVersion != "" {
		header.Add("api-version", data.APIVersion)
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/v2/payment_methods/%s/payments?%s", c.Opt.XenditURL, id, queryString),
		c.Opt.SecretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreatePaymentRequest provides the following functionalities:
// 1. Initializes the payment/capture process for Cards, E-wallets and Direct Debits
// 2. Creates a payment method object along with the payment initialization
func (c *Client) CreatePaymentRequest(data *CreatePaymentRequestParams) (*xendit.PaymentRequest, *xendit.Error) {
	return c.CreatePaymentRequestWithContext(context.Background(), data)
}

// CreatePaymentMethodWithContext provides the following functionalities:
// 1. Initializes the payment/capture process for Cards, E-wallets and Direct Debits
// 2. Creates a payment method object along with the payment initialization
func (c *Client) CreatePaymentRequestWithContext(ctx context.Context, data *CreatePaymentRequestParams) (*xendit.PaymentRequest, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.PaymentRequest{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}
	if data.APIVersion != "" {
		header.Add("api-version", data.APIVersion)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/payment_requests", c.Opt.XenditURL),
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

// ListPaymentRequests returns a list of matching Payment Request objects based on the provided query.
func (c *Client) ListPaymentRequests(data *ListPaymentRequestsParams) (*ListPaymentRequestsResponse, *xendit.Error) {
	return c.ListPaymentRequestsWithContext(context.Background(), data)
}

// ListPaymentRequestsWithContext returns a list of matching Payment Request objects based on the provided query.
func (c *Client) ListPaymentRequestsWithContext(ctx context.Context, data *ListPaymentRequestsParams) (*ListPaymentRequestsResponse, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &ListPaymentRequestsResponse{}
	var queryString string

	if data != nil {
		queryString = data.QueryString()
	}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}
	if data.APIVersion != "" {
		header.Add("api-version", data.APIVersion)
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/payment_requests?%s", c.Opt.XenditURL, queryString),
		c.Opt.SecretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetPaymentRequestById returns the corresponding Payment Method that matches the provided ID
func (c *Client) GetPaymentRequestById(data *GetPaymentRequestsByIdParams) (*xendit.PaymentRequest, *xendit.Error) {
	return c.GetPaymentRequestByIdWithContext(context.Background(), data)
}

// GetPaymentRequestByIdWithContext returns the corresponding Payment Method that matches the provided ID.
func (c *Client) GetPaymentRequestByIdWithContext(ctx context.Context, data *GetPaymentRequestsByIdParams) (*xendit.PaymentRequest, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.PaymentRequest{}

	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	if data.APIVersion != "" {
		header.Add("api-version", data.APIVersion)
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/payment_requests/%s", c.Opt.XenditURL, data.ID),
		c.Opt.SecretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// CreateRefund initialized the refund process for the provided amount for a given successful payment
func (c *Client) CreateRefund(data *CreateRefundParams) (*xendit.Refund, *xendit.Error) {
	return c.CreateRefundWithContext(context.Background(), data)
}

// CreateRefundWithContext initialized the refund process for the provided amount for a given successful payment
func (c *Client) CreateRefundWithContext(ctx context.Context, data *CreateRefundParams) (*xendit.Refund, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.Refund{}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}
	if data.APIVersion != "" {
		header.Add("api-version", data.APIVersion)
	}

	err := c.APIRequester.Call(
		ctx,
		"POST",
		fmt.Sprintf("%s/refunds", c.Opt.XenditURL),
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

// ListPaymentMethods returns a list of matching Refund objects based on the provided query
func (c *Client) ListRefunds(data *ListRefundsParams) (*ListRefundsResponse, *xendit.Error) {
	return c.ListRefundsWithContext(context.Background(), data)
}

// ListPaymentMethodsWithContext returns a list of matching Refund objects based on the provided query
func (c *Client) ListRefundsWithContext(ctx context.Context, data *ListRefundsParams) (*ListRefundsResponse, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &ListRefundsResponse{}
	var queryString string

	if data != nil {
		queryString = data.QueryString()
	}
	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}
	if data.APIVersion != "" {
		header.Add("api-version", data.APIVersion)
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/refunds?%s", c.Opt.XenditURL, queryString),
		c.Opt.SecretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// GetRefundbyId returns the corresponding Refund that matches the provided ID
func (c *Client) GetRefundById(data *GetRefundByIdParams) (*xendit.Refund, *xendit.Error) {
	return c.GetRefundByIdWithContext(context.Background(), data)
}

// GetRefundbyIdWithContext returns the corresponding Refund that matches the provided ID
func (c *Client) GetRefundByIdWithContext(ctx context.Context, data *GetRefundByIdParams) (*xendit.Refund, *xendit.Error) {
	if err := validator.ValidateRequired(ctx, data); err != nil {
		return nil, validator.APIValidatorErr(err)
	}

	response := &xendit.Refund{}

	header := http.Header{}

	if data.ForUserID != "" {
		header.Add("for-user-id", data.ForUserID)
	}

	if data.APIVersion != "" {
		header.Add("api-version", data.APIVersion)
	}

	err := c.APIRequester.Call(
		ctx,
		"GET",
		fmt.Sprintf("%s/refunds/%s", c.Opt.XenditURL, data.ID),
		c.Opt.SecretKey,
		nil,
		nil,
		&response,
	)
	if err != nil {
		return nil, err
	}

	return response, nil
}
