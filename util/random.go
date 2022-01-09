package util

import (
	"math/rand"
	"net/mail"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
	rand.Seed(time.Now().UnixNano())
}

// RandomInt generates a random integer between min and max
func RandomInt(min, max int64) int64 {
	return min + rand.Int63n(max - min + 1)
}

// RandomString generates a random string of length n
func RandomString(n int) string {
	var sb strings.Builder
	k := len(alphabet)

	for i := 0; i < n; i++ {
		c := alphabet[rand.Intn(k)]
		sb.WriteByte(c)
	}

	return sb.String()
}

// RandomFullName generates a random full name
func RandomFullName() string {
	return RandomString(6) + " " + RandomString(5)
}

// RandomEmail generates a random email
func RandomEmail() string {
	randEmail := RandomString(8)+"@"+RandomString(4)+".com"
	if IsValidEmail(randEmail) {
		return randEmail
	}
	return "bad-email"
}

// IsValidEmail checks to see if email is valid format
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}