package main

import (
	"fmt"
	"log"
	"net/http"
)

func handleGet(w http.ResponseWriter, r *http.Request) {
	// Verifica que el método sea GET
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	// Responde con un mensaje simple
	fmt.Fprintln(w, "¡Hola, has realizado un GET!")
}

func main() {
	// Asigna la función handleGet a la ruta raíz
	http.HandleFunc("/", handleGet)

	fmt.Println("Servidor escuchando en http://localhost:8080")
	// Inicia el servidor en el puerto 8080
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
