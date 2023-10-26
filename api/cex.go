package api

import (
	"encoding/json"
	"fmt"
	"io"
	"jsflor/cryptomasters/data"
	"net/http"
	"strings"
)

const apiUrl = "https://cex.io/api/ticker/%s/USD"

func GetRate(currency string) (*data.Rate, error) {
	if len(currency) != 3 {
		return nil, fmt.Errorf("3 characters required; received %v\n", len(currency))
	}

	res, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(currency)))
	if err != nil {
		return nil, err
	}

	var response CEXResponse

	if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(body, &response)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("status code received: %v", res.StatusCode)
	}
	return &data.Rate{Currency: currency, Price: response.Bid}, nil
}
