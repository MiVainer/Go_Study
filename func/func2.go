package main

import (
	"fmt"

)

func sum(nums ...int) {
    fmt.Print(nums, " -> ")
    total := 0
    for _, num := range nums {
        total += num
    }
    fmt.Println(total)
}

func main() {
	// Вариант 1
	sum(2, 4, 3, 12)

	// Вариант 2, передаем срез, преобразовав его в список аргументов с использованием ...
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}