package main

import (
	"fmt"

)

func addval(nval int, delta int) {
    nval += delta
}


func main() {
	
	n := 42
	addval(n, 3)
	fmt.Println(n)
}