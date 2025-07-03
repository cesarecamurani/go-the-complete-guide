package main

import "fmt"

func main() {
	// Arrays
	//var fruits [4]string
	//fruits = [4]string{"Apple", "Banana", "Orange", "Pear"}

	var fruits = [4]string{"Apple"}
	fruits[1] = "Orange"

	prices := [4]float64{10.99, 9.99, 11.99, 12.0}

	fmt.Println(prices)
	fmt.Println(prices[3])
	fmt.Println(fruits)

	// Slices (DON'T MAKE COPIES, IF MANIPULATED CHANGE THE ORIGINAL ARRAY!)
	discountedPrices := prices[1:3]
	fmt.Println(discountedPrices)
	fmt.Println(len(discountedPrices), cap(discountedPrices))

	// Starting directly with slices
	dynamicPrices := []float64{7.99, 8.99, 9.99}

	dynamicPrices[2] = 10.99
	updatedPrices := append(dynamicPrices, 6.99)

	fmt.Println(updatedPrices)
}
