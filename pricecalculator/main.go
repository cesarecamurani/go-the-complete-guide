package main

import (
	"example.com/pricecalculator/filemanager"
	"example.com/pricecalculator/prices"
	"fmt"
)

func main() {
	var taxRates = []float64{0.0, 0.07, 0.1, 0.15}
	doneChans := make([]chan bool, len(taxRates))
	errorChans := make([]chan error, len(taxRates))

	for index, taxRate := range taxRates {
		doneChans[index] = make(chan bool)
		errorChans[index] = make(chan error)

		fm := filemanager.New("prices.txt", fmt.Sprintf("taxed_prices_%.0f.json", taxRate*100))
		//cmdm := cmdmanager.New()
		taxRateJob := prices.NewTaxedPriceJob(fm, taxRate)
		//taxRateJob := prices.NewTaxedPriceJob(cmdm, taxRate)

		go taxRateJob.Process(doneChans[index], errorChans[index])
	}

	for index := range taxRates {
		select {
		case err := <-errorChans[index]:
			if err != nil {
				fmt.Printf("Error processing tax rate: %v\n", err)
			}
		case <-doneChans[index]:
			fmt.Printf("Successfully processed tax rate: %.2f\n", taxRates[index])
		}
	}
}
