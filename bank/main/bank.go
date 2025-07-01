package main

import (
	"example.com/bank/fileops"
	"fmt"
	"github.com/Pallinder/go-randomdata"
)

const BalanceFileName = "balance.txt"

func main() {
	var accountBalance, err = fileops.FetchFloatFromFile(BalanceFileName)

	if err != nil {
		fmt.Println("Failed to fetch balance")
		fmt.Println("Error:", err)
		fmt.Println("----------")
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Help desk phone number:", randomdata.PhoneNumber())

	for {
		displayOptions()

		var userChoice int

		fmt.Print("Your choice: ")

		_, scanErr := fmt.Scan(&userChoice)
		if scanErr != nil {
			return
		}

		switch userChoice {
		case 1:
			fmt.Println("Your current balance is", accountBalance)
			fmt.Println()
		case 2:
			handleDeposit(accountBalance)
		case 3:
			handleWithdrawal(accountBalance)
		case 4:
			fmt.Println("Thanks for choosing Go Bank!")
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Please select a valid option")
			fmt.Println()
		}
	}
}
