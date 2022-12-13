package xendit

import "gorm.io/datatypes"

type Payment struct {
	ID                string            `json:"id"`
	PaymentRequestID  string            `json:"payment_request_id,omitempty"`
	ReferenceID       string            `json:"reference_id"`
	CustomerID        string            `json:"customer_id,omitempty"`
	Currency          string            `json:"currency"`
	Amount            float64           `json:"amount"`
	Country           string            `json:"country"`
	Status            string            `json:"status"`
	PaymentMethod     PaymentMethod2    `json:"payment_method" gorm:"embedded;embedded_prefix:pm_"`
	ChannelProperties ChannelProperties `json:"channel_properties" gorm:"embedded;embedded_prefix:cp_"`
	PaymentDetail     PaymentDetail     `json:"payment_detail" gorm:"embedded;embedded_prefix:pd_"`
	FailureCode       string            `json:"failure_code,omitempty"`
	Created           string            `json:"created"`
	Updated           string            `json:"updated"`
	Metadata          datatypes.JSONMap `json:"metadata,omitempty"`
}

type ChannelProperties struct {
	RedeemPoints         string `json:"redeem_properties,omitempty"`
	SuccessReturnUrl     string `json:"success_return_url,omitempty"`
	FailureReturnUrl     string `json:"failure_return_url,omitempty"`
	CancelReturnUrl      string `json:"cancel_return_url,omitempty"`
	RequiredAuth         string `json:"required_auth,omitempty"`
	MobileNumber         string `json:"mobile_number,omitempty"`
	Cashtag              string `json:"cashtag,omitempty"`
	CardLastFour         string `json:"card_last_four,omitempty"`
	CardExpiry           string `json:"card_expiry,omitempty"`
	Email                string `json:"email,omitempty"`
	SkipThreeDSecure     bool   `json:"skip_three_d_secure,omitempty"`
	CardOnFileType       string `json:"cardonfile_type,omitempty"`
	CustomerName         string `json:"customer_name,omitempty"`
	PaymentCode          string `json:"payment_code,omitempty"`
	ExpiresAt            string `json:"expires_at,omitempty"`
	VirtualAccountNumber string `json:"virtual_account_number,omitempty"`
	SuggestedAmount      string `json:"suggested_amount,omitempty"`
	QrString             string `json:"qr_string"`
}

type PaymentDetail struct {
	Remarks string `json:"remarks"`
}

type PaymentMethod2 struct {
	ID                 string             `json:"id"`
	BusinessID         string             `json:"business_id"`
	CustomerID         string             `json:"customer_id,omitempty"`
	ReferenceID        string             `json:"reference_id,omitempty"`
	Reusability        string             `json:"reusability"`
	Country            string             `json:"country"`
	Status             string             `json:"status"`
	Actions            Actions2           `json:"actions,omitempty" gorm:"embedded;embedded_prefix:act_"`
	PaymentRequestID   string             `json:"payment_request_id"`
	Currency           string             `json:"currency"`
	Amount             float64            `json:"amount"`
	Type               string             `json:"type"`
	Ewallet            Ewallet            `json:"ewallet,omitempty" gorm:"embedded;embedded_prefix:ew_"`
	DirectDebit        DirectDebit        `json:"direct_debit,omitempty" gorm:"embedded;embedded_prefix:dd_"`
	Card               Card               `json:"card,omitempty" gorm:"embedded;embedded_prefix:cd_"`
	OverTheCounter     OverTheCounter     `json:"over_the_counter,omitempty" gorm:"embedded;embedded_prefix:otc_"`
	VirtualAccount     VirtualAccount2    `json:"virtual_account,omitempty" gorm:"embedded;embedded_prefix:va_"`
	QrCode             QrCode             `json:"qr_code,omitempty" gorm:"embedded;embedded_prefix:qr_"`
	Description        string             `json:"description,omitempty"`
	BillingInformation BillingInformation `json:"billing_information,omitempty" gorm:"embedded;embedded_prefix:bi_"`
	Created            string             `json:"created"`
	Updated            string             `json:"updated"`
	Metadata           datatypes.JSONMap  `json:"metadata,omitempty"`
}

