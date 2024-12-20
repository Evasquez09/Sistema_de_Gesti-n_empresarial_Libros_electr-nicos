package servicios

import (
	"database/sql"
	"fmt"
	"log"
	"sistema_gestion_libros/modelos"
	"strings"
)

type autorService struct {
	db *sql.DB
}

func NewAutorService(db *sql.DB) IAutorService {
	return &autorService{db: db}
}

func (s *autorService) AgregarAutor(autor modelos.Autor) error {
	_, err := s.db.Exec("INSERT INTO autores (nombre) VALUES (?)", autor.Nombre)
	return err
}

func (s *autorService) ActualizarAutor(autor modelos.Autor) error {
	res, err := s.db.Exec("UPDATE autores SET nombre=? WHERE id=?", autor.Nombre, autor.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("autor con ID %d no encontrado", autor.ID)
	}
	return nil
}

func (s *autorService) EliminarAutor(id int) error {
	res, err := s.db.Exec("DELETE FROM autores WHERE id=?", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("autor con ID %d no encontrado", id)
	}
	return nil
}

func (s *autorService) ObtenerTodos() []modelos.Autor {
	rows, err := s.db.Query("SELECT id, nombre FROM autores")
	if err != nil {
		return []modelos.Autor{}
	}
	defer rows.Close()

	var autores []modelos.Autor
	for rows.Next() {
		var autor modelos.Autor
		rows.Scan(&autor.ID, &autor.Nombre)
		autores = append(autores, autor)
	}
	return autores
}

func (s *autorService) ExisteAutor(nombre string) bool {
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM autores WHERE nombre = ?", nombre).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (s *autorService) ObtenerAutorPorID(id int) (modelos.Autor, bool) {
	var autor modelos.Autor
	err := s.db.QueryRow("SELECT id, nombre FROM autores WHERE id = ?", id).Scan(&autor.ID, &autor.Nombre)
	if err != nil {
		return modelos.Autor{}, false
	}
	return autor, true
}

func (s *autorService) VerAutores() {
	rows, err := s.db.Query("SELECT id, nombre FROM autores")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var autor modelos.Autor
		rows.Scan(&autor.ID, &autor.Nombre)
		fmt.Printf("ID: %d, Nombre: %s\n", autor.ID, autor.Nombre)
	}
}

func (s *autorService) BuscarAutores(query string) []modelos.Autor {
	q := "%" + strings.ToLower(query) + "%"
	rows, err := s.db.Query("SELECT id, nombre FROM autores WHERE LOWER(nombre) LIKE ?", q)
	if err != nil {
		log.Println("Error en la búsqueda de autores:", err) // Log para depurar
		return []modelos.Autor{}
	}
	defer rows.Close()

	var autores []modelos.Autor
	for rows.Next() {
		var a modelos.Autor
		err := rows.Scan(&a.ID, &a.Nombre)
		if err != nil {
			log.Println("Error al escanear autor:", err) // Log para depurar
			continue
		}
		autores = append(autores, a)
	}
	return autores
}
