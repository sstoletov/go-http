package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

var (
	users = map[int]User{
		1: {ID: 1, Name: "Alice", Age: 25},
		2: {ID: 2, Name: "Bob", Age: 30},
	}
	nextID = 3
	mu     sync.Mutex
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Go HTTP Server")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case http.MethodGet:
		getUsers(w, r)

	case http.MethodPost:
		createUser(w, r)

	case http.MethodDelete:
		deleteUser(w, r)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.URL.Query().Get("id")

	if id != "" {
		userID, err := strconv.Atoi(id)
		if err != nil {
			http.Error(w, "invalid id", http.StatusBadRequest)
			return
		}

		mu.Lock()
		user, ok := users[userID]
		mu.Unlock()

		if !ok {
			http.Error(w, "user not found", http.StatusNotFound)
			return
		}

		json.NewEncoder(w).Encode(user)
		return
	}

	mu.Lock()
	list := make([]User, 0, len(users))
	for _, user := range users {
		list = append(list, user)
	}
	mu.Unlock()

	json.NewEncoder(w).Encode(list)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "invalid json", http.StatusBadRequest)
		return
	}

	mu.Lock()
	user.ID = nextID
	nextID++
	users[user.ID] = user
	mu.Unlock()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	userID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	mu.Lock()
	defer mu.Unlock()

	if _, ok := users[userID]; !ok {
		http.Error(w, "user not found", http.StatusNotFound)
		return
	}

	delete(users, userID)

	fmt.Fprintf(w, "user %d deleted\n", userID)
}

func main() {
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/users", usersHandler)

	fmt.Println("Server started on :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
