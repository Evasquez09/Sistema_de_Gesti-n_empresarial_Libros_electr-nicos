package utilidades

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// LeerEntrada lee una línea de la entrada estándar y la devuelve sin espacios adicionales
func LeerEntrada(mensaje string) string {
	fmt.Print(mensaje)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	entrada := scanner.Text()
	return strings.TrimSpace(entrada)
}
