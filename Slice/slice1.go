package main

import (
	"fmt"

)

func main() {
	src := []string{"a", "b", "c", "d", "e", "f"}
	//Создали строковый срез dst длиной как срез src
	dst := make([]string, len(src))
	fmt.Println("Срез src: ", src)
	fmt.Println("Срез dst: ", dst)
	//копируем в срез dst элементы среза src
	copy(dst, src)
	fmt.Println("Результат после копирования:", dst)

	//Присваиваем срезу sl1 часть среза dst
	sl1 := dst[2:5]
	fmt.Println("sl1:", sl1)
	// Срез sl2 включает элементы с 1 по последний среза dst 
	sl2 := dst[1:]
	fmt.Println("sl2:", sl2)
}