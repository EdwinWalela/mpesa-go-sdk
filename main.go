package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const AUTH_ENDPOINT = "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"

func getAuth() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", AUTH_ENDPOINT, nil)

	if err != nil {
		panic(err)
	}

	consumerKey := os.Getenv("CONSUMER_KEY")
	consumerSecret := os.Getenv("CONSUMER_SECRET")

	secret := "Basic " + base64.StdEncoding.EncodeToString([]byte(consumerKey+":"+consumerSecret))

	req.Header.Add("Authorization", secret)

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.StatusCode)

	// return auth
}

func main() {
	err := godotenv.Load()

	if err != nil {
		panic(err)
	}

	getAuth()
}
