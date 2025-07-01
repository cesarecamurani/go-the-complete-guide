package main

import (
	"example.com/bank/fileops"
	"fmt"
)

func displayOptions() {
	fmt.Println("Please select an option:")
	fmt.Println("1. Check Balance")
	fmt.Println("2. Deposit Money")
	fmt.Println("3. Withdraw Money")
	fmt.Println("4. Exit")
}

func handleDeposit(accountBalance float64) {
	fmt.Println("How much would you like to deposit?")

	var depositAmount float64

	_, err := fmt.Scan(&depositAmount)
	if err != nil {
		return
	}

	if depositAmount <= 0 {
		fmt.Println("Please enter a positive amount")
		return
	}

	accountBalance += depositAmount

	fmt.Println("Your updated balance is", accountBalance)
	fmt.Println()

	fileops.SaveFloatToFile(accountBalance, BalanceFileName)
}

func handleWithdrawal(accountBalance float64) {
	fmt.Println("How much would you like to withdraw?")

	var withdrawAmount float64

	_, err := fmt.Scan(&withdrawAmount)
	if err != nil {
		return
	}

	if withdrawAmount <= 0 {
		fmt.Println("Please enter a positive amount")
		fmt.Println()
		return
	}

	if withdrawAmount > accountBalance {
		fmt.Println("The amount you chose is greater than your balance!")
		fmt.Println()
		return
	}

	accountBalance -= withdrawAmount

	fmt.Println("Your updated balance is", accountBalance)
	fmt.Println()

	fileops.SaveFloatToFile(accountBalance, BalanceFileName)
}
