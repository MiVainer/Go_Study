package main

import (
    "fmt"
	"strings"


)


func main() {
	var n int 
	var s string
	fmt.Scan(&n)
	s ="*"
	for i:= 1; i<=n; i++ {
		repeated := strings.Repeat(s, i)
		fmt.Printf(repeated + "\n")
	}

}




