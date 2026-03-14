package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "time"
)

// calculateAge возвращает полное количество лет на текущую дату
func calculateAge(birth, now time.Time) int {
    years := now.Year() - birth.Year()
    // Если день рождения в этом году ещё не наступил, вычитаем год
    if now.Month() < birth.Month() || (now.Month() == birth.Month() && now.Day() < birth.Day()) {
        years--
    }
    return years
}

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Введите дату рождения (ДД.ММ.ГГГГ): ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input) // удаляем символы новой строки и пробелы

    // Задаём ожидаемый формат (эталонная дата Go: 02.01.2006)
    layout := "02.01.2006"
    // Парсим введённую строку в UTC (предполагаем, что время не важно)
    birthDate, err := time.Parse(layout, input)
    if err != nil {
        fmt.Println("Ошибка: неверный формат даты. Используйте ДД.ММ.ГГГГ, например 15.04.1992")
        return
    }
    
	// Проверяем, что дата не в будущем и не слишком старая (опционально)
    now := time.Now()
    if birthDate.After(now) {
        fmt.Println("Дата рождения не может быть в будущем.")
        return
    }

	age := calculateAge(birthDate, now)
    fmt.Printf("Ваш возраст: %d лет\n", age)


}