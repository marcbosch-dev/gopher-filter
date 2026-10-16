package main

import (
	"bufio"
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

	if len(os.Args) != 2 {
		fmt.Println("Error: You must provide a JSON file")
		os.Exit(1)
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error: Couldn't not read file ", os.Args[1])
		os.Exit(1)
	}
	defer file.Close()

	var usuarios []Usuario
	decorder := json.NewDecoder(file)
	err = decorder.Decode(&usuarios)
	if err != nil {
		fmt.Println(err)
		fmt.Print("Error: la estructura de los datos de JSON es incorrecta")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
filter:
	for {
		showMenu()
		if !scanner.Scan() {
			fmt.Print("Error: could not read the Stdin")
			break filter
		}
		option := scanner.Text()

		if err := scanner.Err(); err != nil {
			fmt.Println("Error: ", err)
			break filter
		}

		switch option {
		case "1":

		case "6":
			break filter
		}
	}
}

func showMenu() {
	fmt.Println("Choise Atributte:")
	fmt.Println("1. Id")
	fmt.Println("2. Name")
	fmt.Println("3. Email")
	fmt.Println("4. Salary")
	fmt.Println("5. State")
	fmt.Println("6. Exit.")
}
