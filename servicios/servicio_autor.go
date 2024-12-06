package servicios

import (
	"errors"
	"fmt"
	"sistema_gestion_libros/modelos"
	"strings"
)

// Se encapsulan las variables en una estructura y se accede a ellas a través de métodos
type autorService struct {
	autores        []modelos.Autor
	autorIDCounter int
}

// NewAutorService crea una nueva instancia del servicio de autores
func NewAutorService() IAutorService {
	return &autorService{
		autores:        []modelos.Autor{},
		autorIDCounter: 1,
	}
}

// AgregarAutor agrega un autor a la lista de autores
func (s *autorService) AgregarAutor(autor modelos.Autor) error {
	for _, a := range s.autores {
		if a.Nombre == autor.Nombre {
			return errors.New("el autor ya existe en el sistema")
		}
	}
	autor.ID = s.autorIDCounter
	s.autores = append(s.autores, autor)
	s.autorIDCounter++
	fmt.Printf("Autor agregado correctamente: %s (ID: %d)\n", autor.Nombre, autor.ID)
	return nil
}

// ExisteAutor verifica si existe un autor con el nombre dado
func (s *autorService) ExisteAutor(nombre string) bool {
	for _, autor := range s.autores {
		if autor.Nombre == nombre {
			return true
		}
	}
	return false
}

// ObtenerAutorPorID obtiene un autor por su ID
func (s *autorService) ObtenerAutorPorID(id int) (modelos.Autor, bool) {
	for _, autor := range s.autores {
		if autor.ID == id {
			return autor, true
		}
	}
	return modelos.Autor{}, false
}

// VerAutores muestra la lista de autores
func (s *autorService) VerAutores() {
	fmt.Println("\n--- Lista de Autores ---")
	for _, autor := range s.autores {
		fmt.Printf("ID: %d, Nombre: %s\n", autor.ID, autor.Nombre)
	}
}

// BuscarAutores busca autores por nombre parcial
func (s *autorService) BuscarAutores(query string) []modelos.Autor {
	var resultados []modelos.Autor
	for _, autor := range s.autores {
		if strings.Contains(strings.ToLower(autor.Nombre), strings.ToLower(query)) {
			resultados = append(resultados, autor)
		}
	}
	return resultados
}
