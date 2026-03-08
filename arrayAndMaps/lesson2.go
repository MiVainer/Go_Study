package main

import (
	"fmt"

)

func main() {
	nums := []int{2, 3, 4}
	for idx, num := range nums {
		if num == 3 {
			fmt.Println("index:", idx)
		}
	}
}