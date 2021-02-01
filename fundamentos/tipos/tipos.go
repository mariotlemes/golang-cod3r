package main

import (
	"fmt"
	"math"
	"reflect"
)

func main() {
	fmt.Println(1, 2, 1000)
	fmt. Println("Literal inteiro é: ", reflect.TypeOf(32000))

	var b uint8 = 255
	fmt.Println("O byte é", reflect.TypeOf(b))

	maxValeMaxInt64 := math.MaxInt64
	fmt.Println("O máximo valor int64 é", maxValeMaxInt64)
	fmt.Printf("O tipo do valor %v é %v", maxValeMaxInt64, reflect.TypeOf(maxValeMaxInt64))

	var i2 rune = 'a'
	fmt.Println("O rune é", reflect.TypeOf(i2))
	fmt.Println("O valor da variável i2 é:", i2)

	bo := true
	fmt.Println(reflect.TypeOf(bo))
	fmt.Println(!bo)

	s1 := "O meu nome é Mario Teixeira Lemes!"
	fmt.Println(s1 + "!!")
	fmt.Println("O tamanho da string é:", len(s1))

	//string com múltiplas linhas

	s2 := `teste
	mario
	junior
	teste `

	fmt.Println("A variável s2 possui o tamanho:", len(s2))


}
