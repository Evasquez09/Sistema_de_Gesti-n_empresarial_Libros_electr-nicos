package servicios

import (
	"errors"
	"fmt"
	"sistema_gestion_libros/modelos"
	"strings"
)

type libroService struct {
	libros         []modelos.Libro
	libroIDCounter int
}

// NewLibroService crea una instancia del servicio de libros
func NewLibroService() ILibroService {
	return &libroService{
		libros:         []modelos.Libro{},
		libroIDCounter: 1,
	}
}

// AgregarLibro agrega un libro a la lista
func (s *libroService) AgregarLibro(libro modelos.Libro) error {
	for _, l := range s.libros {
		if l.Titulo == libro.Titulo {
			return errors.New("el libro ya existe en el sistema")
		}
	}
	libro.ID = s.libroIDCounter
	s.libros = append(s.libros, libro)
	s.libroIDCounter++
	fmt.Printf("Libro agregado correctamente: %s (ID: %d)\n", libro.Titulo, libro.ID)
	return nil
}

// ExisteLibro verifica si un libro existe por título
func (s *libroService) ExisteLibro(titulo string) bool {
	for _, libro := range s.libros {
		if libro.Titulo == titulo {
			return true
		}
	}
	return false
}

// ObtenerLibroPorID obtiene un libro por su ID
func (s *libroService) ObtenerLibroPorID(id int) (modelos.Libro, bool) {
	for _, libro := range s.libros {
		if libro.ID == id {
			return libro, true
		}
	}
	return modelos.Libro{}, false
}

// VerLibros muestra la lista completa de libros
func (s *libroService) VerLibros() {
	fmt.Println("\n--- Lista de Libros ---")
	for _, libro := range s.libros {
		fmt.Printf("ID: %d, Título: %s, Autor: %s, Categoría: %s\n", libro.ID, libro.Titulo, libro.Autor, libro.Categoria)
	}
}

// BuscarLibros busca libros por título, autor o categoría
func (s *libroService) BuscarLibros(query string) []modelos.Libro {
	var resultados []modelos.Libro
	for _, libro := range s.libros {
		if strings.Contains(strings.ToLower(libro.Titulo), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(libro.Autor), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(libro.Categoria), strings.ToLower(query)) {
			resultados = append(resultados, libro)
		}
	}
	return resultados
}
