package util

import (
	"math/rand"
	"strings"
	"time"
)

func GetRandomString() string {
	alphabet := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	n := 20

	alphabetSize := len(alphabet)
	var sb strings.Builder

	for i := 0; i < n; i++ {
		ch := alphabet[rand.Intn(alphabetSize)]
		sb.WriteRune(ch)
	}

	s := sb.String()
	return s
}

func GetRandomIntn(n int) int {
	rand.Seed(time.Now().UnixNano())

	return rand.Intn(n)
}
