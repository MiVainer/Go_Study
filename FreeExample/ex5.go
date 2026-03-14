package main

import (
    "fmt"


)


func main() {
	var a, b float64
	fmt.Println("Введи 2 числа через пробел:")
	var c string
	fmt.Scan(&a, &b)
	fmt.Println("Выбери математическую операцию +, -, *, /:")
	fmt.Scan(&c)
	switch {
	case c == "+":
		fmt.Println("В результате сложения получилось:", a+b)
	case c == "-":
		fmt.Println("В результате вычитания получилось:", a-b)
	case c == "*":
		fmt.Println("В результате умножения получилось:", a*b)
	case c == "/":
		if b == 0 {
			fmt.Println("Ошибка, деление на 0. Введите исло отличное от 0.")
		} else {
			fmt.Println("В результате деления получилось:", a/b)
		}
	default:
		fmt.Println("Я немогу совершить выбранную операцию.") 

	}
	
	
}