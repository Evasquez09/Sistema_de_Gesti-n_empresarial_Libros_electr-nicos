package handlers

import (
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"sistema_gestion_libros/modelos"
	"sistema_gestion_libros/servicios"
	"strconv"
	"strings"
)

func MenuHandler(templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Renderizar el template "menu.html"
		err := templates.ExecuteTemplate(w, "menu.html", nil)
		if err != nil {
			http.Error(w, "Error al renderizar el template: "+err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func BuscarHandler(libroService servicios.ILibroService, autorService servicios.IAutorService, categoriaService servicios.ICategoriaService, templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Verifica que el método sea GET
		if r.Method != http.MethodGet {
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
			return
		}

		// Obtén el parámetro de búsqueda
		query := r.URL.Query().Get("q")
		if query == "" {
			http.Error(w, "El parámetro de búsqueda es obligatorio", http.StatusBadRequest)
			return
		}

		// Imprime la búsqueda para depuración
		log.Println("Búsqueda recibida:", query)

		// Realiza las búsquedas
		libros := libroService.BuscarLibros(query)
		autores := autorService.BuscarAutores(query)
		categorias := categoriaService.BuscarCategorias(query)

		// Verifica si no hay resultados
		if len(libros) == 0 && len(autores) == 0 && len(categorias) == 0 {
			http.Error(w, "No se encontraron resultados", http.StatusNotFound)
			return
		}

		// Crea un mapa con los resultados
		resultados := map[string]interface{}{
			"libros":     libros,
			"autores":    autores,
			"categorias": categorias,
		}

		// Renderiza los resultados en la plantilla
		err := templates.ExecuteTemplate(w, "resultados.html", resultados)
		if err != nil {
			http.Error(w, "Error al renderizar los resultados de la búsqueda", http.StatusInternalServerError)
		}
	}
}

func AutoresHandler(autorService servicios.IAutorService, templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		realMethod := r.Method
		if m := r.FormValue("_method"); m != "" {
			realMethod = strings.ToUpper(m)
		}

		switch realMethod {
		case http.MethodGet:
			autores := autorService.ObtenerTodos()
			if r.Header.Get("Accept") == "application/json" {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(autores)
				return
			}

			// Renderizar plantilla si no es JSON
			err := templates.ExecuteTemplate(w, "autores.html", autores)
			if err != nil {
				http.Error(w, "Error al renderizar la plantilla de autores", http.StatusInternalServerError)
			}

		case http.MethodPost:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Datos inválidos", http.StatusBadRequest)
				return
			}
			nombre := r.FormValue("nombre")
			if nombre == "" {
				http.Error(w, "El nombre del autor es obligatorio", http.StatusBadRequest)
				return
			}
			autor := modelos.Autor{Nombre: nombre}
			err = autorService.AgregarAutor(autor)
			if err != nil {
				http.Error(w, "Error al agregar el autor: "+err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/autores", http.StatusSeeOther)

		case http.MethodPut:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Datos inválidos", http.StatusBadRequest)
				return
			}
			id := r.FormValue("id")
			nombre := r.FormValue("nombre")
			autorID, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "ID inválido", http.StatusBadRequest)
				return
			}
			autor := modelos.Autor{ID: autorID, Nombre: nombre}
			err = autorService.ActualizarAutor(autor)
			if err != nil {
				http.Error(w, "Error al actualizar el autor: "+err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/autores", http.StatusSeeOther)

		case http.MethodDelete:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Datos inválidos", http.StatusBadRequest)
				return
			}
			id := r.FormValue("id")
			autorID, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "ID inválido", http.StatusBadRequest)
				return
			}
			err = autorService.EliminarAutor(autorID)
			if err != nil {
				http.Error(w, "Error al eliminar el autor: "+err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/autores", http.StatusSeeOther)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}
}

// LibrosHandler
func LibrosHandler(libroService servicios.ILibroService, autorService servicios.IAutorService, categoriaService servicios.ICategoriaService, templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		realMethod := r.Method
		if m := r.FormValue("_method"); m != "" {
			realMethod = strings.ToUpper(m)
		}

		switch realMethod {
		case http.MethodGet:
			libros := libroService.ObtenerTodos()
			if r.Header.Get("Accept") == "application/json" {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(libros)
				return
			}
			err := templates.ExecuteTemplate(w, "libros.html", libros)
			if err != nil {
				http.Error(w, "Error al renderizar la plantilla de libros", http.StatusInternalServerError)
			}

		case http.MethodPost:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Datos inválidos", http.StatusBadRequest)
				return
			}
			titulo := r.FormValue("titulo")
			autor := r.FormValue("autor")
			categoria := r.FormValue("categoria")
			if titulo == "" || autor == "" || categoria == "" {
				http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
				return
			}
			libro := modelos.Libro{Titulo: titulo, Autor: autor, Categoria: categoria}
			err = libroService.AgregarLibro(libro)
			if err != nil {
				http.Error(w, "Error al agregar el libro: "+err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/libros", http.StatusSeeOther)

		case http.MethodPut:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Datos inválidos", http.StatusBadRequest)
				return
			}
			id := r.FormValue("id")
			titulo := r.FormValue("titulo")
			autor := r.FormValue("autor")
			categoria := r.FormValue("categoria")
			libroID, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "ID inválido", http.StatusBadRequest)
				return
			}
			libro := modelos.Libro{ID: libroID, Titulo: titulo, Autor: autor, Categoria: categoria}
			err = libroService.ActualizarLibro(libro)
			if err != nil {
				http.Error(w, "Error al actualizar el libro: "+err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/libros", http.StatusSeeOther)

		case http.MethodDelete:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Datos inválidos", http.StatusBadRequest)
				return
			}
			id := r.FormValue("id")
			libroID, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "ID inválido", http.StatusBadRequest)
				return
			}
			err = libroService.EliminarLibro(libroID)
			if err != nil {
				http.Error(w, "Error al eliminar el libro: "+err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/libros", http.StatusSeeOther)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}
}
func PrestamosHandler(prestamoService servicios.IPrestamoService, libroService servicios.ILibroService, templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		realMethod := r.Method
		if m := r.FormValue("_method"); m != "" {
			realMethod = strings.ToUpper(m)
		}

		switch realMethod {
		case http.MethodGet:
			prestamos := prestamoService.ObtenerActivos()
			if r.Header.Get("Accept") == "application/json" {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(prestamos)
				return
			}
			err := templates.ExecuteTemplate(w, "prestamos.html", prestamos)
			if err != nil {
				http.Error(w, "Error al renderizar la plantilla de préstamos", http.StatusInternalServerError)
			}

		case http.MethodPost:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Datos inválidos", http.StatusBadRequest)
				return
			}
			libroIDStr := r.FormValue("libro_id")
			estudiante := r.FormValue("estudiante")
			libroID, err := strconv.Atoi(libroIDStr)
			if err != nil || estudiante == "" {
				http.Error(w, "Libro ID y Estudiante son obligatorios", http.StatusBadRequest)
				return
			}
			libro, existe := libroService.ObtenerLibroPorID(libroID)
			if !existe {
				http.Error(w, "El libro no existe", http.StatusNotFound)
				return
			}
			err = prestamoService.CrearPrestamo(libro.ID, libro.Titulo, estudiante, 3)
			if err != nil {
				http.Error(w, "Error al crear el préstamo: "+err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/prestamos", http.StatusSeeOther)

		case http.MethodPut:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Datos inválidos", http.StatusBadRequest)
				return
			}
			libroIDStr := r.FormValue("libro_id")
			libroID, err := strconv.Atoi(libroIDStr)
			if err != nil {
				http.Error(w, "Libro ID inválido", http.StatusBadRequest)
				return
			}
			err = prestamoService.RegistrarDevolucion(libroID)
			if err != nil {
				http.Error(w, "Error al registrar la devolución: "+err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/prestamos", http.StatusSeeOther)

		case http.MethodDelete:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Datos inválidos", http.StatusBadRequest)
				return
			}
			libroIDStr := r.FormValue("libro_id")
			libroID, err := strconv.Atoi(libroIDStr)
			if err != nil {
				http.Error(w, "Libro ID inválido", http.StatusBadRequest)
				return
			}
			err = prestamoService.RegistrarDevolucion(libroID) // En este caso, la devolución elimina el préstamo
			if err != nil {
				http.Error(w, "Error al eliminar el préstamo: "+err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/prestamos", http.StatusSeeOther)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}
}
func CategoriasHandler(categoriaService servicios.ICategoriaService, templates *template.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		realMethod := r.Method
		if m := r.FormValue("_method"); m != "" {
			realMethod = strings.ToUpper(m)
		}

		switch realMethod {
		case http.MethodGet:
			categorias := categoriaService.ObtenerTodas()
			if r.Header.Get("Accept") == "application/json" {
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(categorias)
				return
			}
			err := templates.ExecuteTemplate(w, "categorias.html", categorias)
			if err != nil {
				http.Error(w, "Error al renderizar la plantilla de categorías", http.StatusInternalServerError)
			}

		case http.MethodPost:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Datos inválidos", http.StatusBadRequest)
				return
			}
			nombre := r.FormValue("nombre")
			if nombre == "" {
				http.Error(w, "El nombre de la categoría es obligatorio", http.StatusBadRequest)
				return
			}
			categoria := modelos.Categoria{Nombre: nombre}
			err = categoriaService.AgregarCategoria(categoria)
			if err != nil {
				http.Error(w, "Error al agregar la categoría: "+err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/categorias", http.StatusSeeOther)

		case http.MethodPut:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Datos inválidos", http.StatusBadRequest)
				return
			}
			id := r.FormValue("id")
			nombre := r.FormValue("nombre")
			categoriaID, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "ID inválido", http.StatusBadRequest)
				return
			}
			categoria := modelos.Categoria{ID: categoriaID, Nombre: nombre}
			err = categoriaService.ActualizarCategoria(categoria)
			if err != nil {
				http.Error(w, "Error al actualizar la categoría: "+err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/categorias", http.StatusSeeOther)

		case http.MethodDelete:
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Datos inválidos", http.StatusBadRequest)
				return
			}
			id := r.FormValue("id")
			categoriaID, err := strconv.Atoi(id)
			if err != nil {
				http.Error(w, "ID inválido", http.StatusBadRequest)
				return
			}
			err = categoriaService.EliminarCategoria(categoriaID)
			if err != nil {
				http.Error(w, "Error al eliminar la categoría: "+err.Error(), http.StatusInternalServerError)
				return
			}
			http.Redirect(w, r, "/categorias", http.StatusSeeOther)

		default:
			http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		}
	}
}
