package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ianeinser/xendit-go"
	"github.com/ianeinser/xendit-go/transaction"
	"github.com/joho/godotenv"
)

func main() {
	godotenvErr := godotenv.Load()
	if godotenvErr != nil {
		log.Fatal(godotenvErr)
	}
	xendit.Opt.SecretKey = os.Getenv("SECRET_KEY")

	getListTransactionData := transaction.GetListTransactionParams{
		Limit: 10,
	}

	respList, err := transaction.GetListTransaction(&getListTransactionData)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("retrieved list transaction: %+v\n", respList)

	getTransactionData := transaction.GetTransactionParams{
		TransactionID: respList.Data[0].ID,
	}

	resp, err := transaction.GetTransaction(&getTransactionData)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("retrieved transaction: %+v\n", resp)

}
