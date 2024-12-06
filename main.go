package main

import (
	"errors"
	"fmt"
	"sistema_gestion_libros/modelos"
	"sistema_gestion_libros/servicios"
	"sistema_gestion_libros/utilidades"
	"strconv"
)

var limiteMaximoPrestamos int = 3

// Se crean variables globales que implementan las interfaces definidas en servicios
var (
	libroService     servicios.ILibroService     = servicios.NewLibroService()
	autorService     servicios.IAutorService     = servicios.NewAutorService()
	categoriaService servicios.ICategoriaService = servicios.NewCategoriaService()
	prestamoService  servicios.IPrestamoService  = servicios.NewPrestamoService()
)

func main() {
	cargarDatosPredeterminados()

	for {
		fmt.Println("\n--- Sistema de Gestión de Libros Electrónicos ---")
		fmt.Println("1. Agregar libro")
		fmt.Println("2. Ver libros")
		fmt.Println("3. Agregar autor")
		fmt.Println("4. Ver autores")
		fmt.Println("5. Agregar categoría")
		fmt.Println("6. Ver categorías")
		fmt.Println("7. Crear préstamo")
		fmt.Println("8. Ver préstamos activos")
		fmt.Println("9. Registrar devolución")
		fmt.Println("10. Historial de préstamos")
		fmt.Println("11. Buscar libros, autores o categorías")
		fmt.Println("0. Salir")

		opcion := utilidades.LeerEntrada("Seleccione una opción: ")

		switch opcion {
		case "1":
			agregarLibro()
		case "2":
			verLibros()
		case "3":
			agregarAutor()
		case "4":
			verAutores()
		case "5":
			agregarCategoria()
		case "6":
			verCategorias()
		case "7":
			crearPrestamo()
		case "8":
			verPrestamos()
		case "9":
			registrarDevolucion()
		case "10":
			verHistorialPrestamos()
		case "11":
			realizarBusqueda()
		case "0":
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opción no válida, intente de nuevo.")
		}
	}
}

func realizarBusqueda() {
	query := utilidades.LeerEntrada("Ingrese texto de búsqueda: ")
	fmt.Println("\nResultados de búsqueda en libros:")
	resultadosLibros := libroService.BuscarLibros(query)
	for _, libro := range resultadosLibros {
		fmt.Printf("ID: %d, Título: %s, Autor: %s, Categoría: %s\n", libro.ID, libro.Titulo, libro.Autor, libro.Categoria)
	}

	fmt.Println("\nResultados de búsqueda en autores:")
	resultadosAutores := autorService.BuscarAutores(query)
	for _, autor := range resultadosAutores {
		fmt.Printf("ID: %d, Nombre: %s\n", autor.ID, autor.Nombre)
	}

	fmt.Println("\nResultados de búsqueda en categorías:")
	resultadosCategorias := categoriaService.BuscarCategorias(query)
	for _, categoria := range resultadosCategorias {
		fmt.Printf("ID: %d, Nombre: %s\n", categoria.ID, categoria.Nombre)
	}
}

func cargarDatosPredeterminados() {
	_ = autorService.AgregarAutor(modelos.Autor{Nombre: "Gabriel García Márquez"})
	_ = autorService.AgregarAutor(modelos.Autor{Nombre: "Isabel Allende"})
	_ = autorService.AgregarAutor(modelos.Autor{Nombre: "J.K. Rowling"})

	_ = categoriaService.AgregarCategoria(modelos.Categoria{Nombre: "Literatura"})
	_ = categoriaService.AgregarCategoria(modelos.Categoria{Nombre: "Fantasía"})
	_ = categoriaService.AgregarCategoria(modelos.Categoria{Nombre: "Ciencia Ficción"})

	_ = libroService.AgregarLibro(modelos.Libro{Titulo: "Cien Años de Soledad", Autor: "Gabriel García Márquez", Categoria: "Literatura"})
	_ = libroService.AgregarLibro(modelos.Libro{Titulo: "La Casa de los Espíritus", Autor: "Isabel Allende", Categoria: "Literatura"})
	_ = libroService.AgregarLibro(modelos.Libro{Titulo: "Harry Potter y la Piedra Filosofal", Autor: "J.K. Rowling", Categoria: "Fantasía"})
	fmt.Println("Datos predeterminados cargados exitosamente.")
}

