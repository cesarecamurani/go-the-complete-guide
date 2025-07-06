package main

import "fmt"

type multiplyFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}

	doubled := doubleNumbers(&numbers)
	doubledWithMultiply := multiplyNumbers(&numbers, double)
	doubledWithAnonymousFn := multiplyNumbers(&numbers, func(num int) int {
		return num * 2
	})
	doubledWithClosures := multiplyNumbers(&numbers, createMultiplier(2))

	tripledNumbers := multiplyNumbers(&numbers, triple)
	tripledWithClosures := multiplyNumbers(&numbers, createMultiplier(3))

	fmt.Println(numbers)
	fmt.Println(doubled)
	fmt.Println(doubledWithMultiply)
	fmt.Println(doubledWithAnonymousFn)
	fmt.Println(doubledWithClosures)
	fmt.Println(tripledNumbers)
	fmt.Println(tripledWithClosures)

	// Recursion example
	fact := factorial(6)
	fmt.Println(fact)

	// Variadic function example
	total := sum(3, 1, 2, 3, 4, 5, 10)
	fmt.Println(total)

	// With slices
	sliceTotal := sum(3, numbers...)
	fmt.Println(sliceTotal)
}

func multiplyNumbers(numbers *[]int, multiply multiplyFn) []int {
	var multipliedNums []int

	for _, num := range *numbers {
		multipliedNums = append(multipliedNums, multiply(num))
	}

	return multipliedNums
}

func doubleNumbers(numbers *[]int) []int {
	var dNums []int

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

// Closure
func createMultiplier(factor int) multiplyFn {
	return func(number int) int {
		return number * factor
	}
}

// Recursion
func factorial(number int) int {
	if number == 0 {
		return 1
	}

	return number * factorial(number-1)
}

// Variadic function (accepts a variable number of arguments with the three dots)
func sum(initialValue int, numbers ...int) int {
	total := initialValue

	for _, num := range numbers {
		total += num
	}

	return total
}
