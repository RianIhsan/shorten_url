package random

import (
	"math/rand"
	"time"
)

func Generate() (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	rand.Seed(time.Now().UnixNano())

	shortURL := make([]byte, 8)
	for i := range shortURL {
		shortURL[i] = charset[rand.Intn(len(charset))]
	}

	return string(shortURL), nil
}
