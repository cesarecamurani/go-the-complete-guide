package main

import (
	"example.com/price-calculator/filemanager"
	"example.com/price-calculator/prices"
	"fmt"
)

func main() {
	var taxRates = []float64{0.0, 0.07, 0.1, 0.15}

	for _, taxRate := range taxRates {
		fm := filemanager.New("prices.txt", fmt.Sprintf("taxed_prices_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		taxRateJob := prices.NewTaxedPriceJob(fm, taxRate)
		//taxRateJob := prices.NewTaxedPriceJob(cmdm, taxRate)
		err := taxRateJob.Process()
		if err != nil {
			fmt.Println("Error processing job")
			fmt.Println(err)
		}
	}
}
