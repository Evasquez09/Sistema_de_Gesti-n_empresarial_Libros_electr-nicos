package servicios

import (
	"fmt"
	"sistema_gestion_libros/modelos"
	"strings"
)

var categorias []modelos.Categoria
var categoriaIDCounter int = 1 // Contador para asignar IDs únicos

// Agrega una nueva categoría
func AgregarCategoria(categoria modelos.Categoria) {
	for _, c := range categorias {
		if c.Nombre == categoria.Nombre {
			fmt.Println("La categoría ya existe en el sistema.")
			return
		}
	}
	categoria.ID = categoriaIDCounter
	categorias = append(categorias, categoria)
	categoriaIDCounter++
	fmt.Printf("Categoría agregada correctamente: %s (ID: %d)\n", categoria.Nombre, categoria.ID)
}

// Verifica si una categoría con el nombre dado existe
func ExisteCategoria(nombre string) bool {
	for _, categoria := range categorias {
		if categoria.Nombre == nombre {
			return true
		}
	}
	return false
}

// Muestra todas las categorías disponibles
func VerCategorias() {
	fmt.Println("\n--- Lista de Categorías ---")
	for _, categoria := range categorias {
		fmt.Printf("ID: %d, Nombre: %s\n", categoria.ID, categoria.Nombre)
	}
}

// Obtiene una categoría por su ID
func ObtenerCategoriaPorID(id int) (modelos.Categoria, bool) {
	for _, categoria := range categorias {
		if categoria.ID == id {
			return categoria, true
		}
	}
	return modelos.Categoria{}, false
}

func BuscarCategorias(query string) []modelos.Categoria {
	var resultados []modelos.Categoria
	for _, categoria := range categorias {
		if strings.Contains(strings.ToLower(categoria.Nombre), strings.ToLower(query)) {
			resultados = append(resultados, categoria)
		}
	}
	return resultados
}
