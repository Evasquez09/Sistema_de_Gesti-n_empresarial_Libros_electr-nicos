package main

import (
	"fmt"
	"sistema_gestion_libros/modelos"
	"sistema_gestion_libros/servicios"
	"sistema_gestion_libros/utilidades"
	"strconv"
)

func main() {
	cargarDatosPredeterminados() // Cargar datos predeterminados al iniciar el programa

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
		case "0":
			fmt.Println("Saliendo del sistema...")
			return
		default:
			fmt.Println("Opción no válida, intente de nuevo.")
		}
	}
}

// Función para cargar datos predeterminados
func cargarDatosPredeterminados() {
	// Agregar autores predeterminados
	servicios.AgregarAutor(modelos.Autor{Nombre: "Gabriel García Márquez"})
	servicios.AgregarAutor(modelos.Autor{Nombre: "Isabel Allende"})
	servicios.AgregarAutor(modelos.Autor{Nombre: "J.K. Rowling"})

	// Agregar categorías predeterminadas
	servicios.AgregarCategoria(modelos.Categoria{Nombre: "Literatura"})
	servicios.AgregarCategoria(modelos.Categoria{Nombre: "Fantasía"})
	servicios.AgregarCategoria(modelos.Categoria{Nombre: "Ciencia Ficción"})

	// Agregar libros predeterminados
	servicios.AgregarLibro(modelos.Libro{Titulo: "Cien Años de Soledad", Autor: "Gabriel García Márquez", Categoria: "Literatura"})
	servicios.AgregarLibro(modelos.Libro{Titulo: "La Casa de los Espíritus", Autor: "Isabel Allende", Categoria: "Literatura"})
	servicios.AgregarLibro(modelos.Libro{Titulo: "Harry Potter y la Piedra Filosofal", Autor: "J.K. Rowling", Categoria: "Fantasía"})

	fmt.Println("Datos predeterminados cargados exitosamente.")
}

// Función para agregar un libro
func agregarLibro() {
	titulo := utilidades.LeerEntrada("Ingrese el título del libro: ")
	if servicios.ExisteLibro(titulo) {
		fmt.Println("El libro ya existe en el sistema.")
		return
	}

	fmt.Println("\n--- Seleccione una Categoría (ID) ---")
	servicios.VerCategorias()
	categoriaIDStr := utilidades.LeerEntrada("Ingrese el ID de la categoría: ")
	categoriaID, _ := strconv.Atoi(categoriaIDStr)
	categoria, existe := servicios.ObtenerCategoriaPorID(categoriaID)
	if !existe {
		fmt.Println("Categoría no encontrada.")
		return
	}

	fmt.Println("\n--- Seleccione un Autor (ID) ---")
	servicios.VerAutores()
	autorIDStr := utilidades.LeerEntrada("Ingrese el ID del autor: ")
	autorID, _ := strconv.Atoi(autorIDStr)
	autor, existe := servicios.ObtenerAutorPorID(autorID)
	if !existe {
		fmt.Println("Autor no encontrado.")
		return
	}

	nuevoLibro := modelos.Libro{Titulo: titulo, Autor: autor.Nombre, Categoria: categoria.Nombre}
	servicios.AgregarLibro(nuevoLibro)
	fmt.Println("Libro agregado correctamente.")
}

// Función para agregar un autor
func agregarAutor() {
	nombre := utilidades.LeerEntrada("Ingrese el nombre del autor: ")
	if servicios.ExisteAutor(nombre) {
		fmt.Println("El autor ya existe en el sistema.")
		return
	}
	servicios.AgregarAutor(modelos.Autor{Nombre: nombre})
	fmt.Println("Autor agregado correctamente.")
}

// Función para agregar una categoría
func agregarCategoria() {
	nombre := utilidades.LeerEntrada("Ingrese el nombre de la categoría: ")
	if servicios.ExisteCategoria(nombre) {
		fmt.Println("La categoría ya existe en el sistema.")
		return
	}
	servicios.AgregarCategoria(modelos.Categoria{Nombre: nombre})
	fmt.Println("Categoría agregada correctamente.")
}

// Función para crear un préstamo
func crearPrestamo() {
	fmt.Println("\n--- Lista de Libros Disponibles (ID) ---")
	servicios.VerLibros() // Mostrar la lista de libros con ID

	libroIDStr := utilidades.LeerEntrada("Ingrese el ID del libro para el préstamo: ")
	libroID, _ := strconv.Atoi(libroIDStr)
	libro, existe := servicios.ObtenerLibroPorID(libroID)
	if !existe {
		fmt.Println("El libro no está disponible en la biblioteca.")
		return
	}

	estudiante := utilidades.LeerEntrada("Ingrese el nombre del estudiante: ")
	servicios.CrearPrestamo(libro.Titulo, estudiante)
}
