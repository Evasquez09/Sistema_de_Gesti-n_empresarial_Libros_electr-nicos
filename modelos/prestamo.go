package modelos

import "time"

type Prestamo struct {
	LibroID    int
	Libro      string
	Estudiante string
	Fecha      time.Time
	Enlace     string
}
