package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Sistema de Gestión de Libros Electrónicos")
	})
	fmt.Println("Servidor iniciado en puerto 8080")
	http.ListenAndServe(":8080", r)
}
