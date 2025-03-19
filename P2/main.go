package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// Configuración de la conexión a la base de datos
const (
	host     = "localhost" // O el nombre del contenedor si está en Docker
	port     = 5432
	user     = "admin"
	password = "admin"
	dbname   = "mi_base"
)

var db *sql.DB

// Función para manejar solicitudes GET en la raíz
func handleGet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "¡Hola, has realizado un GET!")
}

// Función para manejar una consulta a la base de datos
func handleDBQuery(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	// Ejecuta una consulta (ejemplo: obtener la fecha actual en PostgreSQL)
	var currentTime string
	err := db.QueryRow("SELECT NOW()").Scan(&currentTime)
	if err != nil {
		http.Error(w, "Error al consultar la base de datos", http.StatusInternalServerError)
		log.Println("Error en la consulta:", err)
		return
	}

	// Responde con la fecha y hora actual de PostgreSQL
	fmt.Fprintf(w, "Hora actual en la base de datos: %s\n", currentTime)
}

func main() {
	// Conexión a la base de datos
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error al abrir la conexión a la base de datos:", err)
	}
	defer db.Close()

	// Verificar la conexión a la base de datos
	err = db.Ping()
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	fmt.Println("¡Conectado a PostgreSQL!")

	// Rutas del servidor
	http.HandleFunc("/", handleGet)
	http.HandleFunc("/db", handleDBQuery) // Nueva ruta para la base de datos

	fmt.Println("Servidor escuchando en http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
