package utils

import (
	"fmt"
	"time"
)

// Вывод времени работы программы
func OutputTimer() {
	startTime := time.Now() // Фиксируем время запуска программы
	for {
		hours, minutes, seconds := calculateElapsedTime(startTime)
		fmt.Printf("\nВремя с момента запуска программы: %02d:%02d:%02d\n", hours, minutes, seconds)
		time.Sleep(1 * time.Second)
	}
}

// Вычисляем время работы программы
func calculateElapsedTime(startTime time.Time) (int, int, int) {
	elapsed := time.Since(startTime)
	hours := int(elapsed.Hours())
	minutes := int(elapsed.Minutes()) % 60
	seconds := int(elapsed.Seconds()) % 60
	return hours, minutes, seconds
}
