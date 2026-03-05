package main

import (
	"fmt"

)

func main() {
	str := "го!"
	//Преобразуем строку в срез байт
	bytes := []byte(str)
	fmt.Println(bytes)

	// Преобразование строки в юникод символы (руны)
	runes := []rune(str)
	fmt.Println(runes)

	// Строка реализована как массив байт, а не рун.
	fmt.Println(bytes[1])
}