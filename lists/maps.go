package main

import "fmt"

func main() {
	websites := map[string]string{
		"Golang": "https://go.dev/",
		"Google": "https://www.google.com",
		"Amazon": "https://www.amazon.com",
	}

	fmt.Println(websites["Golang"])

	websites["Github"] = "https://github.com"
	fmt.Println(websites)

	delete(websites, "Amazon")
	fmt.Println(websites)
}
