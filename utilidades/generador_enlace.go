package utilidades

import (
	"crypto/rand"
	"encoding/hex"
)

func GenerarEnlaceUnico() string {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return "https://mibiblioteca.com/prestamo/" + hex.EncodeToString(bytes)
}
