package main

import (
    "fmt"


)


func main() {
	var a, b int
	fmt.Println("Введи 2 целых числа через пробел:")
	fmt.Scan(&a, &b)
	switch  {
	case a > b:
		fmt.Println("А больше чем Б, а =",a)
	case b > a:
		fmt.Println("Б больше чем А, Б =",b)
	default:
		fmt.Println("Числа равны") 

	}
	
}