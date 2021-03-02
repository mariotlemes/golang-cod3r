package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64 {
		"G": {
			"Gabriela Silva": 14535.32,
			"Guga Pereira": 12323.33,
		},

		"J": {
			"José Alves": 12323.12,
			"João José": 1231.22,
		},
		"P": {
			"Pedro S": 123.22,
			"Pato": 1212.23,
		},
	}

	fmt.Println(funcsPorLetra)

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}





}