package servicios

import (
	"fmt"
	"sistema_gestion_libros/modelos"
	"sistema_gestion_libros/utilidades"
	"time"
)

var prestamos []modelos.Prestamo

func CrearPrestamo(libro string, estudiante string) {
	enlace := utilidades.GenerarEnlaceUnico()
	prestamo := modelos.Prestamo{
		Libro:      libro,
		Estudiante: estudiante,
		Fecha:      time.Now(),
		Enlace:     enlace,
	}
	prestamos = append(prestamos, prestamo)
	fmt.Printf("Préstamo creado para el libro '%s'. Enlace de acceso: %s\n", libro, enlace)
}

func VerPrestamos() {
	fmt.Println("\n--- Lista de Préstamos Activos ---")
	for _, prestamo := range prestamos {
		fmt.Printf("Libro: %s, Estudiante: %s, Fecha: %s, Enlace: %s\n",
			prestamo.Libro, prestamo.Estudiante, prestamo.Fecha.Format("02-01-2006"), prestamo.Enlace)
	}
}
