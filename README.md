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
    - [if-else-if](#if-else-if)
    - [if com Init](#if-com-init)
    - [For](#for)
      - [For com uma condição - parecido com if](#for-com-uma-condi%C3%A7%C3%A3o---parecido-com-if)
      - [For tradicional](#for-tradicional)
      - [For "infinito"](#for-infinito)
    - [Switch](#switch)
      - [Switch #01](#switch-01)
      - [Switch #02](#switch-02)
      - [Switch #03](#switch-03)
  - [Arrays](#arrays)
    - [Definição](#defini%C3%A7%C3%A3o)
    - [Atribuição múltipla](#atribui%C3%A7%C3%A3o-m%C3%BAltipla)
    - [Tamanho de um array e convertendo para um tipo](#tamanho-de-um-array-e-convertendo-para-um-tipo)
    - [Tamanho do array indefinido](#tamanho-do-array-indefinido)
    - [Percorrer um for usando range](#percorrer-um-for-usando-range)
      - [for para acessar índice e valor](#for-para-acessar-%C3%ADndice-e-valor)
      - [for ignorando o índice usando _](#for-ignorando-o-%C3%ADndice-usando-_)
    - [for ignorando os valores](#for-ignorando-os-valores)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# [Go Lang - Explorando a linguagem do Google](https://www.udemy.com/course/curso-go/)

Este repositório tem como principal objetivo fazer memória dos recursos de golang.

Data de início do aprendizado: **01/02/2021**.

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
  - Soma -> +
  - Subtração -> -
  - Multiplicação -> *
  - Módulo -> %
  - E -> &
  - Ou -> |

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
  - Estrutura:

```
    func nome_da_funcao (parâmetro tipo, ...) tipo_retorno {
        //lógica da função
    }
```

- Exemplo:

```
    func somar (a int, b int) int {
        return a + b
    }
```

  - função principal: func **main** () {}
  - função main() precisa estar definida no package main.


 ### Ponteiros 
  - Go não tem aritmética de ponteiros;
  - Como declarar -> var p *int = nil //este ponteiro é um inteiro e deve apontar para nil
  - i : = 1 // atribuição simplificada da variável i
  - p = &i // end. de memória da variável i é atribuída ao ponteiro p
  - *p //valor do ponteiro é retornado
    

## Estruturas de Controle

### if-else
   - Estrutura:
```
    if condição {
        //se condição for true;
    } else {
        //se condição for false;
    }
```
   - else não é obrigatório;
   - a condição não é cercada por ().
   - tanto o if quanto else precisam estar delimitados por {}

### if-else-if

- Estrutura:
```
    if condição 1 {
        //se condição 1 for true;
    } else if condicao 2 {
        //se condição 2 for true;
    } else if condicao n {
        // se condição n for true;
    } else {
        // todas anteriores são False
    }
```
### if com Init
- Ideia: Trazer a inicialização de uma
variável da estrutura de controle for para o if.
Portanto, é possível criar uma variável que
  receba uma valor e que só funciona dentro do if.

- Estrutura:
```
    if i := 2; i == 2 {
      fmt.Println("A variavél i vale 2")
    } else {
      fmt.Println("A variável i não é igual a 2")
    }
```
Obs: No lugar de uma atribuição estática pode-se 
utilizar uma função previamente definida.
### For
#### For com uma condição - parecido com if
-Exemplo:
```
i := 1
for i <= 10 {
  fmt.Println(i)
  i++
}
```
#### For tradicional
- Exemplo:
```
for i := 0; i <=20; i += 2 {
  fmt.Printf("%d ", i)
}
```
#### For "infinito"
- Exemplo:

```
for {
  fmt.Println("Para sempre!")
  time.Sleep(time.Second * 2)
}
```
Obs: A **única** estrutura de 
repetição em go é o **for**, não
existe **while** e suas variantes.
### Switch
- Usado para múltiplas seleções. Diferentemente
  do if não testa uma condição. É apenas uma
  expressão associada e que vai ser testada
  em um dos cases.
- Por padrão, em go o switch entra em um case
  e depois sai da estrutura switch.
- Palavra **fallthough** significa passe
  para o próximo case.
- É possível colocar vários cases em apenas
  um case (Ex: case 4, 5).
- Caso não entre em nenhum case, o case
  **default** é executado.

#### Switch #01
- Exemplo:

```
switch nota {
  case 10:
	 fallthrough
  case 9:
	return "A"
  case 8, 7:
	return "B"
  case 6, 5:
  	return "C"
  case 4, 3:
	return "D"
  default:
	return "Nota inválida!"
}
```

#### Switch #02
- Switch sem nenhuma expressão associada.
- Neste caso, procura-se pelo case 
  verdadeiro.
- Exemplo:
```
t := time.Now()
switch { //switch true
case t.Hour() < 12:
  fmt.Println("Bom dia!")
case t.Hour() < 18:
  fmt.Println("Boa tarde!")
default:
  fmt.Println("Boa noite!")
}
```

#### Switch #03
- O tipo **interface** é 
  genérico.
Posso usar switch para verificar qual
  é o tipo do parâmetro passado por
  parâmetro.
  
- Exemplo:
```
func tipo(i interface{}) string {
  switch i.(type) {
  case int:
    return "Inteiro"
  case float32, float64:
    return "float"
  case string:
	return "string"
  case func():
	return "função"
  default:
  	return "n sei"
  }
}
```
## Arrays
- Estruturas ~~~~fixas e homogêneas (mesmo 
  tipo).
  
- Não precisa ser inicializado. Por 
padrão, os arrays são criados e inicializados 
  com zeros do tipo definido.
### Definição
- nome_array [tamanho] tipo
- Exemplo:
```
var notas [3] float64
```
### Atribuição múltipla
- Exemplo:
```
notas[0], notas[1], notas[2] = 3.3, 1.2, 4.4
```

### Tamanho de um array e convertendo para um tipo
- Exemplo:
```
float64(len(notas))
```

### Tamanho do array indefinido
- Exemplo:
```
numeros := [...] int {1,2,3,4,5}
```
Obs: O compilador que vai contar. Neste caso o
array **numeros** é composto por
5 inteiros.

### Percorrer um for usando range
#### for para acessar índice e valor
```
for i, numero := range numeros {
  fmt.Printf("%d) %d\n", i, numero)
}
```

#### for ignorando o índice usando _
```
for _, numeros := range numeros {
  fmt.Println(numeros)
}
```

### for ignorando os valores
```
for numeros := range numeros {
  fmt.Println(numeros) //serão impressos os indices
}
```
