package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sistema_gestion_libros/handlers"
	"sistema_gestion_libros/servicios"

	_ "github.com/go-sql-driver/mysql"
)

var templates *template.Template

func main() {
	db := conectarBD()

	// Inicializar los servicios
	autorService := servicios.NewAutorService(db)
	categoriaService := servicios.NewCategoriaService(db)
	libroService := servicios.NewLibroService(db)
	prestamoService := servicios.NewPrestamoService(db)

	// Cargar templates
	templates = template.Must(template.ParseGlob("templates/*.html"))

	// Configurar enrutador
	mux := http.NewServeMux()

	// Ruta para el menú principal
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		renderTemplate(w, "index.html", nil)
	})

	// Handlers para las demás rutas
	mux.HandleFunc("/autores", handlers.AutoresHandler(autorService, templates))
	mux.HandleFunc("/autores/", handlers.AutoresHandler(autorService, templates))
	mux.HandleFunc("/categorias", handlers.CategoriasHandler(categoriaService, templates))
	mux.HandleFunc("/categorias/", handlers.CategoriasHandler(categoriaService, templates))
	mux.HandleFunc("/libros", handlers.LibrosHandler(libroService, autorService, categoriaService, templates))
	mux.HandleFunc("/prestamos", handlers.PrestamosHandler(prestamoService, libroService, templates))
	mux.HandleFunc("/buscar", handlers.BuscarHandler(libroService, autorService, categoriaService, templates))
	fmt.Println("Servidor iniciado en http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func conectarBD() *sql.DB {
	dsn := "gestion_libros:gestion_libros2024$@tcp(127.0.0.1:3306)/gestion_libros?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	if err := db.Ping(); err != nil {
		log.Fatal("No se pudo hacer ping a la base de datos:", err)
	}
	return db
}

func renderTemplate(w http.ResponseWriter, filename string, data interface{}) {
	err := templates.ExecuteTemplate(w, filename, data)
	if err != nil {
		http.Error(w, "Error al renderizar la plantilla", http.StatusInternalServerError)
	}
}
