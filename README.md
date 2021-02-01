<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**  *generated with [DocToc](https://github.com/thlorenz/doctoc)*

- [golang-cod3r.com.br](#golang-cod3rcombr)
    - [Comentários gerais](#coment%C3%A1rios-gerais)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# golang-cod3r.com.br

Este repositório tem como principal objetivo fazer memória dos recursos de golang. Data de início do aprendizado: 01/02/2021.

## Fundamentos

### 1. Comentários gerais
  - comentários podem ser feitos por /* */ ou //

### 2. Comandos úteis
  - $ got vet arquivo.go //faz uma verificação no arquivo.go em busca de inconsistências
  - $ go build arquivo.go //compila o arquivo.go
  - $ ./arquivo //executa arquivo
  - $ go run arquivo.go //compila e executa 
 
### 3. Tipos
  - reflect.TypeOf(32000) //método para retornar o tipo de uma variável/valor
  - var i2 rune = "a" // comando rune mostra o valor da tabela ASCII. Se printarmos o valor de i2 o resultado é 97 (valor ascii para o caracterere 'a').
  - valueMaxInt64 = max.MaxInt64 // método MaxInt64 da biblioteca math retorna o maior valor inteiro possível com 64 bits.
  
### 4. Valores zeros (padrão)
  - variáveis float e int64 quando criadas mas não inicializadas com valor possui o valor padrão de 0.
  - variável bool possui o valor padrão false.
  - ponteiro possui o valor zero como **nil**. O valor nil é o mesmo que null de outras linguagens, referindo-se a vazio.
  
 ### 4. Conversões
  - strconv.Itoa(123) //converte  int para string
  - strconv.Atoi("123") //converte string para int
  - num, _ := strconv.Atoi("123") //variável num caso a conversão tenha sucesso e _ caso dê erro. Como declarei o símbolo _ no lugar de uma variável para tratar o erro, consigo executar o código mesmo não tratando o erro posteriormente. Caso utilize uma variável err no lugar do símbolo *_*, necessariamente precisarei tratar o erro.
