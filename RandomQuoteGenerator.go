package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type Quote struct {
	Quote    string `json:"q"`
	Author   string `json:"a"`
	Category string `json:"c"`
}

func main() {
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

	fmt.Println(quotes[0].Quote)
}