type Actions2 struct {
	Method  string `json:"method,omitempty"`
	UrlType string `json:"url_type,omitempty"`
	Action  string `json:"action,omitempty"`
	Url     string `json:"url,omitempty"`
}

type Ewallet struct {
	ChannelCode       string            `json:"channel_code,omitempty"`
	ChannelProperties ChannelProperties `json:"channel_properties,omitempty" gorm:"embedded;embedded_prefix:cp_"`
	Account           Account2          `json:"account,omitempty" gorm:"embedded;embedded_prefix:acc_"`
}

type Account2 struct {
	AccountDetails string  `json:"account_details,omitempty"`
	Name           string  `json:"name,omitempty"`
	Balance        float64 `json:"balance,omitempty"`
	PointBalance   float64 `json:"point_balance,omitempty"`
}

type DirectDebit struct {
	ChannelCode       string            `json:"channel_code,omitempty"`
	Type              string            `json:"type,omitempty"`
	ChannelProperties ChannelProperties `json:"channel_properties,omitempty" gorm:"embedded;embedded_prefix:cp_"`
	BankAccount       BankAccount       `json:"bank_account,omitempty" gorm:"embedded;embedded_prefix:ba_"`
	DebitCard         DebitCard         `json:"debit_card,omitempty" gorm:"embedded;embedded_prefix:dc_"`
}

type BankAccount struct {
	MaskedBankAccountNumber string `json:"masked_bank_account_number,omitempty"`
	BankAccountHash         string `json:"bank_account_hash,omitempty"`
}

type DebitCard struct {
	MobileNumber string `json:"mobile_number,omitempty"`
	CardLastFour string `json:"card_last_four,omitempty"`
	CardExpiry   string `json:"card_expiry,omitempty"`
	Email        string `json:"email,omitempty"`
}

type Card struct {
	Currency          string            `json:"currency,omitempty"`
	ChannelProperties ChannelProperties `json:"channel_properties,omitempty" gorm:"embedded;embedded_prefix:cp_"`
	CardInformation   CardInformation   `json:"card_information,omitempty" gorm:"embedded;embedded_prefix:ci_"`
}

type CardInformation struct {
	TokenID          string `json:"token_id,omitempty"`
	CardNumber       string `json:"card_number,omitempty"`
	MaskedCardNumber string `json:"masked_card_number,omitempty"`
	ExpiryMonth      string `json:"expiry_month,omitempty"`
	ExpiryYear       string `json:"expiry_year,omitempty"`
	CardholderName   string `json:"cardholder_name,omitempty"`
	Cvv              string `json:"cvv,omitempty"`
	Fingerprint      string `json:"fingerprint,omitempty"`
	Type             string `json:"type,omitempty"`
	Network          string `json:"network,omitempty"`
	Country          string `json:"country,omitempty"`
	Issuer           string `json:"issuer,omitempty"`
}

type OverTheCounter struct {
	ChannelCode       string            `json:"channel_code,omitempty"`
	Currency          string            `json:"currency,omitempty"`
	Amount            float64           `json:"amount,omitempty"`
	ChannelProperties ChannelProperties `json:"channel_properties,omitempty" gorm:"embedded;embedded_prefix:cp_"`
}

type VirtualAccount2 struct {
	ChannelCode       string            `json:"channel_code,omitempty"`
	Currency          string            `json:"currency,omitempty"`
	Amount            float64           `json:"amount,omitempty"`
	ChannelProperties ChannelProperties `json:"channel_properties,omitempty" gorm:"embedded;embedded_prefix:cp_"`
}

