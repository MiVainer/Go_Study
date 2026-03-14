package main

import (
	"fmt"

)

//простой сумматор
func sum(a int, b int) int {
	return a + b
}
//функция возвращающее частное и остаток от деления
func divide(divisible, divisor int) (int, int) {
    quotient := divisible / divisor
    remainder := divisible % divisor
    return quotient, remainder
}

func main() {

	fmt.Println(sum(2, 3))

	// Можно разделить переменные
	q, r := divide(10, 3)
	fmt.Println("Частное равно: ", q)
	fmt.Println("Остаток равен: ", r)
}