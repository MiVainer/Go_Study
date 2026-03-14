package main

import (
    "fmt"
	"sort"
)

func minOfThree(a, b, c int) int {
	numbers:= []int{a, b, c}
	sort.Ints(numbers)
	return numbers[0]
}

func main() {
	var a,b,c int
	fmt.Println("Введи 3 целых числа для сравнения")
	fmt.Scan(&a, &b, &c)
	fmt.Println("Минимальное из трех чисел", minOfThree(a, b, c))
}




