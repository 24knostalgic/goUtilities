package go_utilities

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
)

var (
	secretKey = "some random secret key"
	message   = "some random message"
)

// Generate a salt (random encoded string) with 16 bytes of crypto/rand data
// Salts create unique passwords (here, messages) even in the instance of two users choosing the same passwords
// Thus, better to have
func generateSalt() string {
	randomByteArr := make([]byte, 16)
	_, err := rand.Read(randomByteArr)
	if err != nil {
		return ""
	}
	// encode the bytes using URLEncoding method of base64 encoding
	return base64.URLEncoding.EncodeToString(randomByteArr)
}

// GenerateHMAC256Secret generates a keyed-hash i.e., a hash (using hash function such as SHA-2) that uses a secret key to sign a message
// Since HMAC is not a cipher, it can't be decrypted
// The recipient/server takes all the needed input, computes the HMAC using the secure secret key, and checks if the result is equal to the value in say, database
// JWT also uses HMAC to encrypt JSON data/message
// Thus, a hacker can't make sense out of the HMAC encrypted passwords in database
func GenerateHMAC256Secret() string {
	hash := hmac.New(sha256.New, []byte(secretKey)) // create hash by specifying hashing algo and a secret
	salt := generateSalt()

	hash.Write([]byte(message + salt)) // sign message + salt with the 'secretKey'

	sha := hex.EncodeToString(hash.Sum(nil)) // encode the result as hexadecimal string
	return sha
}
