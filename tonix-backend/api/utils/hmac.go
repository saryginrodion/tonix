package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func HmacSha256(v string, secret string) string {
	hmac256 := hmac.New(sha256.New, []byte(secret))

	hmac256.Write([]byte(v))
	dataHmac := hmac256.Sum(nil)

	hmacHex := hex.EncodeToString(dataHmac)

	return hmacHex
}
