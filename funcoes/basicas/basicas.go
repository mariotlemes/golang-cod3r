package main

import "fmt"

func imprimirValor() {
	fmt.Println("Esta função apenas imprime um valor")
}

func funcaoComParametroSemRetorno (parametro1 string, parametro2 string) {
	fmt.Printf("Função com parâmetro mas sem retorno: %s %s \n", parametro1, parametro2)
}

func funcaoSemParametroComRetorno() string {
	return "Retorno da função sem parâmetro mas como retorno"
}

func funcaoCompletaComParametrosDoMesmoTipo (parametro1, parametro2 string) string {
	return fmt.Sprintf("Parâmetros do mesmo tipo: %s %s", parametro1, parametro2)
}

func funcaoMultiplosRetornos () (string, string) {
	return "Retorno 1", "Retorno 2"
}

func main() {

	//f1 := imprimirValor()
	//f2 := funcaoComParametroSemRetorno("Teste", "Teste2")
	f3 := funcaoSemParametroComRetorno()
	fmt.Println(funcaoCompletaComParametrosDoMesmoTipo("Teste4", "Teste5"))

	fmt.Println(funcaoMultiplosRetornos())

	fmt.Println(f3)
}
