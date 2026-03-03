package main

import (
	"fmt"

)

func main() {
	s := make([]string, 3)
	fmt.Printf("empty: %#v\n", s)
	s[0] = "ara"
	s[1] = "papa"
	s[2] = "mama"
	fmt.Println("Новый срез: ", s)
	fmt.Println("Длина среза: ", len(s))
	s = append(s, "dobavil")
	fmt.Println("Измененный срез:", s)
	fmt.Println("___________________________________")

	s1 := []string{"a", "b", "c"}
	fmt.Println("Срез который инициализировали при объявлении: ", s1)

}