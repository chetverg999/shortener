package shortener

import (
	"math/rand"
	"strings"
	"time"
)

// Shortener Функция для генерации случайной строки заданной длины из списка символов

func Shortener(length int) string {
	r := rand.New(rand.NewSource(time.Now().UnixNano())) // Зерно генератора, основанное на текущем времени в наносекундах
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[r.Intn(len(chars))]) // Выписывает случайный символ из набора
	}

	return b.String()
}
