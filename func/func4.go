package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func predicate(n int) bool {
    return n%2 == 0
}

func filter(predicate func(int) bool, iterable []int) []int {
	// отфильтруйте `iterable` с помощью `predicate`
	// и верните отфильтрованный срез
	result := []int{}
	for _, v := range iterable {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

func main() {

	src := readInput()
	res := filter(predicate, src)
	// отфильтруйте `src` так, чтобы остались только четные числа
	// и запишите результат в `res`
	// res := filter(...)
	fmt.Println(res)
}

// ┌─────────────────────────────────┐
// │ не меняйте код ниже этой строки │
// └─────────────────────────────────┘

// readInput считывает целые числа из `os.Stdin`
// и возвращает в виде среза
// разделителем чисел считается пробел
func readInput() []int {
	var nums []int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		nums = append(nums, num)
	}
	return nums
}