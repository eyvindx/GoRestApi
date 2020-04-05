package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	r := chi.NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {

	})
	http.ListenAndServe(":3000", r)
	fmt.Println("Hola Mundo")
	fmt.Println("Hola Mundo")
	fmt.Println("Hola Mundo")
}
