package main

import (
	"example.com/price-calculator/prices"
)

func main() {
	var taxRates = []float64{0.0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		taxRateJob := prices.NewTaxedPriceJob(taxRate)
		taxRateJob.Process()
	}
}
