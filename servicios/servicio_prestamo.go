package servicios

import (
	"fmt"
	"sistema_gestion_libros/modelos"
	"sistema_gestion_libros/utilidades"
	"time"
)

var prestamos []modelos.Prestamo
var historialPrestamos []modelos.Prestamo

// Crear un préstamo
func CrearPrestamo(libroID int, libroTitulo, estudiante string, limiteMaximo int) {
	// Verificar si el estudiante ya excede el límite de préstamos
	conteoPrestamos := contarPrestamosPorEstudiante(estudiante)
	if conteoPrestamos >= limiteMaximo {
		fmt.Printf("El estudiante '%s' ya alcanzó el límite máximo de préstamos (%d).\n", estudiante, limiteMaximo)
		return
	}

	enlace := utilidades.GenerarEnlaceUnico()
	prestamo := modelos.Prestamo{
		LibroID:    libroID,
		Libro:      libroTitulo,
		Estudiante: estudiante,
		Fecha:      time.Now(),
		Enlace:     enlace,
	}
	prestamos = append(prestamos, prestamo)
	fmt.Printf("Préstamo creado para el libro '%s'. Enlace de acceso: %s\n", libroTitulo, enlace)
}

// Registrar devolución de un libro
func RegistrarDevolucion(libroID int) bool {
	for i, prestamo := range prestamos {
		if prestamo.LibroID == libroID {
			historialPrestamos = append(historialPrestamos, prestamo)
			prestamos = append(prestamos[:i], prestamos[i+1:]...)
			fmt.Printf("Devolución registrada: Libro '%s' del estudiante '%s'.\n", prestamo.Libro, prestamo.Estudiante)
			return true
		}
	}
	return false
}

// Ver historial de préstamos
func VerHistorialPrestamos() {
	fmt.Println("\n--- Historial de Préstamos ---")
	for _, prestamo := range historialPrestamos {
		fmt.Printf("Libro: %s, Estudiante: %s, Fecha: %s\n",
			prestamo.Libro, prestamo.Estudiante, prestamo.Fecha.Format("02-01-2006"))
	}
}

// Contar los préstamos activos de un estudiante
func contarPrestamosPorEstudiante(estudiante string) int {
	conteo := 0
	for _, prestamo := range prestamos {
		if prestamo.Estudiante == estudiante {
			conteo++
		}
	}
	return conteo
}

// Función para ver los préstamos activos
func VerPrestamos() {
	fmt.Println("\n--- Lista de Préstamos Activos ---")
	if len(prestamos) == 0 {
		fmt.Println("No hay préstamos activos en este momento.")
		return
	}
	for _, prestamo := range prestamos {
		fmt.Printf("ID del Libro: %d, Libro: %s, Estudiante: %s, Fecha: %s, Enlace: %s\n",
			prestamo.LibroID, prestamo.Libro, prestamo.Estudiante, prestamo.Fecha.Format("02-01-2006"), prestamo.Enlace)
	}
}
