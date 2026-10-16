package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	// filtrados empieza con todos los usuarios y se reduce con cada filtro
	filtrados := usuarios
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
		case "1": // Id: puede buscar por texto (contains) o por limite (below)
			fmt.Println("Filter Id by:")
			fmt.Println("1. Contains (text)")
			fmt.Println("2. Below (number)")
			sub, err := readLine(scanner)
			if err != nil {
				fmt.Println("Error: ", err)
				break filter
			}
			switch sub {
			case "1":
				fmt.Print("Text to search: ")
				texto, err := readLine(scanner)
				if err != nil {
					fmt.Println("Error: ", err)
					break filter
				}
				filtrados = filter(filtrados, func(u Usuario) bool {
					return contains(strconv.Itoa(u.Id), texto)
				})
				printUsuarios(filtrados)
			case "2":
				fmt.Print("Limit: ")
				linea, err := readLine(scanner)
				if err != nil {
					fmt.Println("Error: ", err)
					break filter
				}
				limite, err := strconv.ParseFloat(linea, 64)
				if err != nil {
					fmt.Println("Error: the limit must be a number")
					continue
				}
				filtrados = filter(filtrados, func(u Usuario) bool {
					return below(float64(u.Id), limite)
				})
				printUsuarios(filtrados)
			}
		case "2": // Name: contains
			fmt.Print("Text to search: ")
			texto, err := readLine(scanner)
			if err != nil {
				fmt.Println("Error: ", err)
				break filter
			}
			filtrados = filter(filtrados, func(u Usuario) bool {
				return contains(u.Nombre, texto)
			})
			printUsuarios(filtrados)
		case "3": // Email: contains
			fmt.Print("Text to search: ")
			texto, err := readLine(scanner)
			if err != nil {
				fmt.Println("Error: ", err)
				break filter
			}
			filtrados = filter(filtrados, func(u Usuario) bool {
				return contains(u.Email, texto)
			})
			printUsuarios(filtrados)
		case "4": // Salary: below
			fmt.Print("Limit: ")
			linea, err := readLine(scanner)
			if err != nil {
				fmt.Println("Error: ", err)
				break filter
			}
			limite, err := strconv.ParseFloat(linea, 64)
			if err != nil {
				fmt.Println("Error: the limit must be a number")
				continue
			}
			filtrados = filter(filtrados, func(u Usuario) bool {
				return below(u.Monto, limite)
			})
			printUsuarios(filtrados)
		case "5": // State: isActive
			fmt.Println("Filtrado por usuarios Activos")
			filtrados = filter(filtrados, func(u Usuario) bool {
				return isActive(u.Estado)
			})
			printUsuarios(filtrados)
		case "6": // Continue: modificar usuarios de la lista filtrada
			modifyUser(scanner, filtrados)
		case "7":
			fmt.Println("=== Filtered list ===")
			printUsuarios(filtrados)
			break filter
		}
	}
}

