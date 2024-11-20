package servicios

import (
	"fmt"
	"sistema_gestion_libros/modelos"
	"strings"
)

var libros []modelos.Libro
var libroIDCounter int = 1

// Agrega un libro a la lista de libros
func AgregarLibro(libro modelos.Libro) {
	for _, l := range libros {
		if l.Titulo == libro.Titulo {
			fmt.Println("El libro ya existe en el sistema.")
			return
		}
	}
	libro.ID = libroIDCounter
	libros = append(libros, libro)
	libroIDCounter++
	fmt.Printf("Libro agregado correctamente: %s (ID: %d)\n", libro.Titulo, libro.ID)
}

// Verifica si un libro con el título dado ya existe
func ExisteLibro(titulo string) bool {
	for _, libro := range libros {
		if libro.Titulo == titulo {
			return true
		}
	}
	return false
}

// Obtiene un libro por su ID
func ObtenerLibroPorID(id int) (modelos.Libro, bool) {
	for _, libro := range libros {
		if libro.ID == id {
			return libro, true
		}
	}
	return modelos.Libro{}, false
}

// Muestra la lista de libros
func VerLibros() {
	fmt.Println("\n--- Lista de Libros ---")
	for _, libro := range libros {
		fmt.Printf("ID: %d, Título: %s, Autor: %s, Categoría: %s\n", libro.ID, libro.Titulo, libro.Autor, libro.Categoria)
	}
}

func BuscarLibros(query string) []modelos.Libro {
	var resultados []modelos.Libro
	for _, libro := range libros {
		if strings.Contains(strings.ToLower(libro.Titulo), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(libro.Autor), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(libro.Categoria), strings.ToLower(query)) {
			resultados = append(resultados, libro)
		}
	}
	return resultados
}
