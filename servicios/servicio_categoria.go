package servicios

import (
	"database/sql"
	"errors"
	"log"
	"sistema_gestion_libros/modelos"
	"strings"
)

type categoriaService struct {
	db *sql.DB
}

func NewCategoriaService(db *sql.DB) ICategoriaService {
	return &categoriaService{db: db}
}

func (s *categoriaService) ObtenerTodas() []modelos.Categoria {
	rows, err := s.db.Query("SELECT id, nombre FROM categorias")
	if err != nil {
		log.Println("Error al obtener categorías de la base de datos:", err)
		return []modelos.Categoria{}
	}
	defer rows.Close()

	var categorias []modelos.Categoria
	for rows.Next() {
		var c modelos.Categoria
		err := rows.Scan(&c.ID, &c.Nombre)
		if err != nil {
			log.Println("Error al escanear categoría:", err)
			continue
		}
		log.Println("Categoría obtenida:", c) // Log para depurar
		categorias = append(categorias, c)
	}

	log.Println("Total de categorías obtenidas:", len(categorias)) // Log para depurar
	return categorias
}

func (s *categoriaService) AgregarCategoria(categoria modelos.Categoria) error {
	_, err := s.db.Exec("INSERT INTO categorias (nombre) VALUES (?)", categoria.Nombre)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New("la categoría ya existe en el sistema")
		}
		return err
	}
	return nil
}

func (s *categoriaService) ExisteCategoria(nombre string) bool {
	var count int
	err := s.db.QueryRow("SELECT COUNT(*) FROM categorias WHERE nombre = ?", nombre).Scan(&count)
	if err != nil {
		return false
	}
	return count > 0
}

func (s *categoriaService) VerCategorias() {
	rows, err := s.db.Query("SELECT id, nombre FROM categorias")
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var c modelos.Categoria
		rows.Scan(&c.ID, &c.Nombre)
		// Imprimir si se desea (este método ya no es tan crítico si usamos JSON)
		println(c.ID, c.Nombre)
	}
}

func (s *categoriaService) ObtenerCategoriaPorID(id int) (modelos.Categoria, bool) {
	var cat modelos.Categoria
	err := s.db.QueryRow("SELECT id, nombre FROM categorias WHERE id=?", id).Scan(&cat.ID, &cat.Nombre)
	if err != nil {
		return modelos.Categoria{}, false
	}
	return cat, true
}

func (s *categoriaService) BuscarCategorias(query string) []modelos.Categoria {
	q := "%" + strings.ToLower(query) + "%"
	rows, err := s.db.Query("SELECT id, nombre FROM categorias WHERE LOWER(nombre) LIKE ?", q)
	if err != nil {
		log.Println("Error en la búsqueda de categorías:", err) // Log para depurar
		return []modelos.Categoria{}
	}
	defer rows.Close()

	var categorias []modelos.Categoria
	for rows.Next() {
		var c modelos.Categoria
		err := rows.Scan(&c.ID, &c.Nombre)
		if err != nil {
			log.Println("Error al escanear categoría:", err) // Log para depurar
			continue
		}
		categorias = append(categorias, c)
	}
	return categorias
}

func (s *categoriaService) ActualizarCategoria(cat modelos.Categoria) error {
	res, err := s.db.Exec("UPDATE categorias SET nombre=? WHERE id=?", cat.Nombre, cat.ID)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("categoría no encontrada")
	}
	return nil
}

func (s *categoriaService) EliminarCategoria(id int) error {
	res, err := s.db.Exec("DELETE FROM categorias WHERE id=?", id)
	if err != nil {
		return err
	}
	rowsAffected, _ := res.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("categoría no encontrada")
	}
	return nil
}
