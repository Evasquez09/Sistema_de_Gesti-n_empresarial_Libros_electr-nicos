package servicios

import "sistema_gestion_libros/modelos"

// Definición de interfaces para encapsular la lógica del sistema.
// De esta manera, la implementación puede cambiar sin afectar al main ni a otras capas.

type IAutorService interface {
	AgregarAutor(autor modelos.Autor) error
	ExisteAutor(nombre string) bool
	ObtenerAutorPorID(id int) (modelos.Autor, bool)
	VerAutores()
	BuscarAutores(query string) []modelos.Autor
}

type ICategoriaService interface {
	AgregarCategoria(categoria modelos.Categoria) error
	ExisteCategoria(nombre string) bool
	VerCategorias()
	ObtenerCategoriaPorID(id int) (modelos.Categoria, bool)
	BuscarCategorias(query string) []modelos.Categoria
}

type ILibroService interface {
	AgregarLibro(libro modelos.Libro) error
	ExisteLibro(titulo string) bool
	ObtenerLibroPorID(id int) (modelos.Libro, bool)
	VerLibros()
	BuscarLibros(query string) []modelos.Libro
}

type IPrestamoService interface {
	CrearPrestamo(libroID int, libroTitulo, estudiante string, limiteMaximo int) error
	RegistrarDevolucion(libroID int) error
	VerHistorialPrestamos()
	VerPrestamos()
}
