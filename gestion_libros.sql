-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Servidor: 127.0.0.1
-- Tiempo de generación: 20-12-2024 a las 03:48:05
-- Versión del servidor: 10.4.32-MariaDB
-- Versión de PHP: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Base de datos: `gestion_libros`
--

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `autores`
--

CREATE TABLE `autores` (
  `id` int(11) NOT NULL,
  `nombre` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `autores`
--

INSERT INTO `autores` (`id`, `nombre`) VALUES
(7, 'Agatha Christie'),
(10, 'Arthur Conan Doyle'),
(14, 'Charles Dickens'),
(17, 'Edgar Allan Poe'),
(26, 'Edwin Vasquez'),
(8, 'Ernest Hemingway'),
(19, 'Fiódor Dostoyevski'),
(16, 'Franz Kafka'),
(1, 'Gabriel García Márquez'),
(6, 'George R.R. Martin'),
(20, 'H.P. Lovecraft'),
(15, 'Haruki Murakami'),
(2, 'Isabel Allende'),
(3, 'J.K. Rowling'),
(4, 'J.R.R. Tolkien'),
(9, 'Jane Austen'),
(13, 'Julio Verne'),
(18, 'Leo Tolstoy'),
(11, 'Mark Twain'),
(12, 'Miguel de Cervantes'),
(5, 'Stephen King');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `categorias`
--

CREATE TABLE `categorias` (
  `id` int(11) NOT NULL,
  `nombre` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `categorias`
--

INSERT INTO `categorias` (`id`, `nombre`) VALUES
(7, 'Aventura'),
(3, 'Ciencia Ficción'),
(10, 'Ensayo'),
(2, 'Fantasía'),
(8, 'Histórica'),
(1, 'Literatura'),
(5, 'Misterio'),
(9, 'Poesía'),
(6, 'Romance'),
(4, 'Terror');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `libros`
--

CREATE TABLE `libros` (
  `id` int(11) NOT NULL,
  `titulo` varchar(255) NOT NULL,
  `autor` varchar(255) NOT NULL,
  `categoria` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `libros`
--

INSERT INTO `libros` (`id`, `titulo`, `autor`, `categoria`) VALUES
(1, 'Cien Años de Soledad', 'Gabriel García Márquez', 'Literatura'),
(2, 'El Amor en los Tiempos del Cólera', 'Gabriel García Márquez', 'Literatura'),
(3, 'La Casa de los Espíritus', 'Isabel Allende', 'Literatura'),
(4, 'Harry Potter y la Piedra Filosofal', 'J.K. Rowling', 'Fantasía'),
(5, 'Harry Potter y la Cámara Secreta', 'J.K. Rowling', 'Fantasía'),
(6, 'El Señor de los Anillos: La Comunidad del Anillo', 'J.R.R. Tolkien', 'Fantasía'),
(7, 'El Señor de los Anillos: Las Dos Torres', 'J.R.R. Tolkien', 'Fantasía'),
(8, 'El Resplandor', 'Stephen King', 'Terror'),
(9, 'It', 'Stephen King', 'Terror'),
(10, 'Juego de Tronos', 'George R.R. Martin', 'Fantasía'),
(11, 'Choque de Reyes', 'George R.R. Martin', 'Fantasía'),
(12, 'Diez Negritos', 'Agatha Christie', 'Misterio'),
(13, 'Asesinato en el Orient Express', 'Agatha Christie', 'Misterio'),
(14, 'El Viejo y el Mar', 'Ernest Hemingway', 'Literatura'),
(15, 'Orgullo y Prejuicio', 'Jane Austen', 'Romance'),
(16, 'Sherlock Holmes: Estudio en Escarlata', 'Arthur Conan Doyle', 'Misterio'),
(17, 'Las Aventuras de Tom Sawyer', 'Mark Twain', 'Aventura'),
(18, 'Don Quijote de la Mancha', 'Miguel de Cervantes', 'Literatura'),
(19, 'Veinte mil leguas de viaje submarino', 'Julio Verne', 'Ciencia Ficción'),
(20, 'Oliver Twist', 'Charles Dickens', 'Literatura'),
(21, 'Kafka en la Orilla', 'Haruki Murakami', 'Literatura'),
(22, 'La Metamorfosis', 'Franz Kafka', 'Literatura'),
(23, 'Narraciones Extraordinarias', 'Edgar Allan Poe', 'Terror'),
(24, 'Guerra y Paz', 'Leo Tolstoy', 'Histórica'),
(25, 'Crimen y Castigo', 'Fiódor Dostoyevski', 'Literatura'),
(26, 'La Llamada de Cthulhu', 'H.P. Lovecraft', 'Terror');

-- --------------------------------------------------------

--
-- Estructura de tabla para la tabla `prestamos`
--

CREATE TABLE `prestamos` (
  `id` int(11) NOT NULL,
  `libro_id` int(11) NOT NULL,
  `libro` varchar(255) NOT NULL,
  `estudiante` varchar(255) NOT NULL,
  `fecha_prestamo` datetime NOT NULL,
  `fecha_devolucion` datetime NOT NULL,
  `enlace` varchar(255) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Volcado de datos para la tabla `prestamos`
--

INSERT INTO `prestamos` (`id`, `libro_id`, `libro`, `estudiante`, `fecha_prestamo`, `fecha_devolucion`, `enlace`) VALUES
(8, 1, 'Cien Años de Soledad', 'Edwin Vasquez', '2024-12-19 21:45:05', '2025-01-18 21:45:05', ''),
(10, 2, 'El Amor en los Tiempos del Cólera', 'Edwin Vasquez', '2024-12-19 21:45:15', '2025-01-18 21:45:15', '');

--
-- Índices para tablas volcadas
--

--
-- Indices de la tabla `autores`
--
ALTER TABLE `autores`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `nombre` (`nombre`);

--
-- Indices de la tabla `categorias`
--
ALTER TABLE `categorias`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `nombre` (`nombre`);

--
-- Indices de la tabla `libros`
--
ALTER TABLE `libros`
  ADD PRIMARY KEY (`id`),
  ADD KEY `autor` (`autor`),
  ADD KEY `categoria` (`categoria`);

--
-- Indices de la tabla `prestamos`
--
ALTER TABLE `prestamos`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `libro_id` (`libro_id`,`estudiante`);

--
-- AUTO_INCREMENT de las tablas volcadas
--

--
-- AUTO_INCREMENT de la tabla `autores`
--
ALTER TABLE `autores`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=27;

--
-- AUTO_INCREMENT de la tabla `categorias`
--
ALTER TABLE `categorias`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;

--
-- AUTO_INCREMENT de la tabla `libros`
--
ALTER TABLE `libros`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=28;

--
-- AUTO_INCREMENT de la tabla `prestamos`
--
ALTER TABLE `prestamos`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
