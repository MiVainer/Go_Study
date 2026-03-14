package main

import (
	"errors"
	"fmt"
)

func divideAndRemainder(a, b int) (int, int, error) {
	if b == 0 {
    return 0, 0, errors.New("Деление на ноль!")
} else {
	chast := a / b
	ost := a % b
	return chast, ost, nil
}

}

func main() {
	var a,b int
	fmt.Println("Введи 2 целых числа")
	fmt.Scan(&a, &b)
	chast, ost, err := divideAndRemainder(a, b)
  	if err != nil {
        fmt.Println("Ошибка:", err)
    } else {
        fmt.Printf("%d / %d = %d, остаток %d\n", a, b, chast, ost)
    }
}