func agregarLibro() {
	titulo := utilidades.LeerEntrada("Ingrese el título del libro: ")
	existe := libroService.ExisteLibro(titulo)
	if existe {
		fmt.Println("El libro ya existe en el sistema.")
		return
	}

	verCategorias()
	categoriaID, err := strconv.Atoi(utilidades.LeerEntrada("Seleccione el ID de la categoría: "))
	if err != nil {
		fmt.Println("ID inválido.")
		return
	}
	categoria, existeCat := categoriaService.ObtenerCategoriaPorID(categoriaID)
	if !existeCat {
		fmt.Println("Categoría no encontrada.")
		return
	}

	verAutores()
	autorID, err := strconv.Atoi(utilidades.LeerEntrada("Seleccione el ID del autor: "))
	if err != nil {
		fmt.Println("ID inválido.")
		return
	}
	autor, existeAut := autorService.ObtenerAutorPorID(autorID)
	if !existeAut {
		fmt.Println("Autor no encontrado.")
		return
	}

	err = libroService.AgregarLibro(modelos.Libro{Titulo: titulo, Autor: autor.Nombre, Categoria: categoria.Nombre})
	if err != nil {
		fmt.Println("Error al agregar libro:", err)
		return
	}
	fmt.Println("Libro agregado correctamente.")
}

func agregarAutor() {
	nombre := utilidades.LeerEntrada("Ingrese el nombre del autor: ")
	if autorService.ExisteAutor(nombre) {
		fmt.Println("El autor ya existe en el sistema.")
		return
	}
	if err := autorService.AgregarAutor(modelos.Autor{Nombre: nombre}); err != nil {
		fmt.Println("Error al agregar autor:", err)
		return
	}
	fmt.Println("Autor agregado correctamente.")
}

func agregarCategoria() {
	nombre := utilidades.LeerEntrada("Ingrese el nombre de la categoría: ")
	if categoriaService.ExisteCategoria(nombre) {
		fmt.Println("La categoría ya existe en el sistema.")
		return
	}
	if err := categoriaService.AgregarCategoria(modelos.Categoria{Nombre: nombre}); err != nil {
		fmt.Println("Error al agregar categoría:", err)
		return
	}
	fmt.Println("Categoría agregada correctamente.")
}

func crearPrestamo() {
	fmt.Println("\n--- Lista de Libros Disponibles (ID) ---")
	verLibros()

	libroIDStr := utilidades.LeerEntrada("Ingrese el ID del libro para el préstamo: ")
	libroID, err := strconv.Atoi(libroIDStr)
	if err != nil {
		fmt.Println("ID de libro inválido.")
		return
	}
	libro, existe := libroService.ObtenerLibroPorID(libroID)
	if !existe {
		fmt.Println("El libro no está disponible en la biblioteca.")
		return
	}

	estudiante := utilidades.LeerEntrada("Ingrese el nombre del estudiante: ")
	if err := prestamoService.CrearPrestamo(libroID, libro.Titulo, estudiante, limiteMaximoPrestamos); err != nil {
		fmt.Println("Error al crear préstamo:", err)
		return
	}
}

func registrarDevolucion() {
	libroID, err := strconv.Atoi(utilidades.LeerEntrada("Ingrese el ID del libro devuelto: "))
	if err != nil {
		fmt.Println("ID inválido.")
		return
	}
	if err := prestamoService.RegistrarDevolucion(libroID); err != nil {
		if errors.Is(err, servicios.ErrPrestamoNoEncontrado) {
			fmt.Println("No se encontró un préstamo activo para este libro.")
		} else {
			fmt.Println("Error al registrar devolución:", err)
		}
		return
	}
	fmt.Println("Devolución registrada exitosamente.")
}

func verLibros() {
	libroService.VerLibros()
}

func verAutores() {
	autorService.VerAutores()
}

func verCategorias() {
	categoriaService.VerCategorias()
}

func verPrestamos() {
	prestamoService.VerPrestamos()
}

func verHistorialPrestamos() {
	prestamoService.VerHistorialPrestamos()
}
