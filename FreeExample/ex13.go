package main

import (
	"fmt"
)


func average(nums ...float64) float64 {
    if len(nums) == 0 {
        return 0
    }
    sum := 0.0
    for _, n := range nums {
        sum += n
    }
    return sum / float64(len(nums))
}

func main() {
    fmt.Println(average(2, 6, 3, 5, 8.2))       // 2

}