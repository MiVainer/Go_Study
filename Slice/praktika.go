package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	switch {
	case len(text) <= width:
		fmt.Println("Текст не изменился:", text)
	
	default:
		// Перевели строку в байты
		bytes := []byte(text)
		// взяли количество width байт и сохранили в переменную res
		res := bytes[:width]
		fmt.Println(string(res)+"...") // вывели укороченную строку с точками
	}

}