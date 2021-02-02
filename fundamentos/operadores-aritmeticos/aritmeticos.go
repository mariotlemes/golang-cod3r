package main

import (
	"fmt"
	"math"
)

func main() {
	a := 3
	b := 2

	fmt.Println("Soma:", a+b)
	fmt.Println("Subtração:", a-b)
	fmt.Println("Multiplicação:", a*b)
	fmt.Println("Módulo:", a%b)

	fmt.Println("E:", a&b)
	fmt.Println("Ou:", a|b)
	fmt.Println("Xor:", a^b)

	c := 3.0
	d := 2.0

	//outras operações

	fmt.Println("Maior:", math.Max(float64(a), float64(b)))

	fmt.Println("Menor:", math.Min(float64(a), float64(b)))

	fmt.Println("Exponenciação:", math.Pow(float64(c), float64(d)))

}
