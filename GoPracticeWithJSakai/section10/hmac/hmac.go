package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}

func Server(apiKey, sign string, data []byte) {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	fmt.Println(expectedHMAC == sign)
}

func main() {
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"

	// hmacは認証に使われる
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write([]byte("data"))
	sign := hex.EncodeToString(h.Sum(nil))

	fmt.Println(sign)

	Server(apiKey, sign, []byte("data"))
}
