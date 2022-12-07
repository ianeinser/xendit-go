package xendit

type Payment struct {
	ID                string                 `json:"id"`
	PaymentRequestID  string                 `json:"payment_request_id,omitempty"`
	ReferenceID       string                 `json:"reference_id"`
	CustomerID        string                 `json:"customer_id,omitempty"`
	Currency          string                 `json:"currency"`
	Amount            float64                `json:"amount"`
	Country           string                 `json:"country"`
	Status            string                 `json:"status"`
	PaymentMethod     PaymentMethod2         `json:"payment_method" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	ChannelProperties ChannelProperties      `json:"channel_properties,omitempty" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	PaymentDetail     PaymentDetail          `json:"payment_detail,omitempty" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	FailureCode       string                 `json:"failure_code,omitempty"`
	Created           string                 `json:"created"`
	Updated           string                 `json:"updated"`
	Metadata          map[string]interface{} `json:"metadata,omitempty"`
}

type ChannelProperties struct {
	ReferenceID          string `json:"reference_id"`
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
	ReferenceID string `json:"reference_id"`
	Remarks     string `json:"remarks"`
}

type PaymentMethod2 struct {
	ID                 string                 `json:"id"`
	BusinessID         string                 `json:"business_id"`
	CustomerID         string                 `json:"customer_id,omitempty"`
	ReferenceID        string                 `json:"reference_id,omitempty"`
	Reusability        string                 `json:"reusability"`
	Country            string                 `json:"country"`
	Status             string                 `json:"status"`
	Actions            Actions2               `json:"actions" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	PaymentRequestID   string                 `json:"payment_request_id"`
	Currency           string                 `json:"currency"`
	Amount             float64                `json:"amount"`
	Type               string                 `json:"type"`
	Ewallet            Ewallet                `json:"ewallet,omitempty" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	DirectDebit        DirectDebit            `json:"direct_debit,omitempty" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	Card               Card                   `json:"card,omitempty" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	OverTheCounter     OverTheCounter         `json:"over_the_counter,omitempty" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	VirtualAccount     VirtualAccount2        `json:"virtual_account,omitempty" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	QrCode             QrCode                 `json:"qr_code,omitempty" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	Description        string                 `json:"description,omitempty"`
	BillingInformation BillingInformation     `json:"billing_information,omitempty" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	Created            string                 `json:"created"`
	Updated            string                 `json:"updated"`
	Metadata           map[string]interface{} `json:"metadata,omitempty"`
}

type Actions2 struct {
	ReferenceID string `json:"reference_id,omitempty"`
	Method      string `json:"method"`
	UrlType     string `json:"url_type"`
	Action      string `json:"action"`
	Url         string `json:"url"`
}

type Ewallet struct {
	ReferenceID       string            `json:"reference_id,omitempty"`
	ChannelCode       string            `json:"channel_code"`
	ChannelProperties ChannelProperties `json:"channel_properties" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	Account           Account2          `json:"account,omitempty" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
}

type Account2 struct {
	ReferenceID    string  `json:"reference_id,omitempty"`
	AccountDetails string  `json:"account_details,omitempty"`
	Name           string  `json:"name,omitempty"`
	Balance        float64 `json:"balance,omitempty"`
	PointBalance   float64 `json:"point_balance,omitempty"`
}

type DirectDebit struct {
	ReferenceID       string            `json:"reference_id,omitempty"`
	ChannelCode       string            `json:"channel_code"`
	Type              string            `json:"type"`
	ChannelProperties ChannelProperties `json:"channel_properties" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	BankAccount       BankAccount       `json:"bank_account,omitempty" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	DebitCard         DebitCard         `json:"debit_account,omitempty" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
}

type BankAccount struct {
	ReferenceID             string `json:"reference_id,omitempty"`
	MaskedBankAccountNumber string `json:"masked_bank_account_number,omitempty"`
	BankAccountHash         string `json:"bank_account_hash,omitempty"`
}

