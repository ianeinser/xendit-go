package payment_test

import (
	"fmt"
	"log"

	"github.com/ianeinser/xendit-go"
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

	ewallet := xendit.Ewallet{
		ChannelCode: "OVO",
		ChannelProperties: xendit.ChannelProperties{
			SuccessReturnUrl: "https://your-redirect-website.com/success",
			FailureReturnUrl: "https://your-redirect-website.com/failure",
		},
		Account: xendit.Account2{
			Name:    "Tony Stark",
			Balance: 1000000,
		},
	}

	metadata := map[string]interface{}{
		"meta": "data",
	}

	data := payment.CreatePaymentMethodParams{
		Type:               "EWALLET",
		Reusability:        "MULTIPLE_USE",
		ReferenceId:        "test-reference-id",
		CustomerId:         "fc4c060b-3c41-4707-b7b2-df9c3376edde",
		Country:            "USA",
		BillingInformation: billingInformation,
		Metadata:           metadata,
		Ewallet:            ewallet,
	}

	resp, err := payment.CreatePaymentMethod(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created payment method: %+v\n", resp)
}

func ExampleCreatePaymentRequest() {
	xendit.Opt.SecretKey = "examplesecretkey"

	billingInformation := xendit.BillingInformation{
		Country:       "ID",
		StreetLine1:   "Jl. 123",
		StreetLine2:   "Jl. 456",
		City:          "Jakarta Selatan",
		ProvinceState: "DKI Jakarta",
		PostalCode:    "12345",
	}

	ewallet := xendit.Ewallet{
		ChannelCode: "SHOPEEPAY",
		ChannelProperties: xendit.ChannelProperties{
			SuccessReturnUrl: "https://your-redirect-website.com/success",
			FailureReturnUrl: "https://your-redirect-website.com/failure",
		},
		Account: xendit.Account2{
			Name:    "Tony Stark",
			Balance: 1000000,
		},
	}

	metadata := map[string]interface{}{
		"sku": "ABCDEFGH",
	}

	data := payment.CreatePaymentRequestParams{
		Currency:    "IDR",
		Amount:      1000000,
		ReferenceId: "test-reference-id",
		CustomerId:  "fc4c060b-3c41-4707-b7b2-df9c3376edde",
		Country:     "USA",
		PaymentMethod: xendit.PaymentMethod2{
			Reusability:        "ONE_TIME_USE",
			Country:            "Indonesia",
			Type:               "EWALLET",
			Ewallet:            ewallet,
			BillingInformation: billingInformation,
			Metadata:           metadata,
		},
		PaymentMethodId:   "",
		ChannelProperties: xendit.ChannelProperties{},
		Metadata:          metadata,
	}

	resp, err := payment.CreatePaymentRequest(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created payment request: %+v\n", resp)
}

func ExampleCreateRefund() {
	xendit.Opt.SecretKey = "examplesecretkey"

	metadata := map[string]interface{}{
		"sku": "ABCDEFGH",
	}

	data := payment.CreateRefundParams{
		PaymentRequestId: "pr-6e9778ea-7d62-40fe-8b25-a4d740754c5f",
		ReferenceId:      "b2756a1e-e6cd-4352-9a68-0483aa2b6a2f",
		Currency:         "PHP",
		Amount:           1000000,
		Reason:           "CANCELLATION",
		Metadata:         metadata,
	}

	resp, err := payment.CreateRefund(&data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("created refund: %+v\n", resp)
}
