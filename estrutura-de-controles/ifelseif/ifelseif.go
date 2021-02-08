package main

import "fmt"

func notaParaConceito(nota float64) string {
	if nota >= 9 && nota <= 10 {
		return "A"
	} else if nota >= 7 && nota < 9 {
		return "B"
	} else if nota < 7 && nota >= 5 {
		return "C"
	} else {
		return "D"
	}
}

func main() {
	fmt.Println(notaParaConceito(9.8))
	fmt.Println(notaParaConceito(3.2))
	fmt.Println(notaParaConceito(6.4))
	fmt.Println(notaParaConceito(7.2))
}
