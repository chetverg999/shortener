package shortener

import (
	"math/rand"
	"strings"
	"time"
)

// Функция для генерации случайной строки заданной длины из списка символов
func Shortener() string {
	rand.Seed(time.Now().UnixNano()) // Зерно генератора, основанное на текущем времени в наносекундах
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "abcdefghijklmnopqrstuvwxyz" + "0123456789")
	length := 2 // Длина возвращаемой строки
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))]) // Выписывает случайный символ из набора
	}
	return b.String()
}
