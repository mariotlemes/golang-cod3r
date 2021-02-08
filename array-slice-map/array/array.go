package main

import "fmt"

func main() {
	// arrays são homogêneos e estáticos (fixo)
	var notas [3] float64
	fmt.Println(notas)
	notas[0], notas[1], notas[2] = 3.6, 1.2, 4.4
	fmt.Print(notas)

	total := 0.0
	for i:=0; i < len(notas); i++ {
		total += notas[i]
	}

	//fmt.Print(total)

	media := total / float64(len(notas))

	fmt.Printf("\nMedia: %.2f", media)



}
