package xendit

// CustomerAddress contains data from Xendit's API response of customer's Customer Addres requests.
// For more details see https://xendit.github.io/apireference/?bash#customers.
// For documentation of subpackage customer, checkout https://pkg.go.dev/github.com/ianeinser/xendit-go/customer/
type CustomerAddress struct {
	ReferenceID string `json:"reference_id"`
	Country     string `json:"country" validate:"required"`
	StreetLine1 string `json:"street_line1,omitempty"`
	StreetLine2 string `json:"street_line2,omitempty"`
	City        string `json:"city,omitempty"`
	Province    string `json:"province,omitempty"`
	State       string `json:"state,omitempty"`
	PostalCode  string `json:"postal_code,omitempty"`
	Category    string `json:"category,omitempty"`
	IsPreferred bool   `json:"is_preferred,omitempty"`
}

// Customer contains data from Xendit's API response of customer related requests.
// For more details see https://xendit.github.io/apireference/?bash#customers.
// For documentation of subpackage customer, checkout https://pkg.go.dev/github.com/ianeinser/xendit-go/customer/
type Customer struct {
	ID               string                 `json:"id"`
	ReferenceID      string                 `json:"reference_id"`
	Type             string                 `json:"type"`
	IndividualDetail IndividualDetail       `json:"individual_detail" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	BusinessDetail   BusinessDetail         `json:"business_detail" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	MobileNumber     string                 `json:"mobile_number,omitempty"`
	Email            string                 `json:"email,omitempty"`
	GivenNames       string                 `json:"given_names"`
	MiddleName       string                 `json:"middle_name"`
	Surname          string                 `json:"surname"`
	Description      string                 `json:"description,omitempty"`
	PhoneNumber      string                 `json:"phone_number"`
	Nationality      string                 `json:"nationality"`
	Addresses        []CustomerAddress      `json:"addresses" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	DateOfBirth      string                 `json:"date_of_birth"`
	Metadata         map[string]interface{} `json:"metadata"`
}

type Employment struct {
	ReferenceID      string `json:"reference_id"`
	EmployerName     string `json:"employer_name"`
	NatureOfBusiness string `json:"nature_of_business"`
	RoleDescription  string `json:"role_description"`
}

type IndividualDetail struct {
	ReferenceID  string     `json:"reference_id"`
	GivenNames   string     `json:"given_names"`
	Surname      string     `json:"surname"`
	Nationality  string     `json:"nationality"`
	PlaceOfBirth string     `json:"place_of_birth"`
	DateOfBirth  string     `json:"date_of_birth"`
	Gender       string     `json:"gender"`
	Employment   Employment `json:"employment" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
}

type BusinessDetail struct {
	ReferenceID        string `json:"reference_id"`
	BusinessName       string `json:"business_detail"`
	TradingName        string `json:"trading_name"`
	BusinessType       string `json:"business_type"`
	NatureOfBusiness   string `json:"nature_of_business"`
	BusinessDomicile   string `json:"business_domicile"`
	DateOfRegistration string `json:"date_of_registration"`
}
