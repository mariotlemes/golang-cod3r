package main

import "fmt"

func main() {
	var b byte = 3
	fmt.Println("Valor de b é:", b)

	i := 3
	i += 3 // i = i +3
	i -= 3
	i *= 2
	i %= 2

	fmt.Println(i)

	x, y := 1, 2

	fmt.Println(x, y)

	x, y = y, x // trocar o valor das variáveis

	fmt.Println(x, y)

}
