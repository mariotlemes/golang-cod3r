package main

import "fmt"

func notaParaConceito(nota float64) string {
	var nota2 = int(nota)
	switch nota2 {
	case 10:
		fallthrough
	case 9:
		return "A"
	case 8, 7:
		return "B"
	case 6, 5:
		return "C"
	case 4, 3:
		return "D"
	default:
		return "Nota inválida!"
	}
}

func main() {
	fmt.Println(notaParaConceito(11))
}