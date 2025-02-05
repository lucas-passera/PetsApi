package main

import (
	"PetsApi/database"
	"fmt"
)

func main() {
	// Llama a la función para conectar a la base de datos
	database.ConnectDatabase()

	// Verifica si la conexión fue exitosaaaa
	fmt.Println("La conexión a MySQL fue exitosa.")
}
