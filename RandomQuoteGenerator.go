package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

type Quote struct {
	Quote    string `json:"q"`
	Author   string `json:"a"`
	Category string `json:"c"`
}

func HandleRequest(ctx context.Context) (string, error) {
	response, err := http.Get("https://zenquotes.io/api/quotes/")

	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}

	var quotes []Quote
	err = json.Unmarshal(responseData, &quotes)

	return quotes[0].Quote, nil
}

func main() {
	lambda.Start(HandleRequest)
}
