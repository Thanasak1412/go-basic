package main

import (
	"fmt"
	"github.com/Thanasak1412/http-request/currencyapi"
	"strings"
	"sync"
)

func main() {
	currencies := []string{"BTC", "SUI", "ETH"}
	var wg sync.WaitGroup

	for _, currency := range currencies {
		wg.Add(1)
		go func(currencyCode string) {
			printRate(currencyCode)
			wg.Done()
		}(currency)
	}
	wg.Wait()
}

func printRate(currencyCode string) {
	fmt.Printf("Fetching the price for %v \n", currencyCode)

	price, err := currencyapi.GetRate(strings.ToUpper(currencyCode))

	if err != nil {
		fmt.Println("Error fetching the price for ", currencyCode, err)
	} else {
		fmt.Printf("The currency price of %v is %.2f\n", price.CurrencyCode, price.Price)
	}
}
