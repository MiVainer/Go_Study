package main

import (
	"fmt"

)

func addptr(nptr *int, delta int) {
    *nptr += delta
}


func main() {
	
	n := 42
	addptr(&n, 3)
	fmt.Println(n)
}