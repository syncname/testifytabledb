package util

import (
	"math/rand"
	"strings"
	"time"
)

const alpha = "abcdefghijklmnopqrstuvwxyz"

func init() {
	//rand.New(rand.NewSource(time.Now().UnixNano()))
	rand.Seed(time.Now().UnixNano())
}

func RandomString(n int) string {
	var sb strings.Builder
	k := len(alpha)

	for i := 0; i < n; i++ {
		c := alpha[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

func RandomName() string {
	return RandomString(12)
}

func RandomMail() string {
	var sb strings.Builder
	k := len(alpha)

	for i := 0; i < 10; i++ {
		c := alpha[rand.Intn(k)]
		sb.WriteByte(c)
	}

	sb.WriteString("@todo.com")

	return sb.String()
}
