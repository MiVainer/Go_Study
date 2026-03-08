package main

import (
	"fmt"

)

func main() {
	// Обход строки по символам
	for idx, char := range "ого" {
		fmt.Println(idx, char, string(char))
	}
	// Обход строки по байтам, каждой русской букве принадлежит 2 байта
	str := "ого"
	fmt.Println("Количество байт в слове:", len(str))
	for i := 0; i < len(str); i++ {
    	fmt.Printf("str[%d] = %d\n", i, str[i])
	}
}