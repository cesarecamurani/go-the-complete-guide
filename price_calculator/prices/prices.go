package prices

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const bitSize = 64

type TaxedPriceJob struct {
	TaxRate     float64
	Prices      []float64
	TaxedPrices map[string]float64
}

func NewTaxedPriceJob(taxRate float64) *TaxedPriceJob {
	return &TaxedPriceJob{
		TaxRate: taxRate,
		Prices:  []float64{},
	}
}

func (job *TaxedPriceJob) LoadPrices() {
	file, err := os.Open("prices.txt")
	if err != nil {
		fmt.Printf("failed to open prices file: %v", err)
		return
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		fmt.Printf("failed to open prices file: %v", err)
		return
	}

	prices := make([]float64, len(lines))

	for index, line := range lines {
		floatPrice, parseErr := strconv.ParseFloat(line, bitSize)
		if parseErr != nil {
			return
		}

		prices[index] = floatPrice
	}

	job.Prices = prices
}

func (job *TaxedPriceJob) Process() {
	job.LoadPrices()

	result := make(map[string]string, len(job.Prices))

	for _, price := range job.Prices {
		taxedPrice := price + (price * job.TaxRate)
		result[fmt.Sprintf("%.2f", price)] = fmt.Sprintf("%.2f", taxedPrice)
	}

	fmt.Println(result)
}
