package main

import (
    "fmt"
	"math"


)


func main() {
	var a, b, c float64
	fmt.Println("Введи коэффициенты А,B,C через пробел:")
	fmt.Scan(&a, &b, &c)
	d := math.Pow(b, 2) - 4*a*c
	switch  {
	case d >= 0:
		x1 := (-b + math.Sqrt(d)) / 2*a
		x2 := (-b - math.Sqrt(d)) / 2*a
		fmt.Println("Х1=", x1, "X2=", x2)
	default:
		fmt.Println("Действительных корней нет") 

	}
	
}