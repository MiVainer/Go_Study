package main

import (
    "fmt"


)


func main() {
	var n, fac int
	fmt.Println("Введи целое число:")
	fmt.Scan(&n)
	fac = 1
	for i := 1; i <= n; i++ {
		fac*=i
	}
	fmt.Println("Факториал числа:", fac)
}

