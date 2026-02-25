package main

import (
	"fmt"
)

func main() {
	var source string
	var times int
	fmt.Scan(&source, &times)
	i := 1
	for i <= times {
		fmt.Print(source)
		i ++
	}

}