type DebitCard struct {
	ReferenceID  string `json:"reference_id,omitempty"`
	MobileNumber string `json:"mobile_number,omitempty"`
	CardLastFour string `json:"card_last_four,omitempty"`
	CardExpiry   string `json:"card_expiry,omitempty"`
	Email        string `json:"email,omitempty"`
}

type Card struct {
	ReferenceID       string            `json:"reference_id,omitempty"`
	Currency          string            `json:"currency"`
	ChannelProperties ChannelProperties `json:"channel_properties" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	CardInformation   CardInformation   `json:"card_information" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
}

type CardInformation struct {
	ReferenceID      string `json:"reference_id,omitempty"`
	TokenID          string `json:"token_id"`
	CardNumber       string `json:"card_number,omitempty"`
	MaskedCardNumber string `json:"masked_card_number,omitempty"`
	ExpiryMonth      string `json:"expiry_month"`
	ExpiryYear       string `json:"expiry_year"`
	CardholderName   string `json:"cardholder_name,omitempty"`
	Cvv              string `json:"cvv,omitempty"`
	Fingerprint      string `json:"fingerprint,omitempty"`
	Type             string `json:"type,omitempty"`
	Network          string `json:"network"`
	Country          string `json:"country,omitempty"`
	Issuer           string `json:"issuer,omitempty"`
}

type OverTheCounter struct {
	ReferenceID       string            `json:"reference_id,omitempty"`
	ChannelCode       string            `json:"channel_code"`
	Currency          string            `json:"currency"`
	Amount            float64           `json:"amount,omitempty"`
	ChannelProperties ChannelProperties `json:"channel_properties" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
}

type VirtualAccount2 struct {
	ReferenceID       string            `json:"reference_id,omitempty"`
	ChannelCode       string            `json:"channel_code"`
	Currency          string            `json:"currency"`
	Amount            float64           `json:"amount,omitempty"`
	ChannelProperties ChannelProperties `json:"channel_properties" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
}

type QrCode struct {
	ReferenceID       string            `json:"reference_id,omitempty"`
	ChannelCode       string            `json:"channel_code"`
	Currency          string            `json:"currency"`
	Amount            float64           `json:"amount,omitempty"`
	ChannelProperties ChannelProperties `json:"channel_properties" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
}

type BillingInformation struct {
	ReferenceID   string `json:"reference_id,omitempty"`
	Country       string `json:"country"`
	StreetLine1   string `json:"street_line1"`
	StreetLine2   string `json:"street_line2"`
	City          string `json:"city"`
	ProvinceState string `json:"province_state"`
	PostalCode    string `json:"postal_code"`
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
	PaymentMethod           PaymentMethod2          `json:"payment_method" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	Actions                 Actions                 `json:"actions" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	CaptureMethod           string                  `json:"capture_method"`
	Initiator               string                  `json:"initiator"`
	ChannelProperties       ChannelProperties       `json:"channel_properties" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	ShippingInformation     ShippingInformation2    `json:"shipping_information" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	CardVerificationResults CardVerificationResults `json:"card_verification_result" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	FailureCode             string                  `json:"failure_code"`
	Created                 string                  `json:"created"`
	Updated                 string                  `json:"updated"`
	Metadata                map[string]interface{}  `json:"metadata,omitempty"`
}

type ShippingInformation2 struct {
	Country       string `json:"country"`
	StreetLine1   string `json:"street_line1"`
	StreetLine2   string `json:"street_line2"`
	City          string `json:"city"`
	ProvinceState string `json:"province_state"`
	PostalCode    string `json:"postal_code"`
}

type CardVerificationResults struct {
	ThreeDSecure              ThreeDSecure `json:"three_d_secure" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	CvvResult                 string       `json:"cvv_result"`
	AddressVerificationResult string       `json:"addres_verification_result"`
}

type ThreeDSecure struct {
	ReferenceID         string `json:"reference_id,omitempty"`
	ThreeDSecureFlow    string `json:"three_d_Secure_flow"`
	EciCode             string `json:"eci_code"`
	ThreeDSecureResult  string `json:"three_d_secure_result"`
	ThreeDSecureVersion string `json:"three_d_secure_version"`
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
