package hash

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
)

func Hash(data interface{}) (string, error) {
	marshaled, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	hashed := sha256.Sum256(marshaled)
	hexed := hex.EncodeToString(hashed[:])
	return string(hexed), nil
}
