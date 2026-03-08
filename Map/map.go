package main

import (
	"fmt"
)

func main() {
	// инициализировали карту
	m := make(map[string]int)
	//добавили в карту пары ключ-значение и вывели карту
	m["key"] = 7
	m["other"] = 13
	fmt.Println("map:", m)
	// вывести значение по ключу
	val := m["key"]
	fmt.Println("val:", val)
	// Вывести количество записей (пар ключ-значение)
	fmt.Println("len:", len(m))
	//удалить значение по ключу
	delete(m, "other")
	fmt.Println("map:", m)
}
