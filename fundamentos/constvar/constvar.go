package main

import (
	"fmt"
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.2 // tipo (float64)

	//forma reduzida declara e atribui conteúdo
	//variáveis em go declaradas PRECISAM ser usadas
	area := PI * m.Pow(raio, 2)
	fmt.Printf("O raio é %.2f\n", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(a, b, c, d)

	g, h, i := 2, "teste", false

	fmt.Println(g, h, i)




}
