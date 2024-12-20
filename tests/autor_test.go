package tests

import (
	"sistema_gestion_libros/modelos"
	"sistema_gestion_libros/servicios"
	"testing"
)

func TestAgregarAutor(t *testing.T) {
	s := servicios.NewAutorService()
	err := s.AgregarAutor(modelos.Autor{Nombre: "Test Autor"})
	if err != nil {
		t.Fatalf("No se pudo agregar autor: %v", err)
	}

	if !s.ExisteAutor("Test Autor") {
		t.Errorf("El autor no se agregó correctamente")
	}
}
