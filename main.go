package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// JSON  struct
type Usuario struct {
	Id     int     `json:"id"`
	Nombre string  `json:"nombre"`
	Email  string  `json:"email"`
	Monto  float64 `json:"monto"`
	Estado string  `json:"estado"`
}

func main() {
	fmt.Println("=== Gopher-filter ===")

	fmt.Println(os.Args[1])
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: Couldn't not read file ", os.Args[1])
		os.Exit(1)
	}
	defer file.Close()

	var usuarios []Usuario
	decorder := json.NewDecoder(file)
	// Decodofica para ver que todo esté bien
	err = decorder.Decode(&usuarios)
	if err != nil {
		fmt.Println(err)
		fmt.Print("Error: la estructura de los datos de JSON es incorrecta")
		os.Exit(1)
	}

	fmt.Printf("Se leyeron %d registros\n", len(usuarios))
	for _, u := range usuarios {
		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Monto: %.2f, Estado: %s\n",
			u.Id, u.Nombre, u.Email, u.Monto, u.Estado)
	}
}