type QrCode struct {
	ChannelCode       string            `json:"channel_code,omitempty"`
	Currency          string            `json:"currency,omitempty"`
	Amount            float64           `json:"amount,omitempty"`
	ChannelProperties ChannelProperties `json:"channel_properties,omitempty" gorm:"embedded;embedded_prefix:cp_"`
}

type BillingInformation struct {
	Country       string `json:"country,omitempty"`
	StreetLine1   string `json:"street_line1,omitempty"`
	StreetLine2   string `json:"street_line2,omitempty"`
	City          string `json:"city,omitempty"`
	ProvinceState string `json:"province_state,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
}

type PaymentRequest struct {
	ID                      string                  `json:"id"`
	BusinessID              string                  `json:"business_id"`
	CustomerID              string                  `json:"customer_id,omitempty"`
	ReferenceID             string                  `json:"reference_id,omitempty"`
	Currency                string                  `json:"currency"`
	Amount                  float64                 `json:"amount"`
	CaptureAmount           float64                 `json:"capture_amount"`
	Country                 string                  `json:"country"`
	Status                  string                  `json:"status"`
	Description             string                  `json:"description,omitempty"`
	PaymentMethod           PaymentMethod2          `json:"payment_method" gorm:"embedded;embedded_prefix:pm_"`
	Actions                 Actions2                `json:"actions,omitempty" gorm:"embedded;embedded_prefix:act_"`
	CaptureMethod           string                  `json:"capture_method"`
	Initiator               string                  `json:"initiator"`
	ChannelProperties       ChannelProperties       `json:"channel_properties,omitempty" gorm:"embedded;embedded_prefix:cp_"`
	ShippingInformation     ShippingInformation2    `json:"shipping_information,omitempty" gorm:"embedded;embedded_prefix:si_"`
	CardVerificationResults CardVerificationResults `json:"card_verification_results,omitempty" gorm:"embedded;embedded_prefix:cvr_"`
	FailureCode             string                  `json:"failure_code"`
	Created                 string                  `json:"created"`
	Updated                 string                  `json:"updated"`
	Metadata                datatypes.JSONMap       `json:"metadata,omitempty"`
}

type ShippingInformation2 struct {
	Country       string `json:"country,omitempty"`
	StreetLine1   string `json:"street_line1,omitempty"`
	StreetLine2   string `json:"street_line2,omitempty"`
	City          string `json:"city,omitempty"`
	ProvinceState string `json:"province_state,omitempty"`
	PostalCode    string `json:"postal_code,omitempty"`
}

type CardVerificationResults struct {
	ThreeDSecure              ThreeDSecure `json:"three_d_secure,omitempty" gorm:"embedded;embedded_prefix:tds_"`
	CvvResult                 string       `json:"cvv_result,omitempty"`
	AddressVerificationResult string       `json:"addres_verification_result,omitempty"`
}

type ThreeDSecure struct {
	ThreeDSecureFlow    string `json:"three_d_Secure_flow,omitempty"`
	EciCode             string `json:"eci_code,omitempty"`
	ThreeDSecureResult  string `json:"three_d_secure_result,omitempty"`
	ThreeDSecureVersion string `json:"three_d_secure_version,omitempty"`
}

type Refund struct {
	ID                string  `json:"id"`
	PaymentRequestID  string  `json:"payment_request_id"`
	InvoiceID         string  `json:"invoice_id"`
	PaymentMethodType string  `json:"payment_method_type"`
	ReferenceID       string  `json:"reference_id"`
	Status            string  `json:"status"`
	Country           string  `json:"country"`
	ChannelCode       string  `json:"channel_code"`
	Reason            string  `json:"reason"`
	FailureCode       string  `json:"failure_code"`
	RefundFeeAmount   float64 `json:"refund_fee_amount"`
	Created           string  `json:"created"`
	Updated           string  `json:"updated"`
	Metadata          string  `json:"metadata,omitempty"`
}
