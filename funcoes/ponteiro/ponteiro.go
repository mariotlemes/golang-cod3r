package main

import "fmt"

func inc1(n int) {
	n++
}

// ponteiro é representado por *
func inc2 (n *int) {
	//revisão: * é usado para desreferenciar. Valor a qual o ponteiro aponta.
	*n++
}

func main() {
	n := 1
	inc1(n) //por valor
	fmt.Println(n)

	// revião: & usado para obter o endereço
	inc2(&n) // por referência
	fmt.Println(n)
}

