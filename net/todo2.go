package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Todo struct {
	ID        int    `json:"id"`
	Task      string `json:"task"`
	Completed bool   `json:"completed"`
}

var (
	todos     = make([]Todo, 0)
	idCounter = 1
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/todos", getTodos).Methods("GET")
	r.HandleFunc("/todos", createTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", updateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", deleteTodo).Methods("DELETE")

	// 启动 HTTP 服务器
	http.ListenAndServe(":8080", cors(r))
}

// handler: 获取所有 ToDo
func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

// handler: 创建新的 ToDo
func createTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var t Todo
	json.NewDecoder(r.Body).Decode(&t)
	t.ID = idCounter
	idCounter++
	todos = append(todos, t)
	json.NewEncoder(w).Encode(t)
}

// handler: 更新指定 ID 的 ToDo
func updateTodo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	var updated Todo
	json.NewDecoder(r.Body).Decode(&updated)

	for i, t := range todos {
		if t.ID == id {
			todos[i].Task = updated.Task
			todos[i].Completed = updated.Completed
			json.NewEncoder(w).Encode(todos[i])
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

// handler: 删除指定 ID 的 ToDo
func deleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])
	for i, t := range todos {
		if t.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.Error(w, "Not Found", http.StatusNotFound)
}

// 简易 CORS 中间件
func cors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
