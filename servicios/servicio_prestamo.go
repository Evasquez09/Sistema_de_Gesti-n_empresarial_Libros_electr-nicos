package servicios

import (
	"fmt"
	"sistema_gestion_libros/modelos"
	"sistema_gestion_libros/utilidades"
	"time"
)

var prestamos []modelos.Prestamo
var historialPrestamos []modelos.Prestamo

func CrearPrestamo(libroID int, libroTitulo, estudiante string) {
	enlace := utilidades.GenerarEnlaceUnico()
	prestamos = append(prestamos, modelos.Prestamo{
		LibroID:    libroID,
		Libro:      libroTitulo,
		Estudiante: estudiante,
		Fecha:      time.Now(),
		Enlace:     enlace,
	})
	fmt.Printf("Préstamo creado: %s para %s. Enlace: %s\n", libroTitulo, estudiante, enlace)
}

func VerPrestamos() {
	if len(prestamos) == 0 {
		fmt.Println("No hay préstamos activos.")
		return
	}
	fmt.Println("--- Préstamos Activos ---")
	for _, p := range prestamos {
		fmt.Printf("ID Libro: %d, Título: %s, Estudiante: %s, Fecha: %s\n", p.LibroID, p.Libro, p.Estudiante, p.Fecha.Format("02-01-2006"))
	}
}

func RegistrarDevolucion(libroID int) bool {
	for i, prestamo := range prestamos {
		if prestamo.LibroID == libroID {
			historialPrestamos = append(historialPrestamos, prestamo)
			prestamos = append(prestamos[:i], prestamos[i+1:]...)
			return true
		}
	}
	return false
}

func VerHistorialPrestamos() {
	if len(historialPrestamos) == 0 {
		fmt.Println("No hay historial de préstamos.")
		return
	}
	fmt.Println("--- Historial de Préstamos ---")
	for _, p := range historialPrestamos {
		fmt.Printf("ID Libro: %d, Título: %s, Estudiante: %s, Fecha: %s\n", p.LibroID, p.Libro, p.Estudiante, p.Fecha.Format("02-01-2006"))
	}
}
