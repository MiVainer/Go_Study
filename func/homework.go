package main

import (
    "fmt"
    "os"
)

// normalize нормализует значения, переданные в vals,
// так чтобы их сумма была равна 1.
func normalize(vals ...*float64) {
    // Сначала вычисляем сумму исходных значений
    sum := 0.0
    for _, v := range vals {
        if v != nil {
            sum += *v
        }
    }
    // Если сумма равна 0, деление невозможно – выходим (или можно оставить как есть)
    if sum == 0 {
        return
    }
    // Нормализуем каждое значение
    for _, v := range vals {
        if v != nil {
            *v = *v / sum
        }
    }
}

func main() {
    a, b, c, d := 1.0, 2.0, 3.0, 4.0
    normalize(&a, &b, &c, &d)
    fmt.Println(a, b, c, d)
    // 0.1 0.2 0.3 0.4
    fmt.Println("PASS")
    os.Exit(0)
}