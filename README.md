<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Sumário**

- [Go Lang - Explorando a linguagem do Google](#go-lang---explorando-a-linguagem-do-google)
  - [Fundamentos](#fundamentos)
    - [Comentários gerais](#coment%C3%A1rios-gerais)
    - [Comandos úteis](#comandos-%C3%BAteis)
    - [Tipos](#tipos)
    - [Valores zeros (padrão)](#valores-zeros-padr%C3%A3o)
    - [Operadores](#operadores)
      - [Aritméticos](#aritm%C3%A9ticos)
      - [Lógicos](#l%C3%B3gicos)
      - [Relacionais](#relacionais)
    - [Funções](#fun%C3%A7%C3%B5es)
    - [Ponteiros](#ponteiros)
  - [Estruturas de Controle](#estruturas-de-controle)
    - [if-else](#if-else)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# [Go Lang - Explorando a linguagem do Google](https://www.udemy.com/course/curso-go/)

Este repositório tem como principal objetivo fazer memória dos recursos de golang.

Data de início do aprendizado: 01/02/2021.

## Fundamentos

### Comentários gerais
  - comentários podem ser feitos por /* */ ou //

### Comandos úteis
  - $ got vet arquivo.go //faz uma verificação no arquivo.go em busca de inconsistências
  - $ go build arquivo.go //compila o arquivo.go
  - $ ./arquivo //executa arquivo
  - $ go run arquivo.go //compila e executa 
 
### Tipos
  - reflect.TypeOf(32000) //método para retornar o tipo de uma variável/valor
  - var i2 rune = "a" // comando rune mostra o valor da tabela ASCII. Se printarmos o valor de i2 o resultado é 97 (valor ascii para o caracterere 'a').
  - valueMaxInt64 = max.MaxInt64 // método MaxInt64 da biblioteca math retorna o maior valor inteiro possível com 64 bits.
  
### Valores zeros (padrão)
  - variáveis float e int64 quando criadas mas não inicializadas com valor possui o valor padrão de 0.
  - variável bool possui o valor padrão false.
  - ponteiro possui o valor zero como **nil**. O valor nil é o mesmo que null de outras linguagens, referindo-se a vazio.
  
 ### Conversões
  - strconv.Itoa(123) //converte  int para string
  - strconv.Atoi("123") //converte string para int
  - num, _ := strconv.Atoi("123") //variável num caso a conversão tenha sucesso e _ caso dê erro. Como declarei o símbolo _ no lugar de uma variável para tratar o erro, consigo executar o código mesmo não tratando o erro posteriormente. Caso utilize uma variável err no lugar do símbolo *_*, necessariamente precisarei tratar o erro.

 ### Atribuição
  - i := 3 // atribuição resumida - declara e atribui valor.
  - i += 3 // i = i + 3 
  - i -= 3
  - i *= 2
  - i %= 2

### Operadores 

#### Aritméticos

	a := 3
	b := 2

	fmt.Println("Soma:", a+b)
	fmt.Println("Subtração:", a-b)
	fmt.Println("Multiplicação:", a*b)
	fmt.Println("Módulo:", a%b)

	fmt.Println("E:", a&b)
	fmt.Println("Ou:", a|b)
	fmt.Println("Xor:", a^b)

#### Lógicos
 - && -> e;
 - | -> ou;
 - != -> diferente (! = )
 - ! -> negação

#### Relacionais
 - igual -> == ou d1.Equal(d2)
 - menor -> <
 - menor ou igual ->  < =
 -  maior ou igual -> < =

### Funções
  - func **nome_da_funcao** (parâmetro **tipo**, ...) **tipo_retorno** {}
  - função executada: func **main** () {}
  - função main() precisa estar definida no package main.


 ### Ponteiros 
  - Go não tem aritmética de ponteiros;
  - Como declarar -> var p *int = nil //este ponteiro é um inteiro e deve apontar para nil
  - i : = 1
    p = &i // end. de memória da variável i é atribuída ao ponteiro p
    *p //valor do ponteiro é retornado
    

## Estruturas de Controle

### if-else
   - Estrutura:

    if condição {
        //se condição for true;
    } else {
        //se condição for false;
    }

   - else não é obrigatório;
   - a condição não é cercada por ().
   - tanto o if quanto else precisam estar delimitados por {}
