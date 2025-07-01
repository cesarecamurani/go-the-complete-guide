package main

import "fmt"

const AdultAge = 18

func main() {
	age := 42

	agePointer := &age

	fmt.Println("Age:", *agePointer)

	adultYears := getAdultYears(*agePointer)
	fmt.Println("Years since adulthood:", adultYears)

	//getAdultYears(agePointer)
	//fmt.Println("Years since adulthood:", age)
}

// getAdultYears calculates the number of years a person has spent as an adult based on the given age.
func getAdultYears(age int) int {
	return age - AdultAge
}

// getAdultYears modifies the given age pointer to represent the number of years since reaching adulthood.
//func getAdultYears(age *int) {
//	*age -= AdultAge
//}
