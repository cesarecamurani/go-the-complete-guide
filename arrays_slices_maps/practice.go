package main

import "fmt"

// Time to practice what you learned!
// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.

// Practice
func main() {
	// 1)
	hobbies := [3]string{"Hiking", "Reading", "Cycling"}
	fmt.Println(hobbies)

	// 2)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	// 3)
	bestHobbies := hobbies[:2]
	fmt.Println(bestHobbies)

	// 4
	bestHobbies = bestHobbies[1:3]
	fmt.Println(bestHobbies)

	// 5
	goals := []string{"Learn a new language", "Create a Go app"}
	fmt.Println(goals)

	// 6
	goals[1] = "Use Go in a project"
	goals = append(goals, "Expand my skills")
	fmt.Println(goals)

	// Append another list to an existing one
	newGoals := []string{"For the sake of it", "Make more money"}
	goals = append(goals, newGoals...)
	fmt.Println(goals)

	// 7
	type Product struct {
		id    int
		title string
		price int
	}

	products := []Product{
		{id: 1, title: "Product One", price: 100},
		{id: 2, title: "Product Two", price: 200},
	}
	products = append(products, Product{id: 3, title: "Product Three", price: 300})

	fmt.Println(products)
}
