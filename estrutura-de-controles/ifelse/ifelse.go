package main

import "fmt"

func imprimirResultado (nota float64) {
	if nota >= 7 {
		fmt.Println("Aprovado:", nota)
	} else {
		fmt.Println("Reprovado:", nota)
	}
}

func main() {
	imprimirResultado(3.1)
	imprimirResultado(7.2)
}


