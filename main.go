package main

import (
	"fmt"
	"log"
	"net/http"
)

func holaMundoHandler(w http.ResponseWriter, r *http.Request) {
	// Solo permitir método GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Establecer el tipo de contenido
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Escribir la respuesta
	fmt.Fprintf(w, "hola mundo")
}

func postHandler(w http.ResponseWriter, r *http.Request) {
	// Solo permitir método POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Establecer el tipo de contenido
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Escribir la respuesta
	fmt.Fprintf(w, "Respuesta del handler POST")
}

func main() {
	// Registrar los handlers
	http.HandleFunc("/", holaMundoHandler)
	http.HandleFunc("/post", postHandler)

	// Iniciar el servidor en el puerto 8080
	port := ":8080"
	fmt.Printf("Servidor iniciado en http://localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
