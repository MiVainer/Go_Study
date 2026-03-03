package main

import (
	"fmt"

)

func main() {
	var arr [5]int
	fmt.Println("Первичный массив: ", arr)
	arr[2] = 10
	fmt.Println("Установили аргумент 2 =", arr[2])
	fmt.Println("Длина массива - ", len(arr), "символов" )
	
}