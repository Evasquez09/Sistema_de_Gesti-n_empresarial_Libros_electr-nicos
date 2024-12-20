package servicios

import "sistema_gestion_libros/modelos"

type IAutorService interface {
	AgregarAutor(autor modelos.Autor) error
	ExisteAutor(nombre string) bool
	ObtenerAutorPorID(id int) (modelos.Autor, bool)
	VerAutores()
	BuscarAutores(query string) []modelos.Autor
	ObtenerTodos() []modelos.Autor
	ActualizarAutor(autor modelos.Autor) error
	EliminarAutor(id int) error
}

type ICategoriaService interface {
	AgregarCategoria(categoria modelos.Categoria) error
	ExisteCategoria(nombre string) bool
	VerCategorias()
	ObtenerCategoriaPorID(id int) (modelos.Categoria, bool)
	BuscarCategorias(query string) []modelos.Categoria
	ObtenerTodas() []modelos.Categoria
	ActualizarCategoria(cat modelos.Categoria) error
	EliminarCategoria(id int) error
}

type ILibroService interface {
	AgregarLibro(libro modelos.Libro) error
	ExisteLibro(titulo string) bool
	ObtenerLibroPorID(id int) (modelos.Libro, bool)
	VerLibros()
	BuscarLibros(query string) []modelos.Libro
	ObtenerTodos() []modelos.Libro
	ActualizarLibro(libro modelos.Libro) error
	EliminarLibro(id int) error
}

type IPrestamoService interface {
	CrearPrestamo(libroID int, libroTitulo, estudiante string, limiteMaximo int) error
	RegistrarDevolucion(libroID int) error
	VerHistorialPrestamos()
	VerPrestamos()
	ObtenerActivos() []modelos.Prestamo
	ObtenerHistorial() []modelos.Prestamo
}
