package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Scan(&i)
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
		fallthrough
	case 3:
		fmt.Println("Goodbuy!")
	}
}