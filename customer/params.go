package customer

import (
	"net/url"

	"github.com/ianeinser/xendit-go"
)

// CreateCustomerParams contains parameters for CreateCustomer
type CreateCustomerParams struct {
	IdempotencyKey         string                   `json:"-"`
	ForUserID              string                   `json:"-"`
	APIVersion             string                   `json:"-"`
	ReferenceID            string                   `json:"reference_id" validate:"required"`
	Type                   string                   `json:"type"`
	IndividualDetail       xendit.IndividualDetail  `json:"individual_detail"`
	BusinessDetail         xendit.BusinessDetail    `json:"business_detail"`
	MobileNumber           string                   `json:"mobile_number,omitempty"`
	PhoneNumber            string                   `json:"phone_number,omitempty"`
	HashedPhoneNumber      string                   `json:"hashed_phone_number"`
	Email                  string                   `json:"email,omitempty"`
	Addresses              []xendit.CustomerAddress `json:"addresses,omitempty"`
	Description            string                   `json:"description,omitempty"`
	DateOfRegistration     string                   `json:"date_of_registration"`
	DomicileofRegistration string                   `json:"domicile_of_registration"`
	Metadata               map[string]interface{}   `json:"metadata,omitempty"`
}

// GetCustomerByReferenceIDParams contains parameters for GetCustomerByReferenceID
type GetCustomerByCustomerIDParams struct {
	CustomerID string `json:"-"`
}

// QueryString creates query string from GetCustomerByReferenceIDParams, ignores nil values
func (p *GetCustomerByCustomerIDParams) QueryString() string {
	urlValues := &url.Values{}

	urlValues.Add("", p.CustomerID)

	return urlValues.Encode()
}

// GetCustomerByReferenceIDParams contains parameters for GetCustomerByReferenceID
type GetCustomerByReferenceIDParams struct {
	ReferenceID string `json:"-"`
}

// QueryString creates query string from GetCustomerByReferenceIDParams, ignores nil values
func (p *GetCustomerByReferenceIDParams) QueryString() string {
	urlValues := &url.Values{}

	urlValues.Add("reference_id", p.ReferenceID)

	return urlValues.Encode()
}

// UpdateCustomerParams contains parameters for UpdateCustomer
type UpdateCustomerParams struct {
	ForUserID        string                   `json:"-"`
	APIVersion       string                   `json:"-"`
	IndividualDetail xendit.IndividualDetail  `json:"individual_detail"`
	BusinessDetail   xendit.BusinessDetail    `json:"business_detail"`
	MobileNumber     string                   `json:"mobile_number"`
	Email            string                   `json:"email,omitempty"`
	Addresses        []xendit.CustomerAddress `json:"addresses,omitempty"`
	Metadata         map[string]interface{}   `json:"metadata,omitempty"`
}
