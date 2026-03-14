package main

import (
    "fmt"


)


func main() {
	var a, n, sum int
	fmt.Println("Введи целое число из нескольких знаков:")
	fmt.Scan(&a)
	n = 0
	sum = 0
	for a > 0 {
		n = a % 10
		sum+=n
		a = a / 10
	}
	fmt.Println("Сумма цифр числа:", sum)
	}




