package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ianeinser/xendit-go"
	"github.com/ianeinser/xendit-go/customer"
	"github.com/joho/godotenv"
)

func main() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

	customerAddress := xendit.CustomerAddress{
		Country:     "ID",
		StreetLine1: "Jl. 123",
		StreetLine2: "Jl. 456",
		City:        "Jakarta Selatan",
		Province:    "DKI Jakarta",
		State:       "-",
		PostalCode:  "12345",
	}

	metadata := map[string]interface{}{
		"meta": "data",
	}

	createCustomerData := customer.CreateCustomerParams{
		ReferenceID: "test-reference-id-003",
		Email:       "tes@tes.com",
		Addresses:   []xendit.CustomerAddress{customerAddress},
		Metadata:    metadata,
	}

	resp, err := customer.CreateCustomer(&createCustomerData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created customer: %+v\n", resp)

	getCustomerByReferenceIDData := customer.GetCustomerByReferenceIDParams{
		ReferenceID: resp.ReferenceID,
	}

	resps, errs := customer.GetCustomerByReferenceID(&getCustomerByReferenceIDData)
	if errs != nil {
		log.Fatal(errs)
	}
	fmt.Printf("retrieved customer: %+v\n", resps)
}
