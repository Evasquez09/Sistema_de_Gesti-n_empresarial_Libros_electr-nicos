package utilidades

import (
	"crypto/rand"
	"encoding/hex"
)

// GenerarEnlaceUnico genera un enlace único para la identificación de un préstamo
func GenerarEnlaceUnico() string {
	bytes := make([]byte, 8)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}
	return "https://mibiblioteca.com/prestamo/" + hex.EncodeToString(bytes)
}
