package servicios

import (
	"errors"
	"fmt"
	"sistema_gestion_libros/modelos"
	"strings"
)

type estudianteService struct {
	estudiantes         []modelos.Estudiante
	estudianteIDCounter int
}

func (s *estudianteService) AgregarEstudiante(e modelos.Estudiante) error {
	for _, est := range s.estudiantes {
		if est.Matricula == e.Matricula {
			return errors.New("el estudiante ya existe con esa matrícula")
		}
	}
	e.ID = s.estudianteIDCounter
	s.estudiantes = append(s.estudiantes, e)
	s.estudianteIDCounter++
	fmt.Printf("Estudiante agregado: %s (ID: %d)\n", e.Nombre, e.ID)
	return nil
}

func (s *estudianteService) ObtenerTodos() []modelos.Estudiante {
	return s.estudiantes
}

func (s *estudianteService) ExisteEstudiante(nombre string) bool {
	for _, est := range s.estudiantes {
		if strings.ToLower(est.Nombre) == strings.ToLower(nombre) {
			return true
		}
	}
	return false
}

func (s *estudianteService) ActualizarEstudiante(e modelos.Estudiante) error {
	for i, est := range s.estudiantes {
		if est.ID == e.ID {
			if e.Nombre == "" || e.Matricula == "" {
				return errors.New("nombre y matrícula requeridos")
			}
			s.estudiantes[i] = e
			return nil
		}
	}
	return errors.New("estudiante no encontrado")
}

func (s *estudianteService) EliminarEstudiante(id int) error {
	for i, est := range s.estudiantes {
		if est.ID == id {
			s.estudiantes = append(s.estudiantes[:i], s.estudiantes[i+1:]...)
			return nil
		}
	}
	return errors.New("estudiante no encontrado")
}

func (s *estudianteService) ObtenerEstudiantePorID(id int) (modelos.Estudiante, bool) {
	for _, est := range s.estudiantes {
		if est.ID == id {
			return est, true
		}
	}
	return modelos.Estudiante{}, false
}

func (s *estudianteService) BuscarEstudiantes(query string) []modelos.Estudiante {
	var resultados []modelos.Estudiante
	for _, est := range s.estudiantes {
		if strings.Contains(strings.ToLower(est.Nombre), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(est.Carrera), strings.ToLower(query)) {
			resultados = append(resultados, est)
		}
	}
	return resultados
}
