package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:"ID"`
	Name    string `json:"Name"`
	Content string `json:"Content"`
}

type allTasks []task

var tasks = allTasks{
	{
		ID:      1,
		Name:    "Tarea Uno",
		Content: "Prueba 01",
	},
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Bienvenido hola otra vez")
}

func getTask(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Ingrese una tarea valida")
	}

	json.Unmarshal(reqBody, &newTask)

	newTask.ID = len(tasks) + 1

	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(newTask)
}

func main() {
	// Crear enrutador
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTask).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")

	http.ListenAndServe(":3000", router)
}
