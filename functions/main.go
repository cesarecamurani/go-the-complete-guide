package main

import "fmt"

type multiFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	doubled := doubleNumbers(&numbers)
	doubledWithMultiply := multiplyNumbers(&numbers, double)
	doubledWithAnonymousFn := multiplyNumbers(&numbers, func(num int) int {
		return num * 2
	})

	tripledNumbers := multiplyNumbers(&numbers, triple)

	fmt.Println(numbers)
	fmt.Println(doubled)
	fmt.Println(doubledWithMultiply)
	fmt.Println(doubledWithAnonymousFn)
	fmt.Println(tripledNumbers)
}

func multiplyNumbers(numbers *[]int, multiply multiFn) []int {
	multipliedNums := []int{}

	for _, num := range *numbers {
		multipliedNums = append(multipliedNums, multiply(num))
	}

	return multipliedNums
}

func doubleNumbers(numbers *[]int) []int {
	dNums := []int{}

	for _, num := range *numbers {
		dNums = append(dNums, double(num))
	}

	return dNums
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}
