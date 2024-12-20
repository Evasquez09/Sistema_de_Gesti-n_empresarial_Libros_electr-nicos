package utilidades

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LeerEntrada(mensaje string) string {
	fmt.Print(mensaje)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	entrada := scanner.Text()
	return strings.TrimSpace(entrada)
}
