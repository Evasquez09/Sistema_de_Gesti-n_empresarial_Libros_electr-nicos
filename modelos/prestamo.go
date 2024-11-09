package modelos

import "time"

type Prestamo struct {
	Libro      string    // Título del libro prestado
	Estudiante string    // Nombre del estudiante que pidió el libro prestado
	Fecha      time.Time // Fecha del préstamo
	Enlace     string    // Enlace único para el préstamo
}
