package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello,baby!!!")
}
func byeHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Bye, baby(")
}
func idHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	fmt.Fprintf(w, "ID пользователя: %s", id)
}
func sumHandler(w http.ResponseWriter, r *http.Request) {

	a := r.URL.Query().Get("a")
	b := r.URL.Query().Get("b")
	aInt, err1 := strconv.Atoi(a)
	if err1 != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return

	}
	bInt, err2 := strconv.Atoi(b)
	if err2 != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return

	}

	sum := aInt + bInt

	fmt.Fprintf(w, "%d", sum)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello", helloHandler)
	mux.HandleFunc("GET /bye", byeHandler)
	mux.HandleFunc("GET /sum", sumHandler)
	mux.HandleFunc("GET /users/{id}", idHandler)

	http.ListenAndServe(":8080", mux)

}
