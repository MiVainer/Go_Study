package main

import (
	"fmt"

)

func main() {
	str := "ривет!П"
	//Преобразуем строку в срез байт
	bytes := []byte(str)
	fmt.Println(bytes)

	// Преобразование строки в юникод символы (руны)
	runes := []rune(str)
	fmt.Println(runes)

	// Строка реализована как массив байт, а не рун.
	fmt.Println(bytes[1])

	for i := 0; i < len(str); i++ {
		fmt.Printf("%d ", str[i]) // байты
	}

	for _, r := range str {
		fmt.Printf("%c ", r) // руны
	}

}