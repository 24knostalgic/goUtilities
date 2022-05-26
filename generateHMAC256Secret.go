package go_utilities

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHMAC256Secret() string {
	secret := "some random secret"
	hmac := hmac.New(sha256.New, []byte(secret)) // create new HMAC by specifying hashing algo and a secret

	data := "some random data"
	hmac.Write([]byte(data)) // write data into HMAC

	sha := hex.EncodeToString(hmac.Sum(nil)) // encode the result as hexadecimal string
	return sha
}
