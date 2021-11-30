package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type student struct {
	ID    string
	Name  string
	Grade int
}

var data = []student{
	{"A001", "Andika", 21},
	{"B001", "Bagus", 22},
	{"C001", "Aji", 23},
	{"C002", "Ali", 23},
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "GET" {
		var result, err = json.Marshal(data)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)
		return
	}

	http.Error(w, "", http.StatusMethodNotAllowed)
}


func GetStudent(w http.ResponseWriter, r *http.Request) {
	if !BasicAuth(w, r) {
		return
	}

	if !PostOnly(w, r) {
		return
	}

	w.Header().Set("Content-Type", "application/json")
	var id = r.FormValue("id")
	var result []byte
	var err error

	for _, each := range data {
		if each.ID == id {
			result, err = json.Marshal(each)

			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			w.Write(result)
			return
		}
	}
	http.Error(w, "user not found", http.StatusNotFound)
	return

}

func BasicAuth(w http.ResponseWriter, r *http.Request) bool {
	username, password, ok := r.BasicAuth()
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("unauthorized"))
		return false
	}

	valid := (username == "admin") && (password == "12345")
	if !valid{
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("wrong username or password"))
		return false
	}

	return true
}

func AuthenticationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("unauthorized"))
			return
		}

		valid := (username == "admin") && (password == "12345")
		if !valid{
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("wrong username or password"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func PostOnly(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("method not allowed"))
		return false
	}

	return true
}

func main() {

	mux := http.NewServeMux()

	mux.Handle("/students", AuthenticationMiddleware(http.HandlerFunc(GetStudents)))
	mux.HandleFunc("/student", GetStudent)

	fmt.Println("starting web server at http://localhost:8080/")

	log.Panic(http.ListenAndServe(":8080", mux))
}