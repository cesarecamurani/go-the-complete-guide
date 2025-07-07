package prices

import (
	"example.com/price-calculator/conversion"
	"example.com/price-calculator/filemanager"
	"fmt"
)

type taxPrice map[string]string

type TaxedPriceJob struct {
	TaxRate     float64
	Prices      []float64
	TaxedPrices taxPrice
}

func NewTaxedPriceJob(taxRate float64) *TaxedPriceJob {
	return &TaxedPriceJob{
		TaxRate: taxRate,
		Prices:  []float64{},
	}
}

func (job *TaxedPriceJob) LoadPrices() {
	lines, err := filemanager.ReadLInesFromFile("prices.txt")
	if err != nil {
		fmt.Printf("failed to read lines from file: %v", err)
		return
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		fmt.Printf("failed to convert strings to float: %v", err)
		return
	}

	job.Prices = prices
}

func (job *TaxedPriceJob) Process() {
	job.LoadPrices()

	result := make(taxPrice, len(job.Prices))

	for _, price := range job.Prices {
		taxedPrice := price + (price * job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxedPrice)
	}

	job.TaxedPrices = result

	err := filemanager.WriteJSON(fmt.Sprintf("taxed_prices_%.0f.json", job.TaxRate*100), job)
	if err != nil {
		return
	}
}
