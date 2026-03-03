package main

import (
	"fmt"
	
)

func main() {
	var lang string
	fmt.Scan(&lang)
	switch {
	case lang == "en":
		fmt.Println("English")
	
	case lang == "fr":
		fmt.Println("French")

	case lang == "ru", lang == "rus":
		fmt.Println("Russian")
	default:
		fmt.Println("Unknown")
	}
}