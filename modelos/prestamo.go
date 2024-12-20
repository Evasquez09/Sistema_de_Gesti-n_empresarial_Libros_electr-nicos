package modelos

import "time"

type Prestamo struct {
	LibroID         int       // ID único del libro
	Libro           string    // Título del libro prestado
	Estudiante      string    // Nombre del estudiante que pidió el libro prestado
	FechaPrestamo   time.Time // Fecha del préstamo
	FechaDevolucion time.Time // Fecha estimada para la devolución
	Enlace          string    // Enlace único para el préstamo
}
