package go_utilities

import "golang.org/x/crypto/bcrypt"

func GenerateHashedPassword() string {
	password := "some random password"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(hashedPassword)
}
