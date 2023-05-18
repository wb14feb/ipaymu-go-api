package ipaymu_go_api

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

func GenerateSignature(body string, method string, cfg Client) string {
	bodyHash := sha256.Sum256([]byte(body))
	bodyHashToString := hex.EncodeToString(bodyHash[:])
	stringToSign := "POST:" + cfg.VirtualAccount + ":" + strings.ToLower(bodyHashToString) + ":" + cfg.ApiKey

	h := hmac.New(sha256.New, []byte(cfg.ApiKey))
	h.Write([]byte(stringToSign))
	return hex.EncodeToString(h.Sum(nil))
}
