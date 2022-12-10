package xendit

import "gorm.io/datatypes"

// CustomerAddress contains data from Xendit's API response of customer's Customer Addres requests.
// For more details see https://xendit.github.io/apireference/?bash#customers.
// For documentation of subpackage customer, checkout https://pkg.go.dev/github.com/ianeinser/xendit-go/customer/
type CustomerAddress struct {
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
	ID                string            `json:"id"`
	ReferenceID       string            `json:"reference_id"`
	Type              string            `json:"type"`
	IndividualDetail  IndividualDetail  `json:"individual_detail" gorm:"embedded;embedded_prefix:id_"`
	BusinessDetail    BusinessDetail    `json:"business_detail" gorm:"embedded;embedded_prefix:bd_"`
	MobileNumber      string            `json:"mobile_number,omitempty"`
	Email             string            `json:"email,omitempty"`
	Description       string            `json:"description,omitempty"`
	PhoneNumber       string            `json:"phone_number"`
	HashedPhoneNumber string            `json:"hashed_phone_number"`
	Addresses         []CustomerAddress `json:"addresses" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	Metadata          datatypes.JSONMap `json:"metadata"`
}

type Employment struct {
	EmployerName     string `json:"employer_name"`
	NatureOfBusiness string `json:"nature_of_business"`
	RoleDescription  string `json:"role_description"`
}

type IndividualDetail struct {
	GivenNames   string     `json:"given_names"`
	Surname      string     `json:"surname"`
	Nationality  string     `json:"nationality"`
	PlaceOfBirth string     `json:"place_of_birth"`
	DateOfBirth  string     `json:"date_of_birth"`
	Gender       string     `json:"gender"`
	Employment   Employment `json:"employment" gorm:"embedded;embedded_prefix:emp_"`
}

type BusinessDetail struct {
	BusinessName       string `json:"business_name"`
	BusinessType       string `json:"business_type"`
	NatureOfBusiness   string `json:"nature_of_business"`
	BusinessDomicile   string `json:"business_domicile"`
	DateOfRegistration string `json:"date_of_registration"`
}
