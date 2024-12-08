package main

import (
	"github.com/Thanasak1412/calculate-tax/filemanager"
	"github.com/Thanasak1412/calculate-tax/prices"
	"log"
)

func main() {
	taxRate := 0.7

	tax := *prices.NewTaxIncludedPriceJob(taxRate)

	tax.Process()

	err := filemanager.WriteJson("result.json", tax)
	if err != nil {
		log.Fatal(err)

		return
	}
}
