package main

import (
	"example.com/structs/user"
	"fmt"
)

type str string

func (text str) log() {
	fmt.Println(text)
}

func main() {
	var name str = "Welcome!"

	name.log()

	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthDate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(firstName, lastName, birthDate)
	if err != nil {
		fmt.Println(err)
		return
	}

	adminUser := user.NewAdmin("test@example.com", "123Password")

	adminUser.OutputUserDetails()
	adminUser.ClearUserName()
	adminUser.OutputUserDetails()

	// With pointer (general method)
	//outputUserDetails(&appUser)

	// With method attached to user struct
	appUser.OutputUserDetails()
	appUser.ClearUserName()
	appUser.OutputUserDetails()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)

	var value string

	_, err := fmt.Scanln(&value)
	if err != nil {
		return ""
	}

	return value
}

//func outputUserDetails(user *user) {
//	formattedCreatedAt := user.createdAt.Format("2006-01-02 15:04:05")
//
//	// No dereferencing needed for structs (Go shortcut)
//	fmt.Println(user.firstName, user.firstName, user.birthDate, formattedCreatedAt)
//
//	// Also correct, technically the right way (dereferencing - pick the value from pointer - before accessing it)
//	fmt.Println((*user).firstName, (*user).firstName, (*user).birthDate, formattedCreatedAt)
//}
