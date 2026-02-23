package main

import (
	"fmt"
	"time"
)

func main() {
	//Парсинг строки в длительность
	m, _ := time.ParseDuration("1h30m")
	a, _ := time.ParseDuration("300s")
	b, _ := time.ParseDuration("10m")

	//Перевод переменных в минуты
	fmt.Println("1,5 часа это", m.Minutes(), "минут")
	fmt.Println("300 секунд это", a.Minutes(), "минут!")
	fmt.Println(b.Minutes(), "Минут !=)")

}
