package payment

import (
	"context"

	"github.com/ianeinser/xendit-go"
)

/*
// CreatePaymentMethod creates new payment method for user
func CreatePaymentMethod(data *CreatePaymentMethodParams) (*xendit.PaymentMethod2, *xendit.Error) {
	return CreatePaymentMethodWithContext(context.Background(), data)
}

// CreatePaymentMethodWithContext creates new payment
func CreatePaymentMethodWithContext(ctx context.Context, data *CreatePaymentMethodParams) (*xendit.PaymentMethod2, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreatePaymentMethodWithContext(ctx, data)
}
*/

// CreatePaymentMethod creates new payment method for user
func CreatePaymentMethod(data *map[string]interface{}) (*xendit.PaymentMethod2, *xendit.Error) {
	return CreatePaymentMethodWithContext(context.Background(), data)
}

// CreatePaymentMethodWithContext creates new payment
func CreatePaymentMethodWithContext(ctx context.Context, data *map[string]interface{}) (*xendit.PaymentMethod2, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreatePaymentMethodWithContext(ctx, data)
}

// ListPaymentMethods returns a list of matching Payment Method objects based on the provided query
func ListPaymentMethods(data *ListPaymentMethodsParams) (*ListPaymentMethodsResponse, *xendit.Error) {
	return ListPaymentMethodsWithContext(context.Background(), data)
}

// ListPaymentMethodsWithContext returns a list of matching Payment Method objects based on the provided query
func ListPaymentMethodsWithContext(ctx context.Context, data *ListPaymentMethodsParams) (*ListPaymentMethodsResponse, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.ListPaymentMethodsWithContext(ctx, data)
}

// GetPaymentMethodById returns the corresponding Payment Method that matches the provided ID.
func GetPaymentMethodById(data *GetPaymentMethodByIdParams) (*xendit.PaymentMethod2, *xendit.Error) {
	return GetPaymentMethodByIdWithContext(context.Background(), data)
}

// GetPaymentMethodByIdWithContext returns the corresponding Payment Method that matches the provided ID.
func GetPaymentMethodByIdWithContext(ctx context.Context, data *GetPaymentMethodByIdParams) (*xendit.PaymentMethod2, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetPaymentMethodByIdWithContext(ctx, data)
}

// ListPaymentsByPayentMethodId returns a list of matching Payment objects made on a Payment Method
func ListPaymentsByPaymentMethodId(data *ListPaymentsByPaymentMethodIdParams) (*ListPaymentsByPaymentMethodIdResponse, *xendit.Error) {
	return ListPaymentsByPaymentMethodIdWithContext(context.Background(), data)
}

// ListPaymentsByPayentMethodIdWithContext returns a list of matching Payment objects made on a Payment Method
func ListPaymentsByPaymentMethodIdWithContext(ctx context.Context, data *ListPaymentsByPaymentMethodIdParams) (*ListPaymentsByPaymentMethodIdResponse, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.ListPaymentsByPaymentMethodIdWithContext(ctx, data)
}

/*
// CreatePaymentRequest provides the following functionalities:
// 1. Initializes the payment/capture process for Cards, E-wallets and Direct Debits
// 2. Creates a payment method object along with the payment initialization
func CreatePaymentRequest(data *CreatePaymentRequestParams) (*xendit.PaymentRequest, *xendit.Error) {
	return CreatePaymentRequestWithContext(context.Background(), data)
}

// CreatePaymentRequestWithContext provides the following functionalities:
// 1. Initializes the payment/capture process for Cards, E-wallets and Direct Debits
// 2. Creates a payment method object along with the payment initialization
func CreatePaymentRequestWithContext(ctx context.Context, data *CreatePaymentRequestParams) (*xendit.PaymentRequest, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreatePaymentRequestWithContext(ctx, data)
}
*/

// CreatePaymentRequest provides the following functionalities:
// 1. Initializes the payment/capture process for Cards, E-wallets and Direct Debits
// 2. Creates a payment method object along with the payment initialization
func CreatePaymentRequest(data *map[string]interface{}) (*xendit.PaymentRequest, *xendit.Error) {
	return CreatePaymentRequestWithContext(context.Background(), data)
}

// CreatePaymentRequestWithContext provides the following functionalities:
// 1. Initializes the payment/capture process for Cards, E-wallets and Direct Debits
// 2. Creates a payment method object along with the payment initialization
func CreatePaymentRequestWithContext(ctx context.Context, data *map[string]interface{}) (*xendit.PaymentRequest, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreatePaymentRequestWithContext(ctx, data)
}

// ListPaymentRequests returns a list of matching Payment Request objects based on the provided query.
func ListPaymentRequests(data *ListPaymentRequestsParams) (*ListPaymentRequestsResponse, *xendit.Error) {
	return ListPaymentRequestsWithContext(context.Background(), data)
}

// ListPaymentRequestsWithContext returns a list of matching Payment Request objects based on the provided query.
func ListPaymentRequestsWithContext(ctx context.Context, data *ListPaymentRequestsParams) (*ListPaymentRequestsResponse, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.ListPaymentRequestsWithContext(ctx, data)
}

// GetPaymentRequestById returns the corresponding Payment Method that matches the provided ID
func GetPaymentRequestById(data *GetPaymentRequestByIdParams) (*xendit.PaymentRequest, *xendit.Error) {
	return GetPaymentRequestByIdWithContext(context.Background(), data)
}

// GetPaymentRequestByIdWithContext returns the corresponding Payment Method that matches the provided ID
func GetPaymentRequestByIdWithContext(ctx context.Context, data *GetPaymentRequestByIdParams) (*xendit.PaymentRequest, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetPaymentRequestByIdWithContext(ctx, data)
}

/*
// CreateRefund initialized the refund process for the provided amount for a given successful payment
func CreateRefund(data *CreateRefundParams) (*xendit.Refund, *xendit.Error) {
	return CreateRefundWithContext(context.Background(), data)
}

// CreateRefundWithContext initialized the refund process for the provided amount for a given successful payment
func CreateRefundWithContext(ctx context.Context, data *CreateRefundParams) (*xendit.Refund, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateRefundWithContext(ctx, data)
}
*/

// CreateRefund initialized the refund process for the provided amount for a given successful payment
func CreateRefund(data *map[string]interface{}) (*xendit.Refund, *xendit.Error) {
	return CreateRefundWithContext(context.Background(), data)
}

// CreateRefundWithContext initialized the refund process for the provided amount for a given successful payment
func CreateRefundWithContext(ctx context.Context, data *map[string]interface{}) (*xendit.Refund, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateRefundWithContext(ctx, data)
}

// ListPaymentMethods returns a list of matching Refund objects based on the provided query
func ListRefunds(data *ListRefundsParams) (*ListRefundsResponse, *xendit.Error) {
	return ListRefundsWithContext(context.Background(), data)
}

// ListPaymentMethodsWithContext returns a list of matching Refund objects based on the provided query
func ListRefundsWithContext(ctx context.Context, data *ListRefundsParams) (*ListRefundsResponse, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.ListRefundsWithContext(ctx, data)
}

// GetRefundbyId returns the corresponding Refund that matches the provided ID
func GetRefundById(data *GetRefundByIdParams) (*xendit.Refund, *xendit.Error) {
	return GetRefundByIdWithContext(context.Background(), data)
}

// GetRefundbyIdWithContext returns the corresponding Refund that matches the provided ID
func GetRefundByIdWithContext(ctx context.Context, data *GetRefundByIdParams) (*xendit.Refund, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetRefundByIdWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
