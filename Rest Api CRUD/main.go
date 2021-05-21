package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

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

func getTasks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(&tasks)
}

func getTask(w http.ResponseWriter, r *http.Request) {
	log.Println(r)
	identifier := mux.Vars(r)
	log.Println(identifier["id"])
	taskID, err := strconv.Atoi(identifier["id"])
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err != nil {
		log.Println("ID invalido")
		fmt.Fprintf(w, "ID invalido")
		return
	}
	log.Println("Task length:", len(tasks))
	if taskID > len(tasks) || taskID <= 0 {
		log.Println("ID fuera del rango")
		fmt.Fprintf(w, "ID fuera del rango")
		return
	}
	task := tasks[taskID-1]
	log.Println(task)
	json.NewEncoder(w).Encode(task)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	log.Println(w)
	log.Println(r)
	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(w, "Ingrese una tarea valida")
		return
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
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")

	http.ListenAndServe(":3000", router)
}
