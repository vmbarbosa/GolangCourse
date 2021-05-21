package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/gorilla/mux"
)

func Test_indexRoute(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	request, _ := http.NewRequest("GET", "", nil)

	tests := []struct {
		name string
		args args
	}{
		{
			name: "test 1",
			args: args{
				w: httptest.NewRecorder(),
				r: request,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			indexRoute(tt.args.w, tt.args.r)
		})
	}
}

func Test_getTasks(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	request, _ := http.NewRequest("GET", "tasks", nil)
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test 1",
			args: args{
				w: httptest.NewRecorder(),
				r: request,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getTasks(tt.args.w, tt.args.r)
		})
	}
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tasks", getTasks).Methods("GET")
	router.HandleFunc("/tasks", createTask).Methods("POST")
	router.HandleFunc("/tasks/{id}", getTask).Methods("GET")
	return router
}

func Test_getTask(t *testing.T) {

	//Test 1

	request, _ := http.NewRequest("GET", "/tasks/0", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	respuesta, _ := ioutil.ReadAll(response.Body)

	if string(respuesta) != "ID fuera del rango" {
		t.Errorf("Test 1 - getTasks() = %v, want %v", response.Body, "ID fuera del rango")
	}

	//Test 2

	request, _ = http.NewRequest("GET", "/tasks/p", nil)
	response = httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	respuesta, _ = ioutil.ReadAll(response.Body)

	if string(respuesta) != "ID invalido" {
		t.Errorf("Test 2 - getTasks() = %v, want %v", response.Body, "ID invalido")
	}

	//Test 3

	request, _ = http.NewRequest("GET", "/tasks/1", nil)
	response = httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	respuesta, _ = ioutil.ReadAll(response.Body)

	task_test := task{
		ID:      1,
		Name:    "Tarea Uno",
		Content: "Prueba 01",
	}

	task_print, _ := json.Marshal(task_test)
	r := regexp.MustCompile("\\s+")
	replace := r.ReplaceAllString(string(respuesta), " ")
	if replace != (string(task_print) + " ") {
		t.Errorf("Test 3 - getTasks() = %v, want %v", replace, string(task_print))
	}

}

func Test_createTask(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			createTask(tt.args.w, tt.args.r)
		})
	}
}

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
