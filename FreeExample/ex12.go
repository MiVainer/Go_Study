package main

import (
	"fmt"
)

func factorial(n int) int {
	switch {
	case n < 0:
		return -1
	case n == 0:
		return 1
	default:
		return n * factorial(n-1)
	}
}

func main() {
	var n int
	fmt.Println("Введи требуемый факториал")
	fmt.Scan(&n)
	result := factorial(n)
	if result != -1 {
		fmt.Println("Факториал числа:", result)
	
	} else {
		fmt.Println("Ошибка, отрицательное число")
	}

	
}
