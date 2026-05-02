package main

import (
	"fmt"
)

// Если функция изменяет содержимое карты, передавайте ее как значение.
func changeMap(m map[string]int) {
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
}

// Если функция полностью заменяет саму карту,
// передавайте ее как значение и возвращайте новое значение.
func clearMap(m map[string]int) map[string]int {
	m = map[string]int{}
	return m
}

// Либо передавайте карту как указатель.
func clearMapPtr(m *map[string]int) {
	*m = map[string]int{}
}

func main() {
	m := map[string]int{}
	fmt.Println("Пустая карта:", m)
	m["four"] = 4
	fmt.Println("добавили ключ-значение:", m)
	changeMap(m)
	fmt.Println(m)

	m = clearMap(m)
	fmt.Println(m)

	clearMapPtr(&m)
	fmt.Println(m)

}
