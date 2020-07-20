package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actuaizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actuaizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t := &task{
		nombre:      "Completar mi curso de Go",
		descripcion: "Completar mi curso de Go de platzi esta semana",
	}

	fmt.Printf("%+v\n", t)
	t.marcarCompleta()
	t.actuaizarNombre("Finalizar mi curso de Go")
	t.actuaizarDescripcion("Completar mi curso cuanto antes")
	fmt.Printf("%+v\n", t)
}
