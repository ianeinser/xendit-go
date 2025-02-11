package retailoutlet_test

import (
	"fmt"
	"log"
	"time"

	"github.com/ianeinser/xendit-go"
	"github.com/ianeinser/xendit-go/retailoutlet"
)

func ExampleCreateFixedPaymentCode() {
	xendit.Opt.SecretKey = "examplesecretkey"

	data := retailoutlet.CreateFixedPaymentCodeParams{
		ExternalID:       "retailoutlet-external-id",
		RetailOutletName: xendit.RetailOutletNameAlfamart,
		Name:             "Michael Jackson",
		ExpectedAmount:   200000,
	}

	resp, err := retailoutlet.CreateFixedPaymentCode(&data)
	if err != nil {
		log.Fatal(err.ErrorCode)
	}

	fmt.Printf("created retail outlet fixed payment code: %+v\n", resp)
}

func ExampleGetFixedPaymentCode() {
	xendit.Opt.SecretKey = "examplesecretkey"

	getFixedPaymentCodeData := retailoutlet.GetFixedPaymentCodeParams{
		FixedPaymentCodeID: "123",
	}

	resp, err := retailoutlet.GetFixedPaymentCode(&getFixedPaymentCodeData)
	if err != nil {
		log.Fatal(err.ErrorCode)
	}

	fmt.Printf("retrieved retail outlet fixed payment code: %+v\n", resp)
}

func ExampleGetPaymentByFixedPaymentCode() {
	xendit.Opt.SecretKey = "examplesecretkey"

	getPaymentByFixedPaymentCodeData := retailoutlet.GetPaymentByFixedPaymentCodeParams{
		FixedPaymentCodeID: "123",
	}

	resp, err := retailoutlet.GetPaymentByFixedPaymentCode(&getPaymentByFixedPaymentCodeData)
	if err != nil {
		log.Fatal(err.ErrorCode)
	}

	fmt.Printf("retrieved retail outlet list payment by fixed payment code %+v\n", resp)
}

func ExampleUpdateFixedPaymentCode() {
	xendit.Opt.SecretKey = "examplesecretkey"

	expirationDate := time.Now().AddDate(0, 0, 1)

	updateFixedPaymentCodeData := retailoutlet.UpdateFixedPaymentCodeParams{
		FixedPaymentCodeID: "123",
		Name:               "Billy Jackson",
		ExpectedAmount:     2000000,
		ExpirationDate:     &expirationDate,
	}

	resp, err := retailoutlet.UpdateFixedPaymentCode(&updateFixedPaymentCodeData)
	if err != nil {
		log.Fatal(err.ErrorCode)
	}
	fmt.Printf("updated retail outlet fixed payment code: %+v\n", resp)
}
