package go_utilities

import "golang.org/x/crypto/bcrypt"

func GenerateSecret() string {
	text := "some random password"
	hashedText, _ := bcrypt.GenerateFromPassword([]byte(text), 12)
	return string(hashedText)
}
