package main

import (
	"fmt"

)

func main() {
	var number int
	fmt.Scanf("%d", &number)

	// Посчитайте, сколько раз каждая цифра встречается
	// в числе `number`. Запишите результат в карту `counter`,
	// где ключом является цифра числа, а значением -
	// сколько раз она встречается
	counter := make(map[int]int)
	for number > 0 {
		digit := number % 10
		counter[digit]++
		fmt.Println(counter)
		fmt.Println(counter[digit])
		number = number / 10
	}

	fmt.Println(counter)
}