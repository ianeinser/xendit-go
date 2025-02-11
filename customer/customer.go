package customer

import (
	"context"

	"github.com/ianeinser/xendit-go"
)

/*
// CreateCustomer creates new customer
func CreateCustomer(data *CreateCustomerParams) (*xendit.Customer, *xendit.Error) {
	return CreateCustomerWithContext(context.Background(), data)
}

// CreateCustomerWithContext creates new payment
func CreateCustomerWithContext(ctx context.Context, data *CreateCustomerParams) (*xendit.Customer, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateCustomerWithContext(ctx, data)
}
*/

// CreateCustomer creates new customer
func CreateCustomer(data *map[string]interface{}) (*xendit.Customer, *xendit.Error) {
	return CreateCustomerWithContext(context.Background(), data)
}

// CreateCustomerWithContext creates new payment
func CreateCustomerWithContext(ctx context.Context, data *map[string]interface{}) (*xendit.Customer, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.CreateCustomerWithContext(ctx, data)
}

// GetCustomerByCustomerID gets customer by customer ID
func GetCustomerByCustomerID(data *GetCustomerByCustomerIDParams) (*[]xendit.Customer, *xendit.Error) {
	return GetCustomerByCustomerIDWithContext(context.Background(), data)
}

// GetCustomerByCustomerIDWithContext gets customer by customer ID
func GetCustomerByCustomerIDWithContext(ctx context.Context, data *GetCustomerByCustomerIDParams) (*[]xendit.Customer, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetCustomerByCustomerIDWithContext(ctx, data)
}

// GetCustomerByReferenceID gets customer by reference ID
func GetCustomerByReferenceID(data *GetCustomerByReferenceIDParams) (*[]xendit.Customer, *xendit.Error) {
	return GetCustomerByReferenceIDWithContext(context.Background(), data)
}

// GetCustomerByReferenceIDWithContext gets customer by reference ID
func GetCustomerByReferenceIDWithContext(ctx context.Context, data *GetCustomerByReferenceIDParams) (*[]xendit.Customer, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.GetCustomerByReferenceIDWithContext(ctx, data)
}

/*
// UpdateCustomer updates the details on a customer
func UpdateCustomer(data *UpdateCustomerParams) (*xendit.Customer, *xendit.Error) {
	return UpdateCustomerWithContext(context.Background(), data)
}

// UpdateCustomerWithContext updates the details on a customer
func UpdateCustomerWithContext(ctx context.Context, data *UpdateCustomerParams) (*xendit.Customer, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.UpdateCustomerWithContext(ctx, data)
}
*/

// UpdateCustomer updates the details on a customer
func UpdateCustomer(data *map[string]interface{}) (*xendit.Customer, *xendit.Error) {
	return UpdateCustomerWithContext(context.Background(), data)
}

// UpdateCustomerWithContext updates the details on a customer
func UpdateCustomerWithContext(ctx context.Context, data *map[string]interface{}) (*xendit.Customer, *xendit.Error) {
	client, err := getClient()
	if err != nil {
		return nil, err
	}

	return client.UpdateCustomerWithContext(ctx, data)
}

func getClient() (*Client, *xendit.Error) {
	return &Client{
		Opt:          &xendit.Opt,
		APIRequester: xendit.GetAPIRequester(),
	}, nil
}
