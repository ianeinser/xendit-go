package payment_test

import (
	"fmt"
	"log"

	"github.com/ianeinser/xendit-go"
	"github.com/ianeinser/xendit-go/customer"
	"github.com/ianeinser/xendit-go/payment"
)

func ExampleCreatePaymentMethod() {
	xendit.Opt.SecretKey = "examplesecretkey"

	billingInformation := xendit.BillingInformation{
		Country:       "ID",
		StreetLine1:   "Jl. 123",
		StreetLine2:   "Jl. 456",
		City:          "Jakarta Selatan",
		ProvinceState: "DKI Jakarta",
		PostalCode:    "12345",
	}

	metadata := map[string]interface{}{
		"meta": "data",
	}

	data := payment.CreatePaymentMethodParams{
		Type:               "",
		Reusability:        "",
		ReferenceID:        "",
		CustomerID:         "",
		Country:            "",
		Description:        "",
		BillingInformation: billingInformation,
		Metadata:           metadata,
	}

	resp, err := payment.CreatePaymentMethod(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created payment method: %+v\n", resp)
}

func ExampleGeListPaymentsByPaymentMethodId() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := customer.GetCustomerByReferenceIDParams{
		ReferenceID: "test-reference-id-002",
	}

	resp, err := customer.GetCustomerByReferenceID(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved customer: %+v\n", resp)
}
