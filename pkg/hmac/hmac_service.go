package hmacservice

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHexSignature(v []string) string {
	h := sha256.New()
	var str []byte
	for _, s := range v {
		str = append(str, []byte(s)...)
	}
	h.Write(str)
	signature := h.Sum(nil)
	return hex.EncodeToString(signature)
}

func GenerateHMAC(v string, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(v))
	resultHmac := h.Sum(nil)
	return hex.EncodeToString(resultHmac)
}

func ValidateHMAC(mac string, v string, key string) bool {
	expected := GenerateHMAC(v, key)
	m1 := []byte(mac)
	m2 := []byte(expected)
	compResult := hmac.Equal(m1, m2)
	return compResult
}
