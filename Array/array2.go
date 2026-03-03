package main

import (
	"fmt"

)

func main() {
	arr:=[...]int{1, 0, 2, 3, 1}
	arr1:=[...]string{"app", "sps"}
	fmt.Println("Первичные массивы: ", arr, arr1)
	arr1[1] = "sss"
	fmt.Println("Установили аргумент 2 строкового массива =", arr1[1])
	fmt.Println("Длина 1 массива - ", len(arr), "символов" )
	fmt.Println("Длина 2 массива - ", len(arr1), "символа" )
	fmt.Println("Массивы приняли следующий вид: ", arr, arr1)
	
}