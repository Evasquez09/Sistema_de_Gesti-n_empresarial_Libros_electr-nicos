package servicios

import (
	"fmt"
	"sistema_gestion_libros/modelos"
)

var autores []modelos.Autor
var autorIDCounter int = 1 // Contador para asignar IDs únicos

// Agrega un autor a la lista de autores
func AgregarAutor(autor modelos.Autor) {
	for _, a := range autores {
		if a.Nombre == autor.Nombre {
			fmt.Println("El autor ya existe en el sistema.")
			return
		}
	}
	autor.ID = autorIDCounter
	autores = append(autores, autor)
	autorIDCounter++
	fmt.Printf("Autor agregado correctamente: %s (ID: %d)\n", autor.Nombre, autor.ID)
}

// Verifica si un autor con el nombre dado ya existe
func ExisteAutor(nombre string) bool {
	for _, autor := range autores {
		if autor.Nombre == nombre {
			return true
		}
	}
	return false
}

// Obtiene un autor por su ID
func ObtenerAutorPorID(id int) (modelos.Autor, bool) {
	for _, autor := range autores {
		if autor.ID == id {
			return autor, true
		}
	}
	return modelos.Autor{}, false
}

// Muestra la lista de autores
func VerAutores() {
	fmt.Println("\n--- Lista de Autores ---")
	for _, autor := range autores {
		fmt.Printf("ID: %d, Nombre: %s\n", autor.ID, autor.Nombre)
	}
}
