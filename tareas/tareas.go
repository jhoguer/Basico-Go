package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeList(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) imprimirLista() {
	for _, tarea := range t.tasks {
		fmt.Println("Nombre", tarea.nombre)
		fmt.Println("Descripcion", tarea.descripcion)
	}
}

func (t *taskList) imprimirListaCompletados() {
	for _, tarea := range t.tasks {
		if tarea.completado {
			fmt.Println("Nombre", tarea.nombre)
			fmt.Println("Descripcion", tarea.descripcion)
		}
	}
}

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
	t1 := &task{
		nombre:      "Completar mi curso de Go",
		descripcion: "Completar mi curso de Go de platzi esta semana",
	}

	t2 := &task{
		nombre:      "Completar mi curso de Python",
		descripcion: "Completar mi curso de Python de platzi esta semana",
	}

	t3 := &task{
		nombre:      "Completar mi curso de NodeJs",
		descripcion: "Completar mi curso de NodeJs de platzi esta semana",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}

	fmt.Println(lista.tasks[0])
	lista.agregarALista(t3)
	// fmt.Println(len(lista.tasks))
	// lista.eliminarDeList(1)
	// fmt.Println(len(lista.tasks))
	// // fmt.Println(lista.tasks[2])
	// for i := 0; i < len(lista.tasks); i++ {
	// 	fmt.Println("Index", i, "Nombre", lista.tasks[i].nombre)
	// }

	// for index, tarea := range lista.tasks {
	// 	fmt.Println("Index", index, "Nombre", tarea.nombre)
	// }

	lista.imprimirLista()
	lista.tasks[0].marcarCompleta()
	fmt.Println("Tareas Comletadas")
	lista.imprimirListaCompletados()

	mapaTareas := make(map[string]*taskList)

	mapaTareas["Jhon"] = lista

	t4 := &task{
		nombre:      "Completar mi curso de Java",
		descripcion: "Completar mi curso de Java de platzi esta semana",
	}

	t5 := &task{
		nombre:      "Completar mi curso de C#",
		descripcion: "Completar mi curso de C# de platzi esta semana",
	}

	lista2 := &taskList{
		tasks: []*task{
			t4, t5,
		},
	}

	mapaTareas["Ricardo"] = lista2

	fmt.Println("\nTareas de Jhon")
	mapaTareas["Jhon"].imprimirLista()

	fmt.Println("\nTareas de Ricardo")
	mapaTareas["Ricardo"].imprimirLista()

}
