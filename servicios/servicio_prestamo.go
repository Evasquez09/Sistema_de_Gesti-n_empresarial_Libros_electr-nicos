package servicios

import (
	"errors"
	"fmt"
	"sistema_gestion_libros/modelos"
	"sistema_gestion_libros/utilidades"
	"time"
)

// Errores específicos para el manejo de préstamos
var (
	ErrPrestamoNoEncontrado = errors.New("no se encontró un préstamo activo para este libro")
	ErrLimitePrestamos      = errors.New("se alcanzó el límite máximo de préstamos para este estudiante")
)

type prestamoService struct {
	prestamos          []modelos.Prestamo
	historialPrestamos []modelos.Prestamo
}

// NewPrestamoService crea una nueva instancia del servicio de préstamos
func NewPrestamoService() IPrestamoService {
	return &prestamoService{
		prestamos:          []modelos.Prestamo{},
		historialPrestamos: []modelos.Prestamo{},
	}
}

// CrearPrestamo crea un préstamo para un estudiante y un libro determinado.
// Verifica el límite de préstamos permitido antes de crear el nuevo préstamo.
func (s *prestamoService) CrearPrestamo(libroID int, libroTitulo, estudiante string, limiteMaximo int) error {
	// Verificar si el estudiante excede el límite de préstamos
	conteo := s.contarPrestamosPorEstudiante(estudiante)
	if conteo >= limiteMaximo {
		return ErrLimitePrestamos
	}

	enlace := utilidades.GenerarEnlaceUnico()
	prestamo := modelos.Prestamo{
		LibroID:    libroID,
		Libro:      libroTitulo,
		Estudiante: estudiante,
		Fecha:      time.Now(),
		Enlace:     enlace,
	}
	s.prestamos = append(s.prestamos, prestamo)
	fmt.Printf("Préstamo creado para el libro '%s'. Enlace de acceso: %s\n", libroTitulo, enlace)
	return nil
}

// RegistrarDevolucion procesa la devolución de un libro. Si no se encuentra el préstamo activo se devuelve un error.
func (s *prestamoService) RegistrarDevolucion(libroID int) error {
	for i, prestamo := range s.prestamos {
		if prestamo.LibroID == libroID {
			s.historialPrestamos = append(s.historialPrestamos, prestamo)
			s.prestamos = append(s.prestamos[:i], s.prestamos[i+1:]...)
			fmt.Printf("Devolución registrada: Libro '%s' del estudiante '%s'.\n", prestamo.Libro, prestamo.Estudiante)
			return nil
		}
	}
	return ErrPrestamoNoEncontrado
}

// VerHistorialPrestamos muestra el historial de préstamos
func (s *prestamoService) VerHistorialPrestamos() {
	fmt.Println("\n--- Historial de Préstamos ---")
	for _, prestamo := range s.historialPrestamos {
		fmt.Printf("Libro: %s, Estudiante: %s, Fecha: %s\n",
			prestamo.Libro, prestamo.Estudiante, prestamo.Fecha.Format("02-01-2006"))
	}
}

// VerPrestamos muestra la lista de préstamos activos
func (s *prestamoService) VerPrestamos() {
	fmt.Println("\n--- Lista de Préstamos Activos ---")
	if len(s.prestamos) == 0 {
		fmt.Println("No hay préstamos activos en este momento.")
		return
	}
	for _, prestamo := range s.prestamos {
		fmt.Printf("ID del Libro: %d, Libro: %s, Estudiante: %s, Fecha: %s, Enlace: %s\n",
			prestamo.LibroID, prestamo.Libro, prestamo.Estudiante, prestamo.Fecha.Format("02-01-2006"), prestamo.Enlace)
	}
}

// contarPrestamosPorEstudiante retorna el número de préstamos activos para un estudiante dado.
// Esta función es interna del servicio y ejemplifica el concepto de encapsulación.
func (s *prestamoService) contarPrestamosPorEstudiante(estudiante string) int {
	conteo := 0
	for _, prestamo := range s.prestamos {
		if prestamo.Estudiante == estudiante {
			conteo++
		}
	}
	return conteo
}
