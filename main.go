package main

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/neildavies92/cryptolio/jwt"
	"github.com/neildavies92/cryptolio/wallet"
	log "github.com/sirupsen/logrus"
)

func loadEnvFile() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	loadEnvFile()

	requestMethod := "GET"
	requestHost := "api.coinbase.com"
	requestPath := "/v2/accounts/"

	uri := fmt.Sprintf("%s %s%s", requestMethod, requestHost, requestPath)

	jwt, err := jwt.BuildJWT(uri)

	if jwt != "" {
		fmt.Println("JWT set successfully")
	} else {
		log.Errorf("Error: Failed to build JWT: %v", err)
	}

	wallet, err := wallet.GetWallet(jwt)

	if wallet != nil {
		fmt.Println(wallet.Status)
		fmt.Println(wallet.Body)
	} else {
		log.Errorf("Error: Failed to retrieve wallet: %v", err)
	}
}
