package main

import (
	"fmt"
	"sistema_gestion_libros/modelos"
	"sistema_gestion_libros/servicios"
	"sistema_gestion_libros/utilidades"
	"strconv"
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
			servicios.VerLibros()
		case "3":
			agregarAutor()
		case "4":
			servicios.VerAutores()
		case "5":
			agregarCategoria()
		case "6":
			servicios.VerCategorias()
		case "7":
			crearPrestamo()
		case "8":
			servicios.VerPrestamos()
		case "9":
			registrarDevolucion()
		case "10":
			servicios.VerHistorialPrestamos()
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
	resultadosLibros := servicios.BuscarLibros(query)
	for _, libro := range resultadosLibros {
		fmt.Printf("ID: %d, Título: %s, Autor: %s, Categoría: %s\n", libro.ID, libro.Titulo, libro.Autor, libro.Categoria)
	}

	fmt.Println("\nResultados de búsqueda en autores:")
	resultadosAutores := servicios.BuscarAutores(query)
	for _, autor := range resultadosAutores {
		fmt.Printf("ID: %d, Nombre: %s\n", autor.ID, autor.Nombre)
	}

	fmt.Println("\nResultados de búsqueda en categorías:")
	resultadosCategorias := servicios.BuscarCategorias(query)
	for _, categoria := range resultadosCategorias {
		fmt.Printf("ID: %d, Nombre: %s\n", categoria.ID, categoria.Nombre)
	}
}

func cargarDatosPredeterminados() {
	servicios.AgregarAutor(modelos.Autor{Nombre: "Gabriel García Márquez"})
	servicios.AgregarAutor(modelos.Autor{Nombre: "Isabel Allende"})
	servicios.AgregarAutor(modelos.Autor{Nombre: "J.K. Rowling"})

	servicios.AgregarCategoria(modelos.Categoria{Nombre: "Literatura"})
	servicios.AgregarCategoria(modelos.Categoria{Nombre: "Fantasía"})
	servicios.AgregarCategoria(modelos.Categoria{Nombre: "Ciencia Ficción"})

	servicios.AgregarLibro(modelos.Libro{Titulo: "Cien Años de Soledad", Autor: "Gabriel García Márquez", Categoria: "Literatura"})
	servicios.AgregarLibro(modelos.Libro{Titulo: "La Casa de los Espíritus", Autor: "Isabel Allende", Categoria: "Literatura"})
	servicios.AgregarLibro(modelos.Libro{Titulo: "Harry Potter y la Piedra Filosofal", Autor: "J.K. Rowling", Categoria: "Fantasía"})
	fmt.Println("Datos predeterminados cargados exitosamente.")
}

func agregarLibro() {
	titulo := utilidades.LeerEntrada("Ingrese el título del libro: ")
	if servicios.ExisteLibro(titulo) {
		fmt.Println("El libro ya existe en el sistema.")
		return
	}

	servicios.VerCategorias()
	categoriaID, _ := strconv.Atoi(utilidades.LeerEntrada("Seleccione el ID de la categoría: "))
	categoria, existe := servicios.ObtenerCategoriaPorID(categoriaID)
	if !existe {
		fmt.Println("Categoría no encontrada.")
		return
	}

	servicios.VerAutores()
	autorID, _ := strconv.Atoi(utilidades.LeerEntrada("Seleccione el ID del autor: "))
	autor, existe := servicios.ObtenerAutorPorID(autorID)
	if !existe {
		fmt.Println("Autor no encontrado.")
		return
	}

	servicios.AgregarLibro(modelos.Libro{Titulo: titulo, Autor: autor.Nombre, Categoria: categoria.Nombre})
	fmt.Println("Libro agregado correctamente.")
}

func agregarAutor() {
	nombre := utilidades.LeerEntrada("Ingrese el nombre del autor: ")
	if servicios.ExisteAutor(nombre) {
		fmt.Println("El autor ya existe en el sistema.")
		return
	}
	servicios.AgregarAutor(modelos.Autor{Nombre: nombre})
	fmt.Println("Autor agregado correctamente.")
}

func agregarCategoria() {
	nombre := utilidades.LeerEntrada("Ingrese el nombre de la categoría: ")
	if servicios.ExisteCategoria(nombre) {
		fmt.Println("La categoría ya existe en el sistema.")
		return
	}
	servicios.AgregarCategoria(modelos.Categoria{Nombre: nombre})
	fmt.Println("Categoría agregada correctamente.")
}

func crearPrestamo() {
	servicios.VerLibros()
	libroID, _ := strconv.Atoi(utilidades.LeerEntrada("Seleccione el ID del libro: "))
	libro, existe := servicios.ObtenerLibroPorID(libroID)
	if !existe {
		fmt.Println("El libro no está disponible.")
		return
	}

	estudiante := utilidades.LeerEntrada("Ingrese el nombre del estudiante: ")
	servicios.CrearPrestamo(libroID, libro.Titulo, estudiante)
}

func registrarDevolucion() {
	libroID, _ := strconv.Atoi(utilidades.LeerEntrada("Ingrese el ID del libro devuelto: "))
	if servicios.RegistrarDevolucion(libroID) {
		fmt.Println("Devolución registrada exitosamente.")
	} else {
		fmt.Println("No se encontró un préstamo activo para este libro.")
	}
}
