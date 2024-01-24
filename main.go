package main

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

type Todo struct {
	Name       string
	IsComplete bool
}

var todos = []Todo{
	{"Wash the car", false},
	{"Get groceries", true},
	{"Buy movie tickets", false},
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/delete/", handleDelete)
	http.HandleFunc("/complete/", handleComplete)
	http.HandleFunc("/add", handleAdd)

	fmt.Println("Listening on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	home := home("todo.List()", todos)
	home.Render(r.Context(), w)
}

func handleDelete(w http.ResponseWriter, r *http.Request) {
	index := indexFromPath(r.URL.Path)

	var newTodos []Todo
	for i := 0; i < len(todos); i++ {
		if i != index {
			newTodos = append(newTodos, todos[i])
		}
	}
	todos = newTodos

	todoList := todoList(todos)
	todoList.Render(r.Context(), w)
}

func handleComplete(w http.ResponseWriter, r *http.Request) {
	index := indexFromPath(r.URL.Path)

	todos[index].IsComplete = !todos[index].IsComplete

	status := todoStatus(todos[index].IsComplete, index)
	status.Render(r.Context(), w)
}

func handleAdd(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	todoName := r.FormValue("name")

	if todoName != "" {
		todos = append(todos, Todo{todoName, false})
	}

	todoList := todoList(todos)
	todoList.Render(r.Context(), w)
}

func indexFromPath(rawPath string) int {
	path := strings.Split(rawPath, "/")
	if len(path) != 3 {
		panic("incorrectly formed URL path")
	}

	index, err := strconv.Atoi(path[2])
	if err != nil {
		panic(err)
	}

	return index
}
