package utils

import (
"crypto/sha256"
"encoding/hex"
"fmt"
"strings"
)

const (
	name = "booking-service"
)

// GetRedisKey get redis key
func GetRedisKey(value interface{}, values ...interface{}) string {
	k := make([]string, len(values)+2)
	k[0] = name
	k[1] = fmt.Sprint(value)
	for i, v := range values {
		k[i+2] = fmt.Sprint(v)
	}
	return strings.Join(k, "_")
}

// HashUserPassword hash user's pssword
func HashUserPassword(salt, password string) (string, error) {
	h := sha256.New()
	_, err := h.Write([]byte(password))
	if err != nil {
		return "", err
	}
	_, err = h.Write([]byte(salt))
	if err != nil {
		return "", err
	}
	digest := h.Sum(nil)
	hashedPassword := hex.EncodeToString(digest)
	return hashedPassword, nil
}
