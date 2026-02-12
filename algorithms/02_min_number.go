package main

import("fmt")

// Функция принимает слайс и ВОЗВРАЩАЕТ ДВА ИНТА (само число и его индекс)
func MinNumber(minNumber []int) (int, int){
    min := minNumber[0]   // Считаем первое число минимальным
    minIndex := 0         // Запоминаем, что его индекс — 0
    
    // i — индекс (позиция), val — само число
    for i, val := range minNumber {
        // Если текущее число меньше нашего "рекорда"
        if val < min {
            min = val     // Обновляем минимальное число
            minIndex = i  // Запоминаем его новую позицию (индекс)
        }
    }
    return min, minIndex  // Отдаем оба результата
}

func main(){
    list := []int{15, 2, 45, -5, 10}
    
    // Распаковываем результат в две переменные: val (число) и i (индекс)
    val, i := MinNumber(list)
    
    // %d — подставляет целое число в строку
    fmt.Printf("Min number: %d, Index: %d\n", val, i)
}
