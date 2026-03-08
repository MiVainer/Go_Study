package main

import (
	"fmt"

)

func main() {
	m := map[string]string{"a": "apple", "b": "banana"}
	//итерация по записям
	for key, val := range m {
		fmt.Printf("%s -> %s\n", key, val)
	}
	fmt.Printf("----------------------------\n")
	//итерация по ключам
	for key := range m {
    	fmt.Println("key:", key)
	}
}