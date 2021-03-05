package main

import "fmt"

var soma = func (a int, b int) int {
	return a + b
}

func main() {
	fmt.Println(soma(3,4))

	var sub = func (a int, b int) int {
		return a -b
	}

	fmt.Println(sub(2, 3))
}
