package wallet

import (
	"fmt"
	"net/http"
)

func GetWallet(jwt string) (*http.Response, error) {

	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.coinbase.com/v2/accounts/", nil)

	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+jwt)

	resp, err := client.Do(req)

	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		return resp, nil
	}
	return nil, nil
}
