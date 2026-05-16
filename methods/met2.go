package main

import (
	"fmt"
	"math"

)

type circle struct {
    radius int
}

func (c *circle) area() float64 {
    return math.Pi * math.Pow(float64(c.radius), 2)
}


func main() {

	cptr := &circle{radius: 5}
	fmt.Println("circle area:", cptr.area())
}