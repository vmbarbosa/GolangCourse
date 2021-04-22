package main

func (t task) completarTarea() {
	t.completado = true
}

func (t task) actualizarNombre(nombre string) {
	t.nombre = nombre
}
