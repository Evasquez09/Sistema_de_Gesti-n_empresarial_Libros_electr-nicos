package servicios

import (
	"database/sql"
	"errors"
	"fmt"
	"sistema_gestion_libros/modelos"
	"time"
)

var (
	ErrPrestamoNoEncontrado = errors.New("no se encontró un préstamo activo para este libro")
	ErrLimitePrestamos      = errors.New("se alcanzó el límite máximo de préstamos para este estudiante")
)

type prestamoService struct {
	db *sql.DB
}

func (s *prestamoService) CrearPrestamo(libroID int, libroTitulo, estudiante string, limiteMaximo int) error {
	// Lógica de creación de préstamo
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM prestamos WHERE estudiante = ?", estudiante).Scan(&count)
	if err != nil {
		return err
	}
	if count >= limiteMaximo {
		return fmt.Errorf("El estudiante %s ha alcanzado el límite de préstamos", estudiante)
	}
	fechaPrestamo := time.Now()
	fechaDevolucion := fechaPrestamo.AddDate(0, 0, 30) // 30 días de préstamo
	_, err = s.db.Exec("INSERT INTO prestamos (libro_id, libro, estudiante, fecha_prestamo, fecha_devolucion) VALUES (?, ?, ?, ?, ?)",
		libroID, libroTitulo, estudiante, fechaPrestamo, fechaDevolucion)
	return err
}

func (s *prestamoService) RegistrarDevolucion(libroID int) error {
	res, err := s.db.Exec("DELETE FROM prestamos WHERE libro_id = ?", libroID)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("Préstamo no encontrado para el libro con ID %d", libroID)
	}
	return nil
}

func (s *prestamoService) ObtenerActivos() []modelos.Prestamo {
	rows, err := s.db.Query("SELECT libro_id, libro, estudiante, fecha_prestamo, fecha_devolucion FROM prestamos")
	if err != nil {
		return []modelos.Prestamo{}
	}
	defer rows.Close()
	var prestamos []modelos.Prestamo
	for rows.Next() {
		var p modelos.Prestamo
		err := rows.Scan(&p.LibroID, &p.Libro, &p.Estudiante, &p.FechaPrestamo, &p.FechaDevolucion)
		if err != nil {
			continue
		}
		prestamos = append(prestamos, p)
	}
	return prestamos
}

func NewPrestamoService(db *sql.DB) IPrestamoService {
	return &prestamoService{db: db}
}

func (s *prestamoService) VerHistorialPrestamos() {
	rows, err := s.db.Query("SELECT libro_id, libro, estudiante, fecha_prestamo, fecha_devolucion, enlace FROM prestamos")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()
	fmt.Println("\n--- Historial de Préstamos ---")
	for rows.Next() {
		var p modelos.Prestamo
		rows.Scan(&p.LibroID, &p.Libro, &p.Estudiante, &p.FechaPrestamo, &p.FechaDevolucion, &p.Enlace)
		fmt.Printf("ID Libro: %d, Libro: %s, Estudiante: %s, Fecha de Préstamo: %s, Fecha de Devolución: %s, Enlace: %s\n",
			p.LibroID, p.Libro, p.Estudiante, p.FechaPrestamo.Format("02-01-2006"), p.FechaDevolucion.Format("02-01-2006"), p.Enlace)
	}
}

func (s *prestamoService) VerPrestamos() {
	rows, err := s.db.Query("SELECT libro_id, libro, estudiante, fecha_prestamo, fecha_devolucion, enlace FROM prestamos")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()
	fmt.Println("\n--- Lista de Préstamos Activos ---")
	for rows.Next() {
		var p modelos.Prestamo
		rows.Scan(&p.LibroID, &p.Libro, &p.Estudiante, &p.FechaPrestamo, &p.FechaDevolucion, &p.Enlace)
		fmt.Printf("ID Libro: %d, Libro: %s, Estudiante: %s, Fecha de Préstamo: %s, Fecha de Devolución: %s, Enlace: %s\n",
			p.LibroID, p.Libro, p.Estudiante, p.FechaPrestamo.Format("02-01-2006"), p.FechaDevolucion.Format("02-01-2006"), p.Enlace)
	}
}

func (s *prestamoService) ObtenerHistorial() []modelos.Prestamo {
	return []modelos.Prestamo{}
}
