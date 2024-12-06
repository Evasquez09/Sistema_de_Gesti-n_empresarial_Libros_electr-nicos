package servicios

import (
	"errors"
	"fmt"
	"sistema_gestion_libros/modelos"
	"strings"
)

type categoriaService struct {
	categorias         []modelos.Categoria
	categoriaIDCounter int
}

// NewCategoriaService crea una nueva instancia del servicio de categorías
func NewCategoriaService() ICategoriaService {
	return &categoriaService{
		categorias:         []modelos.Categoria{},
		categoriaIDCounter: 1,
	}
}

// AgregarCategoria agrega una nueva categoría
func (s *categoriaService) AgregarCategoria(categoria modelos.Categoria) error {
	for _, c := range s.categorias {
		if c.Nombre == categoria.Nombre {
			return errors.New("la categoría ya existe en el sistema")
		}
	}
	categoria.ID = s.categoriaIDCounter
	s.categorias = append(s.categorias, categoria)
	s.categoriaIDCounter++
	fmt.Printf("Categoría agregada correctamente: %s (ID: %d)\n", categoria.Nombre, categoria.ID)
	return nil
}

// ExisteCategoria verifica si una categoría existe por nombre
func (s *categoriaService) ExisteCategoria(nombre string) bool {
	for _, categoria := range s.categorias {
		if categoria.Nombre == nombre {
			return true
		}
	}
	return false
}

// VerCategorias muestra todas las categorías disponibles
func (s *categoriaService) VerCategorias() {
	fmt.Println("\n--- Lista de Categorías ---")
	for _, categoria := range s.categorias {
		fmt.Printf("ID: %d, Nombre: %s\n", categoria.ID, categoria.Nombre)
	}
}

// ObtenerCategoriaPorID obtiene una categoría por su ID
func (s *categoriaService) ObtenerCategoriaPorID(id int) (modelos.Categoria, bool) {
	for _, categoria := range s.categorias {
		if categoria.ID == id {
			return categoria, true
		}
	}
	return modelos.Categoria{}, false
}

// BuscarCategorias busca categorías por nombre parcial
func (s *categoriaService) BuscarCategorias(query string) []modelos.Categoria {
	var resultados []modelos.Categoria
	for _, categoria := range s.categorias {
		if strings.Contains(strings.ToLower(categoria.Nombre), strings.ToLower(query)) {
			resultados = append(resultados, categoria)
		}
	}
	return resultados
}
