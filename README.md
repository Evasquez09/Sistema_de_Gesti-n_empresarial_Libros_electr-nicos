# Sistema de Gestión de Libros Electrónicos

Este sistema permite gestionar autores, categorías, libros y préstamos de forma eficiente mediante una interfaz web y servicios web RESTful. La base de datos utilizada es MariaDB o MySQL.

## Funcionalidades Principales

### Autores
- **Agregar Autor:** Permite agregar un nuevo autor.
- **Actualizar Autor:** Actualiza el nombre de un autor existente.
- **Eliminar Autor:** Elimina un autor de la base de datos.
- **Ver Autores (JSON):** Muestra una lista de todos los autores en formato JSON.

### Categorías
- **Agregar Categoría:** Permite agregar una nueva categoría.
- **Actualizar Categoría:** Actualiza el nombre de una categoría existente.
- **Eliminar Categoría:** Elimina una categoría de la base de datos.
- **Ver Categorías (JSON):** Muestra una lista de todas las categorías en formato JSON.

### Libros
- **Agregar Libro:** Permite agregar un nuevo libro con un título, autor y categoría.
- **Actualizar Libro:** Actualiza los datos de un libro existente.
- **Eliminar Libro:** Elimina un libro de la base de datos.
- **Ver Libros (JSON):** Muestra una lista de todos los libros en formato JSON.

### Préstamos
- **Crear Préstamo:** Registra un préstamo de un libro para un estudiante con una fecha de devolución predeterminada de 30 días.
  - **Validación:** Un estudiante no puede tener dos préstamos activos del mismo libro.
- **Registrar Devolución:** Marca un préstamo como devuelto eliminándolo de la base de datos.
- **Ver Préstamos Activos (JSON):** Muestra todos los préstamos activos.
- **Ver Historial de Préstamos (JSON):** Devuelve un historial (implementación personalizada).
- **Error personalizado para préstamos duplicados:** Si un estudiante intenta prestar el mismo libro dos veces, el sistema devuelve un mensaje descriptivo.

### Búsqueda
- **Buscar Libros, Categorías o Autores:** Permite buscar elementos por palabra clave.

## Características Técnicas

### Base de Datos
Se utiliza MariaDB o MySQL. La estructura de la base de datos incluye:

#### Tablas Principales
1. **autores**
   ```sql
   CREATE TABLE autores (
       id INT AUTO_INCREMENT PRIMARY KEY,
       nombre VARCHAR(255) NOT NULL
   );
   ```

2. **categorias**
   ```sql
   CREATE TABLE categorias (
       id INT AUTO_INCREMENT PRIMARY KEY,
       nombre VARCHAR(255) NOT NULL
   );
   ```

3. **libros**
   ```sql
   CREATE TABLE libros (
       id INT AUTO_INCREMENT PRIMARY KEY,
       titulo VARCHAR(255) NOT NULL,
       autor_id INT,
       categoria_id INT,
       FOREIGN KEY (autor_id) REFERENCES autores(id),
       FOREIGN KEY (categoria_id) REFERENCES categorias(id)
   );
   ```

4. **prestamos**
   ```sql
   CREATE TABLE prestamos (
       id INT AUTO_INCREMENT PRIMARY KEY,
       libro_id INT NOT NULL,
       estudiante VARCHAR(255) NOT NULL,
       fecha_prestamo DATE NOT NULL,
       fecha_devolucion DATE NOT NULL,
       enlace VARCHAR(255) NOT NULL,
       FOREIGN KEY (libro_id) REFERENCES libros(id),
       UNIQUE (libro_id, estudiante)
   );
   ```

### Tecnologías Utilizadas
- **Backend:** Go (Golang)
- **Frontend:** HTML, CSS
- **Base de Datos:** MariaDB/MySQL
- **ORM/Driver:** `github.com/go-sql-driver/mysql`

### Validaciones Importantes
- Un estudiante no puede tener dos préstamos activos del mismo libro.
- Si ocurre un error, se muestran mensajes personalizados.

### Servicios Web RESTful
El sistema ofrece servicios web en formato JSON accesibles desde rutas específicas:
- `GET /autores`: Devuelve todos los autores.
- `POST /autores`: Agrega un nuevo autor.
- `PUT /autores`: Actualiza un autor existente.
- `DELETE /autores`: Elimina un autor.
- (Similar para categorías y libros).

### Enlaces Únicos
Cada préstamo genera un enlace único para identificación.

## Configuración Inicial

1. Clonar este repositorio:
   ```bash
   git clone https://github.com/tu-usuario/sistema-gestion-libros.git
   cd sistema-gestion-libros
   ```

2. Configurar la base de datos en `utilidades/helpers.go`:
   ```go
   const (
       DBUser     = "tu_usuario"
       DBPassword = "tu_contraseña"
       DBName     = "sistema_libros"
   )
   ```

3. Crear la base de datos y ejecutar el script SQL proporcionado:
   ```bash
   mysql -u tu_usuario -p sistema_libros < script.sql
   ```

4. Ejecutar el proyecto:
   ```bash
   go run main.go
   ```

5. Acceder a la aplicación desde:
   ```
   http://localhost:8080
   