// modifyUser permite cambiar un atributo individual de un usuario de la
// lista filtrada. Aplica las reglas de transformación del README:
// normaliza textos (emails y estados en minúsculas) y cada validación
// puede fallar individualmente sin detener el programa.
func modifyUser(scanner *bufio.Scanner, filtrados []Usuario) {
	if len(filtrados) == 0 {
		fmt.Println("There are no users to modify.")
		return
	}

	fmt.Println("=== Modify user ===")
	printUsuarios(filtrados)

	fmt.Print("Id of the user to modify: ")
	linea, err := readLine(scanner)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	id, err := strconv.Atoi(linea)
	if err != nil {
		fmt.Println("Error: the id must be a number")
		return
	}

	// buscar el indice del usuario dentro de la lista filtrada
	indice := -1
	for i, u := range filtrados {
		if u.Id == id {
			indice = i
			break
		}
	}
	if indice == -1 {
		fmt.Println("Error: user with id", id, "not found in the filtered list")
		return
	}

	fmt.Println("Attribute to modify:")
	fmt.Println("1. Name")
	fmt.Println("2. Email")
	fmt.Println("3. Salary")
	fmt.Println("4. State")
	opcion, err := readLine(scanner)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	fmt.Print("New value: ")
	valor, err := readLine(scanner)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	// cada regla valida por separado; si falla, solo se omite la modificacion
	switch opcion {
	case "1": // Name: normalizar texto (quitar espacios)
		nombre := strings.TrimSpace(valor)
		if nombre == "" {
			fmt.Println("Error: the name cannot be empty. User not modified.")
			return
		}
		filtrados[indice].Nombre = nombre
	case "2": // Email: validar formato y normalizar a minusculas
		email := strings.ToLower(strings.TrimSpace(valor))
		if !strings.Contains(email, "@") || !strings.Contains(email, ".") {
			fmt.Println("Error: invalid email format. User not modified.")
			return
		}
		filtrados[indice].Email = email
	case "3": // Salary: debe ser numerico y no negativo
		salario, err := strconv.ParseFloat(valor, 64)
		if err != nil {
			fmt.Println("Error: the salary must be a number. User not modified.")
			return
		}
		if salario < 0 {
			fmt.Println("Error: the salary cannot be negative. User not modified.")
			return
		}
		filtrados[indice].Monto = salario
	case "4": // State: debe ser activo o inactivo
		estado := strings.ToLower(strings.TrimSpace(valor))
		if estado != "activo" && estado != "inactivo" {
			fmt.Println("Error: the state must be 'activo' or 'inactivo'. User not modified.")
			return
		}
		filtrados[indice].Estado = estado
	default:
		fmt.Println("Error: invalid option")
		return
	}

	fmt.Println("User updated:")
	fmt.Printf("Id: %d | Name: %s | Email: %s | Salary: %.2f | State: %s\n",
		filtrados[indice].Id, filtrados[indice].Nombre, filtrados[indice].Email,
		filtrados[indice].Monto, filtrados[indice].Estado)
}

// readLine lee una linea del scanner. Retorna error si no hay mas entrada.
func readLine(scanner *bufio.Scanner) (string, error) {
	if scanner.Scan() {
		return scanner.Text(), nil
	}
	if err := scanner.Err(); err != nil {
		return "", err
	}
	return "", fmt.Errorf("no more input")
}

func printUsuarios(usuarios []Usuario) {
	if len(usuarios) == 0 {
		fmt.Println("No users found.")
		return
	}
	for _, u := range usuarios {
		fmt.Printf("Id: %d | Name: %s | Email: %s | Salary: %.2f | State: %s\n",
			u.Id, u.Nombre, u.Email, u.Monto, u.Estado)
	}
}

// filter aplica la función condición a cada usuario y devuelve un nuevo
// slice solo con aquellos para los que la condición retorna true.
func filter(usuarios []Usuario, condicion func(Usuario) bool) []Usuario {
	resultado := []Usuario{}
	for _, u := range usuarios {
		if condicion(u) {
			resultado = append(resultado, u)
		}
	}
	return resultado
}

func showMenu() {
	fmt.Println("Choise Atributte:")
	fmt.Println("1. Id")
	fmt.Println("2. Name")
	fmt.Println("3. Email")
	fmt.Println("4. Salary")
	fmt.Println("5. State")
	fmt.Println("6. Continue (modify users)")
	fmt.Println("7. Exit.")
}

// contains indica si el string `buscado` aparece dentro de `texto`
// sin tener en cuenta mayúsculas/minúsculas.
func contains(texto, buscado string) bool {
	return strings.Contains(strings.ToLower(texto), strings.ToLower(buscado))
}

// below indica si `valor` es menor que `limite`.
func below(valor, limite float64) bool {
	return valor < limite
}

// isActive indica si el string es "activo", sin distinguir mayúsculas/minúsculas.
func isActive(estado string) bool {
	return strings.EqualFold(estado, "activo")
}
