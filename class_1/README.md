# If e Expressões logicas

### Conversor de moeda

O primeiro exemplo é um simples conversor de moeda. Ele aceita como entrada um valor
string inicial sendo real ou dolar e uma lista de valores, devolvendo uma lista de 
valores convertidos.

A primeira parte do codigo é onde adicionamos o package. O package em Go é simplesmente um 
diretório (pasta) com um ou mais arquivos .go dentro dele. Packages Go fornecem isolamento e 
organização de código semelhante a diretórios e organização de arquivos em um computador.
Todo o código Go vive em um pacote e um pacote é o ponto de entrada para acessar o código Go.

~~~go
package main
~~~
Em seguida os `imports` das libs a serem utilizadas no código.

* O pacote `fmt` é utilizado para leitura e escrita no console.
* O pacote `os` lida com uma série de operações do sistema operacional em que o codigo 
  esta sendo executado.
* O pacote `strconv` fornece uma variedade de funções para conversão de strings para 
  outros formatos.
~~~go
import (
	"fmt"
	"os"
	"strconv"
)
~~~
A linha abaixo representa uma constante. Constante são valores reservados na memoria que não
podem ser alterados. O valor 5.54 representa o valor do dólar no dia da criação desse conteúdo.
~~~go
const valorDolarPorReal = 5.54
~~~
Na sequência é feito a declaração da função main. Como o nome sugere é a principal função do go.
Ela não recebe nenhum parâmetro, nem retorna valores.
~~~go
func main() {
	// ...
}
~~~
O objeto `os.Args` possui uma lista de todos os argumentos passados para o programa. O primmeiro
o índice 0 sempre sera o proprio nome do programa. Sendo assim é feita uma validação se `os.Args`
possui pelo menos três argumentos para ser uma entrada válida. Se for passado menos que três
argumentos, utilizamos a função `Println` do pacote `fmt` para apresentar uma mensagem de ajuda
com o formato de entrada do programa, e na sequência a função `os.Exit` para que o programa finalize.
~~~go
if len(os.Args) < 3 {
    fmt.Println("Uso: conversor <moeda> <valores>")
    os.Exit(1)
}
~~~
Na sequência atribuimos o argumento do índice 1, que deve ser a moeda origem, e dos argumentos do
índice 2 até o fim da lista, os valores a serem convertidos.
~~~go
moedaOrigem := os.Args[1]
valoresOrigem := os.Args[2:len(os.Args)]
~~~
A linha abaixo representa uma variável. Variáveis são valores reservados na memória que precisam
ser alterados durante o tempo de execução do programa.
~~~go
var moedaDestino string
~~~
É necessário verificar qual o valor de origem, se for real o destino será dólar e vice e versa.
~~~go
if moedaOrigem == "real" {
    moedaDestino = "dolar"
} else if moedaOrigem == "dolar" {
    moedaDestino = "real"
} else {
    fmt.Printf("%s não é uma moeda conhecida!", moedaDestino)
    os.Exit(1)
}
~~~
Semelhante a outras linguagens, Go também possui uma estrutura de repetição `for`, e nele ela é unica,
podendo ser usada de diferentes formas. No exemplo abaixo ela é utilizada com o operador `range` para 
obter acesso a cada elemento.  O operador range quando aplicado a um slice, retorna 2 valores para cada
elemento: primeiro o indice do elemento `i` e depois o valor `v`.
~~~go
for i, v := range valoresOrigem {
// ...
}
~~~
Em seguida utilizamos a função `ParseFloat()` do pacote `strconv` para converter uma `string` em um número.
Esta função recebe dois argumentos, o valor a ser convertido e a precisão do valor retornado (32 ou 64 bits). Retornando
dois valores, o valor converto e um erro (que será `nil` em caso de convertido com sucesso). Caso haja um erro é feito 
a validação se a variavel `err` é diferente de `nil`, caso seja é mostado uma mensagem no console para o usuário.
~~~go
valorOrigem, err := strconv.ParseFloat(v, 64)
if err != nil {
    fmt.Printf(
        "O valor %s na posição %d não é um número válido!\n", v, i)
    os.Exit(1)
}
~~~    
Com o valor convertido, é preciso apenas validar qual a moeda de origem
para transformar para o valor da moeda de destino 
~~~go
var valorDestino float64

if moedaOrigem == "real" {
    valorDestino = valorOrigem * valorDolarPorReal
} else {
    valorDestino = valorOrigem / valorDolarPorReal
}
~~~
Por fim é apresentado o valor convertido e sua unidade para o usuario. Ao utilizar `%2.f` a funcão 
`Printf` reconhece como valor float e arredonda o valor para 2 casas decimais.
~~~go
fmt.Printf("%.2f %s = %.2f %s\n", valorOrigem, moedaOrigem, valorDestino, moedaDestino)
~~~
Por fim para executar o programa basta usar o exemplo:
````
go run class_1/conversor_moeda.go dolar 10 24

//Retorno
10.00 dolar = 55.40 real
24.00 dolar = 132.96 real
````
