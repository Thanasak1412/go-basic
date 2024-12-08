package prices

import (
	"fmt"
	"github.com/Thanasak1412/calculate-tax/filemanager"
	"log"
	"math"
	"strconv"
	"strings"
)

type TaxIncludedPriceJob struct {
	InputPrice        []float64          `json:"inputPrice"`
	TaxRate           float64            `json:"taxRate"`
	TaxIncludedPrices map[string]float64 `json:"taxIncludedPrices"`
}

func (j *TaxIncludedPriceJob) Process() {
	result := make(map[string]float64)

	if err := j.LoadPrices(); err != nil {
		log.Fatal(err)

		return
	}

	for _, price := range j.InputPrice {
		result[fmt.Sprintf("%.2f", price)] = FixDecimal(price*(1+j.TaxRate), 2)
	}

	j.TaxIncludedPrices = result
}

func NewTaxIncludedPriceJob(taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		TaxRate: taxRate,
	}
}

func (j *TaxIncludedPriceJob) LoadPrices() error {
	content, err := filemanager.ReadLines("prices.txt")

	if err != nil {
		log.Fatal(err)

		return err
	}

	var inputPrices []float64
	for _, line := range strings.Split(content, "\n") {
		price, _ := strconv.ParseFloat(line, 64)
		value := FixDecimal(price, 2)

		inputPrices = append(inputPrices, value)
	}

	j.InputPrice = inputPrices

	return nil
}

// FixDecimal places to the desire precision and returns a float64
func FixDecimal(value float64, precision int) float64 {
	// Use math.Pow to shift the decimal point, round, and shift back
	scale := math.Pow10(precision)

	return math.Round(value*scale) / scale
}
