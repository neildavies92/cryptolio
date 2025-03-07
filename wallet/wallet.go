package wallet

import (
	"fmt"
	"net/http"
)

func GetWallet(jwt string) (*http.Response, error) {

	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.coinbase.com/v2/accounts/", nil)
	req.Header.Add("Authorization", "Bearer "+jwt)
	req.Header.Add("Content-Type", "application/json")
	if err != nil {
		fmt.Printf("Error: %v", err)
		return nil, err
	}

	resp, err := client.Do(req)
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: request failed with status: ", resp.Status)
		return nil, err
	}

	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		return resp, nil
	}
	return nil, nil
}
