package main

import (
	"math"
    "fmt"


)


func main() {
	var r, S float64
	fmt.Scan(&r)
	S = math.Pow(r, 2) * math.Pi
	fmt.Println("Площадь круга =", S)
}