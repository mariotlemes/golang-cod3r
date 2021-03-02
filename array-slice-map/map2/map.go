package main

import "fmt"

func main() {
	listaFuncionarios := map[string]float64 {
		"José João": 11242.56,
		"Gabriel Junior": 1500.00,
		"Maria Joana": 1200.00,
	}

	listaFuncionarios["Rafael Filho"] = 1350.0

	fmt.Println(listaFuncionarios)

	delete(listaFuncionarios, "Inexistente")

	for nome, salario := range listaFuncionarios {
		fmt.Printf("Nome: %s Salário: R$ %.2f\n", nome, salario)
	}
}