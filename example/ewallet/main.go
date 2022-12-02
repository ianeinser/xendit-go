package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/ianeinser/xendit-go"
	"github.com/ianeinser/xendit-go/ewallet"
	"github.com/joho/godotenv"
)

func main() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

	createPaymentData := ewallet.CreatePaymentParams{
		ExternalID:  "ovo-" + time.Now().String(),
		Amount:      10000,
		Phone:       "081234567890",
		EWalletType: xendit.EWalletTypeOVO,
	}

	resp, err := ewallet.CreatePayment(&createPaymentData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("created payment: %+v\n", resp)

	getPaymentStatusData := ewallet.GetPaymentStatusParams{
		ExternalID:  resp.ExternalID,
		EWalletType: resp.EWalletType,
	}

	resp, err = ewallet.GetPaymentStatus(&getPaymentStatusData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved payment: %+v\n", resp)

	metadata := map[string]interface{}{
		"meta": "data",
	}

	ewalletBasketItem := xendit.EWalletBasketItem{
		ReferenceID: "basket-product-ref-id",
		Name:        "product name",
		Category:    "mechanics",
		Currency:    "IDR",
		Price:       50000,
		Quantity:    5,
		Type:        "type",
		SubCategory: "subcategory",
		Metadata:    metadata,
	}

	createEWalletChargeData := ewallet.CreateEWalletChargeParams{
		ReferenceID:    "test-reference-id",
		Currency:       "IDR",
		Amount:         1688,
		CheckoutMethod: "ONE_TIME_PAYMENT",
		ChannelCode:    "ID_SHOPEEPAY",
		ChannelProperties: map[string]string{
			"success_redirect_url": "https://yourwebsite.com/order/123",
		},
		Basket: []xendit.EWalletBasketItem{
			ewalletBasketItem,
		},
		Metadata: metadata,
	}

	charge, chargeErr := ewallet.CreateEWalletCharge(&createEWalletChargeData)
	if chargeErr != nil {
		log.Fatal(chargeErr)
	}
	fmt.Printf("created e-wallet charge: %+v\n", charge)

	getEWalletChargeStatusData := ewallet.GetEWalletChargeStatusParams{
		ChargeID: charge.ID,
	}

	charge, chargeErr = ewallet.GetEWalletChargeStatus(&getEWalletChargeStatusData)
	if chargeErr != nil {
		log.Fatal(chargeErr)
	}
	fmt.Printf("retrieved e-wallet charge: %+v\n", charge)
}
