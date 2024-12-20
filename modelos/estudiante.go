package modelos

type Estudiante struct {
	ID        int    `json:"id"`
	Nombre    string `json:"nombre"`
	Carrera   string `json:"carrera"`
	Matricula string `json:"matricula"`
}
