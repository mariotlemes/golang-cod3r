package main

import "fmt"

func main() {
	fmt.Print("Mesma")
	fmt.Print(" linha")
	fmt.Println("Nova")
	fmt.Println("Linha")

	x := 3.1491516
	xstring := fmt.Sprint(x)
	fmt.Println("O valor de x é " + xstring)

	fmt.Printf("O valor de x formatado é %.2f.", x)

	a := 1
	b:= 1.869
	c := false
	d := "opa"

	fmt.Printf("\n %d, %.2f, %t, %s", a, b, c, d)

	fmt.Printf("\n %v, %.2v, %v, %v\n", a, b, c, d)

}
