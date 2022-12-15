package xendit

import "gorm.io/datatypes"

// CustomerAddress contains data from Xendit's API response of customer's Customer Addres requests.
// For more details see https://xendit.github.io/apireference/?bash#customers.
// For documentation of subpackage customer, checkout https://pkg.go.dev/github.com/ianeinser/xendit-go/customer/
type CustomerAddress struct {
	ReferenceID string `json:"-"`
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
	IndividualDetail  IndividualDetail  `json:"individual_detail,omitempty" gorm:"embedded;embeddedPrefix:id_"`
	BusinessDetail    BusinessDetail    `json:"business_detail,omitempty" gorm:"embedded;embeddedPrefix:bd_"`
	MobileNumber      string            `json:"mobile_number,omitempty"`
	Email             string            `json:"email,omitempty"`
	Description       string            `json:"description,omitempty"`
	PhoneNumber       string            `json:"phone_number"`
	HashedPhoneNumber string            `json:"hashed_phone_number"`
	Addresses         []CustomerAddress `json:"addresses,omitempty" gorm:"foreignKey:ReferenceID;references:ReferenceID"`
	Metadata          datatypes.JSONMap `json:"metadata"`
}

type Employment struct {
	EmployerName     string `json:"employer_name,omitempty"`
	NatureOfBusiness string `json:"nature_of_business,omitempty"`
	RoleDescription  string `json:"role_description,omitempty"`
}

type IndividualDetail struct {
	GivenNames   string     `json:"given_names,omitempty"`
	Surname      string     `json:"surname,omitempty"`
	Nationality  string     `json:"nationality,omitempty"`
	PlaceOfBirth string     `json:"place_of_birth,omitempty"`
	DateOfBirth  string     `json:"date_of_birth,omitempty"`
	Gender       string     `json:"gender,omitempty"`
	Employment   Employment `json:"employment,omitempty" gorm:"embedded;embeddedPrefix:emp_"`
}

type BusinessDetail struct {
	BusinessName       string `json:"business_name,omitempty"`
	BusinessType       string `json:"business_type,omitempty"`
	NatureOfBusiness   string `json:"nature_of_business,omitempty"`
	BusinessDomicile   string `json:"business_domicile,omitempty"`
	DateOfRegistration string `json:"date_of_registration,omitempty"`
}
