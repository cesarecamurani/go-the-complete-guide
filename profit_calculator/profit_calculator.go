package main

import (
	"errors"
	"fmt"
	"os"
)

const ReadWritePermissionCode = 0644
const ResultsFileName = "results.txt"

func main() {
	revenue, err := getUserInput("Total Revenue Amount: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	expenses, err := getUserInput("Total Expenses Amount: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	taxRate, err := getUserInput("Tax Rate: ")
	if err != nil {
		fmt.Println(err)
		return
	}

	ebt, profit, ratio := calculateProfit(revenue, expenses, taxRate)

	saveResultsToFile(ebt, profit, ratio)

	fmt.Printf("Earnings before tax: %.2f\n", ebt)
	fmt.Printf("Profit: %.2f\n", profit)
	fmt.Printf("Ratio: %.2f\n", ratio)
}

func getUserInput(text string) (float64, error) {
	var userInput float64

	fmt.Print(text)

	_, err := fmt.Scan(&userInput)
	if err != nil {
		return userInput, err
	}

	if userInput <= 0 {
		return 0, errors.New("amount must be a positive number")
	}

	return userInput, nil
}

func calculateProfit(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit

	return ebt, profit, ratio
}

func saveResultsToFile(ebt, profit, ratio float64) {
	//values := []float64{ebt, profit, ratio}
	//
	//content := ""
	//for _, v := range values {
	//	content += fmt.Sprintf("%.2f\n", v)
	//}

	content := fmt.Sprintf("EBT: %.2f\nProfit: %.2f\nRatio: %.2f\n", ebt, profit, ratio)

	err := os.WriteFile(ResultsFileName, []byte(content), ReadWritePermissionCode)
	if err != nil {
		panic(err)
	}
}
