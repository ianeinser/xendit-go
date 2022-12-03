package main

import (
	"fmt"
	"log"
	"time"

	"github.com/ianeinser/xendit-go"
	"github.com/ianeinser/xendit-go/customer"
)

func customerTest() {
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
		ReferenceID: time.Now().String(),
		Email:       "tes@tes.com",
		Addresses:   []xendit.CustomerAddress{customerAddress},
		Metadata:    metadata,
	}

	resp, err := customer.CreateCustomer(&createCustomerData)
	if err != nil {
		log.Panic(err)
	}
	getCustomerByReferenceIDData := customer.GetCustomerByReferenceIDParams{
		ReferenceID: resp.ReferenceID,
	}

	_, err = customer.GetCustomerByReferenceID(&getCustomerByReferenceIDData)
	if err != nil {
		log.Panic(err)
	}

	fmt.Println("Customer integration tests done!")
}
