package servicios

import (
	"database/sql"
	"fmt"
	"log"
	"sistema_gestion_libros/modelos"
	"strings"
)

type libroService struct {
	db *sql.DB
}

func NewLibroService(db *sql.DB) ILibroService {
	return &libroService{db: db}
}

func (s *libroService) AgregarLibro(libro modelos.Libro) error {
	_, err := s.db.Exec("INSERT INTO libros (titulo, autor, categoria) VALUES (?, ?, ?)", libro.Titulo, libro.Autor, libro.Categoria)
	return err
}

func (s *libroService) ActualizarLibro(libro modelos.Libro) error {
	res, err := s.db.Exec("UPDATE libros SET titulo=?, autor=?, categoria=? WHERE id=?", libro.Titulo, libro.Autor, libro.Categoria, libro.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("libro con ID %d no encontrado", libro.ID)
	}
	return nil
}

func (s *libroService) EliminarLibro(id int) error {
	res, err := s.db.Exec("DELETE FROM libros WHERE id=?", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("libro con ID %d no encontrado", id)
	}
	return nil
}

func (s *libroService) ExisteLibro(titulo string) bool {
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM libros WHERE titulo = ?", titulo).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (s *libroService) ObtenerLibroPorID(id int) (modelos.Libro, bool) {
	var l modelos.Libro
	err := s.db.QueryRow("SELECT id, titulo, autor, categoria FROM libros WHERE id=?", id).Scan(&l.ID, &l.Titulo, &l.Autor, &l.Categoria)
	if err != nil {
		return modelos.Libro{}, false
	}
	return l, true
}

func (s *libroService) VerLibros() {
	rows, err := s.db.Query("SELECT id, titulo, autor, categoria FROM libros")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var l modelos.Libro
		rows.Scan(&l.ID, &l.Titulo, &l.Autor, &l.Categoria)
		// imprimir si se desea
	}
}

func (s *libroService) BuscarLibros(query string) []modelos.Libro {
	q := "%" + strings.ToLower(query) + "%"
	rows, err := s.db.Query("SELECT id, titulo, autor, categoria FROM libros WHERE LOWER(titulo) LIKE ? OR LOWER(autor) LIKE ? OR LOWER(categoria) LIKE ?", q, q, q)
	if err != nil {
		log.Println("Error en la búsqueda de libros:", err) // Log para depurar
		return []modelos.Libro{}
	}
	defer rows.Close()

	var libros []modelos.Libro
	for rows.Next() {
		var l modelos.Libro
		err := rows.Scan(&l.ID, &l.Titulo, &l.Autor, &l.Categoria)
		if err != nil {
			log.Println("Error al escanear libro:", err) // Log para depurar
			continue
		}
		libros = append(libros, l)
	}
	return libros
}

func (s *libroService) ObtenerTodos() []modelos.Libro {
	rows, err := s.db.Query("SELECT id, titulo, autor, categoria FROM libros")
	if err != nil {
		return []modelos.Libro{}
	}
	defer rows.Close()
	var libros []modelos.Libro
	for rows.Next() {
		var l modelos.Libro
		rows.Scan(&l.ID, &l.Titulo, &l.Autor, &l.Categoria)
		libros = append(libros, l)
	}
	return libros
}
