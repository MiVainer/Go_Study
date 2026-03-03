package main

import (
	"fmt"
	"time"
)

func main() {
	switch day := time.Now().Weekday(); 
	day {
	case time.Saturday, time.Sunday:
		fmt.Println(day, "is a weekend")
	default:
		fmt.Println(day, "is a weekday")
	}
}