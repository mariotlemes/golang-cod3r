package main

import "fmt"

func main() {
	a1 := [3] int {1,2,3} //array tem tamanho definido

	s1 := [] int {1, 2, 3} //slice não tem tamanho definido. É apenas um pedaço de um array
	fmt.Println(a1)
	fmt.Println(s1)

	a2 := [5] int {1,2, 3, 4, 5}

	s2 := a2 [1:3]

	fmt.Println(a2, s2)



}
