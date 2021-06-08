package urls

import (
	"math/rand"
	"time"
)

const chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateSlug(length int) string {
	rand.Seed(time.Now().UnixNano())
	slug := make([]rune, length)
	for i := 0; i < length; i++ {
		slug[i] = rune(chars[rand.Intn(len(chars))])
	}
	return string(slug)
}