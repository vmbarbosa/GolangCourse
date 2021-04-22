package main

import "fmt"

func main() {
	t := task{
		nombre:      "Crear pipeline en Git Actions",
		descripcion: "Crear pipeline de Golang para validación de pruebas unitarias de Golang",
		completado:  false,
	}
	fmt.Printf("%+v\n", t)

	t.actualizarNombre("Completar curso")

	fmt.Printf("%+v\n", t)
}
