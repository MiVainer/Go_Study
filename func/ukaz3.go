package main

import (
	"fmt"

)

// nums внутри функции — это копия оригинального среза. Изменив nums через append(), функция поменяла копию, а оригинал не изменился.
func appendByVal(nums []int, n int) []int {
    nums = append(nums, n)
	return nums
}

// Чтобы изменить срез в целом, можно использовать указатель:
func appendByPtr(smun *[]int, n int) {
    *smun = append(*smun, n)
}


func main() {
	
	nums := []int{42}
	smun := []int{52}
	appendByVal(nums, 43)
	nums = appendByVal(nums, 43)
	appendByPtr(&smun, 53)
	fmt.Println(nums, smun)
}