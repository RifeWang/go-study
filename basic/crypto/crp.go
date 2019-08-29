package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
	"io"
)

var (
	hmacSecret = []byte("secret") // 密钥
)

func getHmac(s string) string {
	h := hmac.New(sha256.New, hmacSecret)
	io.WriteString(h, s)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func main() {
	c := getHmac("username")
	fmt.Println(c)
}
