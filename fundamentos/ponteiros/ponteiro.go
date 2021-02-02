package main

import "fmt"

func main() {
	i := 1 // go não tem aritmética de ponteiro

	var p *int = nil // ponteiro p do tipo int com 64 bits de tamanho.

	p = &i // pegando o endereço da variável i e apontando pro ponteiro

	*p++ //  está apenas pegando o valor.

	i++

	fmt.Println(p, *p, i, &i)

}
