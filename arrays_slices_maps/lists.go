package main

import "fmt"

// Type alias
type floatMap map[string]float64

func (fm floatMap) print() {
	fmt.Println(fm)
}

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

	// Slices (DON'T MAKE COPIES IF MANIPULATED CHANGE THE ORIGINAL ARRAY!)
	discountedPrices := prices[1:3]
	fmt.Println(discountedPrices)
	fmt.Println(len(discountedPrices), cap(discountedPrices))

	// Starting directly with slices
	dynamicPrices := []float64{7.99, 8.99, 9.99}

	dynamicPrices[2] = 10.99
	updatedPrices := append(dynamicPrices, 6.99)

	fmt.Println(updatedPrices)

	// Make function (Slices)
	userNames := make([]string, 3, 10)

	userNames[0] = "Cesare"
	userNames[1] = "Fred"
	userNames[2] = "Mary"
	userNames = append(userNames, "John")
	userNames = append(userNames, "Jane")
	userNames = append(userNames, "James")

	fmt.Println(userNames)
	fmt.Println(len(userNames), cap(userNames))

	for _, name := range userNames {
		fmt.Println(name)
	}

	// Make a function (Maps) with a type alias
	ratings := make(floatMap, 3)

	ratings["go"] = 4.8
	ratings["react"] = 4.6
	ratings["ruby"] = 4.5

	ratings.print()

	for language, rating := range ratings {
		fmt.Printf("%v has a rating of %v\n", language, rating)
	}
}
