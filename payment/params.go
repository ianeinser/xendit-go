package payment

import (
	"net/url"

	"github.com/ianeinser/xendit-go"
)

// CreatePaymentMethodParams contains parameters for CreatePaymentMethod
type CreatePaymentMethodParams struct {
	IdempotencyKey     string                    `json:"-"`
	ForUserID          string                    `json:"-"`
	Type               string                    `json:"type"`
	Reusability        string                    `json:"reusability"`
	ReferenceID        string                    `json:"reference_id,omitempty"`
	CustomerID         string                    `json:"customer_id,omitempty"`
	Country            string                    `json:"country"`
	Description        string                    `json:"description,omitempty"`
	BillingInformation xendit.BillingInformation `json:"billing_information,omitempty"`
	Metadata           map[string]interface{}    `json:"metadata,omitempty"`
	Ewallet            xendit.Ewallet            `json:"ewallet,omitempty"`
	DirectDebit        xendit.DirectDebit        `json:"direct_debit,omitempty"`
	Card               xendit.Card               `json:"card,omitempty"`
	OverTheCounter     xendit.OverTheCounter     `json:"over_the_counter,omitempty"`
	VirtualAccount     xendit.VirtualAccount2    `json:"virtual_account,omitempty"`
	QrCode             xendit.QrCode             `json:"qr_code,omitempty"`
}

// ListPaymentMethodsParams contains parameters for ListPaymentMethods
type ListPaymentMethodsParams struct {
	ForUserID   string  `json:"-"`
	ID          string  `json:"id"`
	Type        string  `json:"type"`
	Reusability string  `json:"reusability"`
	ReferenceID string  `json:"reference_id,omitempty"`
	CustomerID  string  `json:"customer_id,omitempty"`
	Limit       float64 `json:"limit"`
	AfterId     string  `json:"after_id"`
	BeforeId    string  `json:"before_id"`
}

// QueryString creates query string from ListPaymentMethodsParams, ignores nil values
func (p *ListPaymentMethodsParams) QueryString() string {
	urlValues := &url.Values{}

	urlValues.Add("reference_id", p.ReferenceID)

	return urlValues.Encode()
}

type ListPaymentMethodsResponse struct {
	Data    []xendit.PaymentMethod2 `json:"data"`
	HasMore bool                    `json:"has_more"`
}

// GetPaymentMethodByIdParams contains parameters for GetPaymentMethodById
type GetPaymentMethodByIdParams struct {
	ForUserID string `json:"-"`
	ID        string `json:"-"`
}

// ListPaymentMethodsParams contains parameters for ListPaymentMethods
type ListPaymentsByPaymentMethodIdParams struct {
	ForUserID        string  `json:"-"`
	ID               string  `json:"-"`
	PaymentRequestID string  `json:"payment_request_id"`
	ReferenceID      string  `json:"reference_id,omitempty"`
	Status           string  `json:"status"`
	Limit            float64 `json:"limit"`
	AfterID          string  `json:"after_id"`
	BeforeID         string  `json:"before_id"`
}

// QueryString creates query string from ListPaymentMethodsParams, ignores nil values
func (p *ListPaymentsByPaymentMethodIdParams) QueryString() string {
	urlValues := &url.Values{}

	urlValues.Add("reference_id", p.ReferenceID)

	return urlValues.Encode()
}

type Link struct {
	Href   string `json:"href"`
	Rel    string `json:"rel"`
	Method string `json:"method"`
}

type ListPaymentsByPaymentMethodIdResponse struct {
	Data    []xendit.Payment `json:"data"`
	HasMore bool             `json:"has_more"`
	Links   []Link           `json:"links"`
}

// CreatePaymentRequestParams contains parameters for CreatePaymentRequest
type CreatePaymentRequestParams struct {
	IdempotencyKey    string                   `json:"-"`
	ForUserID         string                   `json:"-"`
	APIVersion        string                   `json:"-"`
	Currency          string                   `json:"currency"`
	Amount            float64                  `json:"amount"`
	ReferenceID       string                   `json:"reference_id,omitempty"`
	CustomerID        string                   `json:"customer_id,omitempty"`
	Country           string                   `json:"country"`
	Description       string                   `json:"description,omitempty"`
	PaymentMethod     xendit.PaymentMethod2    `json:"payment_method"`
	PaymentMethodID   string                   `json:"payment_method_id"`
	ChannelProperties xendit.ChannelProperties `json:"channel_properties"`
	Metadata          map[string]interface{}   `json:"metadata,omitempty"`
}

// ListPaymentRequestsParams contains parameters for ListPaymentRequests
type ListPaymentRequestsParams struct {
	ForUserID   string  `json:"-"`
	ReferenceID string  `json:"reference_id,omitempty"`
	CustomerID  string  `json:"customer_id,omitempty"`
	Type        string  `json:"type"`
	ChannelCode string  `json:"channel_code"`
	Status      string  `json:"status"`
	Limit       float64 `json:"limit"`
	AfterId     string  `json:"after_id"`
	BeforeId    string  `json:"before_id"`
}

// QueryString creates query string from ListPaymentRequestsParams, ignores nil values
func (p *ListPaymentRequestsParams) QueryString() string {
	urlValues := &url.Values{}

	urlValues.Add("reference_id", p.ReferenceID)

	return urlValues.Encode()
}

type ListPaymentRequestsResponse struct {
	Data    []xendit.PaymentRequest `json:"data"`
	HasMore bool                    `json:"has_more"`
}

// ListPaymentRequestsParams contains parameters for CreatePaymentRequest
type GetPaymentRequestsByIdParams struct {
	ForUserID  string `json:"-"`
	APIVersion string `json:"-"`
	ID         string `json:"id"`
}

// CreateRefundParams contains parameters for CreateRefund
type CreateRefundParams struct {
	IdempotencyKey   string                 `json:"-"`
	ForUserID        string                 `json:"-"`
	PaymentRequestID string                 `json:"payment_request_id"`
	ReferenceID      string                 `json:"reference_id,omitempty"`
	InvoiceID        string                 `json:"invoice_id"`
	Currency         string                 `json:"currency"`
	Amount           float64                `json:"amount"`
	Reason           string                 `json:"reason"`
	Metadata         map[string]interface{} `json:"metadata,omitempty"`
}

// ListPaymentRequestsParams contains parameters for ListPaymentRequests
type ListRefundsParams struct {
	ForUserID         string  `json:"-"`
	PaymentRequestID  string  `json:"payment_request_id"`
	InvoiceID         string  `json:"invoice_id"`
	PaymentMethodType string  `json:"payment_method_type"`
	ChannelCode       string  `json:"channel_code"`
	Limit             float64 `json:"limit"`
	AfterID           string  `json:"after_id"`
	BeforeID          string  `json:"before_id"`
}

// QueryString creates query string from ListPaymentRequestsParams, ignores nil values
func (p *ListRefundsParams) QueryString() string {
	urlValues := &url.Values{}

	urlValues.Add("payment_method_type", p.PaymentMethodType)

	return urlValues.Encode()
}

type ListRefundsResponse struct {
	Data    []xendit.Refund `json:"data"`
	HasMore bool            `json:"has_more"`
	Links   []Link          `json:"links"`
}

// ListPaymentRequestsParams contains parameters for CreatePaymentRequest
type GetRefundByIdParams struct {
	ForUserID string `json:"-"`
	ID        string `json:"-"`
}
