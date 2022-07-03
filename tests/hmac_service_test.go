package tests

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	hmacservice "github.com/sonyamoonglade/storage-service/pkg/hmac"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidSignature(t *testing.T) {

	mockHeadersWithoutHmac := []string{"png", "static/images", "5"}

	actualHexSignature := hmacservice.GenerateHexSignature(mockHeadersWithoutHmac)

	h := sha256.New()
	for _, header := range mockHeadersWithoutHmac {
		h.Write([]byte(header))
	}

	VALID_HEX_STRING := hex.EncodeToString(h.Sum(nil))

	assert.Equal(t, actualHexSignature, VALID_HEX_STRING)

}
func TestValidHMAC(t *testing.T) {

	mockHeadersWithoutHmac := []string{"png", "static/images", "5"}
	mockSecretKey := "hello_key"

	signature := hmacservice.GenerateHexSignature(mockHeadersWithoutHmac)

	actualHmac := hmacservice.GenerateHMAC(signature, mockSecretKey)

	h := hmac.New(sha256.New, []byte(mockSecretKey))
	h.Write([]byte(signature))
	expectedHexHmac := hex.EncodeToString(h.Sum(nil))

	assert.Equal(t, expectedHexHmac, actualHmac)

}
