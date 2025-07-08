package prices

import (
	"example.com/price-calculator/conversion"
	"example.com/price-calculator/iomanager"
	"fmt"
)

type taxPriceType map[string]string

type TaxedPriceJob struct {
	IOManager   iomanager.IOManager `json:"-"`
	TaxRate     float64             `json:"tax_rate"`
	Prices      []float64           `json:"prices"`
	TaxedPrices taxPriceType        `json:"taxed_prices"`
}

func NewTaxedPriceJob(iom iomanager.IOManager, taxRate float64) *TaxedPriceJob {
	return &TaxedPriceJob{
		IOManager: iom,
		TaxRate:   taxRate,
		Prices:    []float64{},
	}
}

func (job *TaxedPriceJob) Process() error {
	err := job.loadPrices()
	if err != nil {
		return fmt.Errorf("failed to load prices: %v", err)
	}

	result := make(taxPriceType, len(job.Prices))

	for _, price := range job.Prices {
		taxedPrice := price + (price * job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxedPrice)
	}

	job.TaxedPrices = result

	fmt.Printf("Processed prices with tax rate %.2f\n", job.TaxRate)

	return job.IOManager.WriteResult(job)
}

func (job *TaxedPriceJob) loadPrices() error {
	lines, err := job.IOManager.ReadLines()
	if err != nil {
		return err
	}

	prices, err := conversion.StringsToFloat(lines)
	if err != nil {
		return err
	}

	job.Prices = prices

	return nil
}
