package main

import (
	"fmt"
	"jsflor/cryptomasters/api"
)

func main() {
	rate, err := api.GetRate("BTC")
	if err == nil {
		fmt.Printf("The rate for %v is %2.f\n", rate.Currency, rate.Price)
	}
}
