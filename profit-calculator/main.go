package main

import (
	"fmt"
	"log"
)

func main() {
	var revenue, expenses int
	var ratio, profit, ebt, taxRate float64

	if _, err := fmt.Scan(&revenue, &expenses, &taxRate); err != nil {
		log.Fatal(err)

		return
	}

	ebt = float64(revenue - expenses)
	profit = ebt - (ebt * (taxRate / 100))
	ratio = ebt / profit

	fmt.Printf("profit: %.2f\n", profit)
	fmt.Printf("ratio: %.2f\n", ratio)
}
