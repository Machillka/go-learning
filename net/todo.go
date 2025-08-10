package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

type TodoItem struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

var (
	todoRepo = make(map[int]TodoItem)
	nextId   = 1
	mu       sync.RWMutex
)

func main() {
	mux := http.NewServeMux()
	// 注册 /todos 用于处理 todoitem
	mux.HandleFunc("/todos", TodosHandler)
	mux.HandleFunc("/todos/", TodoHandler) // 处理 /todos/{id}，支持 GET、PUT、DELETE

	addr := ":8080"
	log.Printf("Server listening on %s\n", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("Server start error: %v", err)
	}
}

// 对于整个TodoRepo的处理
func TodosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		mu.RLock()
		list := make([]TodoItem, 0, len(todoRepo))

		for _, t := range todoRepo {
			list = append(list, t)
		}
		mu.RUnlock()

		json.NewEncoder(w).Encode(list)
	case http.MethodPost:
		var t TodoItem

		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, "无效的 Json", http.StatusBadRequest)
			return
		}

		mu.Lock()
		t.Id = nextId
		nextId += 1
		todoRepo[t.Id] = t
		mu.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(t)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

func TodoHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	idStr := strings.TrimPrefix(r.URL.Path, "/todos/")

	id, err := strconv.Atoi(idStr)

	if err != nil {
		http.Error(w, "Error id", http.StatusBadRequest)
		return
	}

	switch r.Method {
	case http.MethodGet:
		mu.RLock()
		t, ok := todoRepo[id]
		mu.RUnlock()

		if !ok {
			http.Error(w, "未找到该 ToDo", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(t)
	case http.MethodPut:
		// 更新指定 Todo
		var t TodoItem
		if err := json.NewDecoder(r.Body).Decode(&t); err != nil {
			http.Error(w, "无效的 JSON", http.StatusBadRequest)
			return
		}

		mu.Lock()
		if _, ok := todoRepo[id]; !ok {
			mu.Unlock()
			http.Error(w, "未找到该 ToDo", http.StatusNotFound)
			return
		}
		t.Id = id
		todoRepo[id] = t
		mu.Unlock()

		json.NewEncoder(w).Encode(t)
	case http.MethodDelete:
		mu.Lock()
		if _, ok := todoRepo[id]; !ok {
			mu.Unlock()
			http.Error(w, "未找到该 ToDo", http.StatusNotFound)
			return
		}
		delete(todoRepo, id)
		mu.Unlock()

		w.WriteHeader(http.StatusNoContent)
	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}
