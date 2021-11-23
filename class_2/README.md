# 2- Criando funções

### 2.1- Funções básicas
O primeiro exemplo vamos entrar no mundo das `func`, as funções são fundamentais no Go. Podemos 
definir funções como pequenas unidades de código que podem abstrair ações, retornar e/ou receber valores,
auxiliando e organizando o código.

A função `imprimirDados` neste exemplo recebe dois argumentos, uma `string` representando o nome, e um `int`
representando a idade, e de forma simples imprime os dois usando a função `fmt.Printf()`: 
~~~go
func imprimirDados(nome string, idade int) {
    fmt.Printf("%s tem %d anos.", nome, idade)
}

func main() {
    imprimirDados("Jhonatan", 31)
}
~~~
Esta outra função como o proprio nome `soma()` ja diz, é usada para somar dois argumentos. Ela recebe os argumentos `n` e `m` 
do tipo `int` e retorna a soma dos dois. Como ambos são do mesmo tipo é possivel apenas separar os dois por virgula e
especificar o tipo apenas uma única vez:
~~~go
func soma(n, m int) int {
    return n + m
}

func main() {
    s := soma(3, 5)
    fmt.Println("A soma é: ", s)
}
~~~
### 2.2- Quicksort
Neste segundo exemplo, vamos implementar um algoritmo de ordenação.

1. Eleger um elemento da lista como pivô e removê-lo
2. Dividir a lista em duas listas distintas, uma com elementos menores que o pivô, e outra com os maiores
3. Ordenar as duas listas recursivamente
4. Retornar a combinação da lista ordenada de elementos menores, o pivô, e a lista ordenada com os elementos
maiores.
   
Este programa se limita a ordenar inteiros, e devem ser passados via argumento na execução. É o mesmo
princípio do conversor de moeda, porém vamos usar a função `strconv.Atoi()` para converter as `strings` em
números inteiros.

Este código é muito semelhante ao que já vimos, exceto pela função nativa `make()`. Essa função cria e
inicializa um slice de inteiros. Diferente de um array que tem um tamanho fixo, um slice, por outro lado,
é uma visualização flexível e de tamanho dinâmico dos elementos de uma matriz. Na prática, os slices
são muito mais comuns do que arrays em go.

Depois de converter a entrada do programa para uma lista de inteiros, a função `quicksort()` é chamada,
passando a lista como argumento, e a lista resultante é impressa como resultado:
~~~go
func main() {
	entrada := os.Args[1:]
	numeros := make([]int, len(entrada))

	for i, n := range entrada {

		numero, err := strconv.Atoi(n)
		if err != nil {
			fmt.Printf("%s não é um número válido", n)
			os.Exit(1)
		}
		numeros[i] = numero
	}

	fmt.Println(quicksort(numeros))
}
~~~
A função `quicksort()` é responsavel por implementar o algoritmo de ordenação:
~~~go
func quicksort(numeros []int) []int {
	// ...
}
~~~
O primeiro passo é verificar se a lista possui mais de um número, caso não, retornar a própria lista. Isso 
é muito importante, essa validação é chamado de "condição de parada", previne que a função seja
executada eternamente:
~~~go
if len(numeros) <= 1 {
    return numeros
}
~~~
O próximo passo é criar uma cópia do slice original para que ele não seja alterado. Para isso é usada a função
`copy()` que basicamente faz uma cópia dos valores do segundo argumento para o primeiro argumento. Após o a 
cópia apenas o novo slice `n` é manipulado:
~~~go
n := make([]int, len(numeros))
copy(n, numeros)
~~~
Em seguida é feita a escolha do `pivo`, basicamente é feito uma divisão pela metade para pegar o elemento mais ou
menos no meio da lista, e armazenamos o índice para utilizar na sequência:
~~~go
indicePivo := len(n) / 2
pivo := n[indicePivo]
~~~
Com o índice do pivô em mãos precisamos removê-lo da lista, e para isso vamos usar a função `append()`. Ela adiciona
um elemento ao final de um slice, ou seja, vamos criar um novo slice que não contenha o índice do `pivo`.
Primeiro fatiamos o slice `n` do primeiro elemento até o `pivo` `n[:indicePivo]`, e segundo partindo imediatamente 
posterior ao `pivo` até o último elemento disponível `[indicePivo+1:]`, este slice será adicionado ao slice-base.
É importante notar o uso de reticências ao final do segundo argumento, este informa que todos os elementos do segundo slice 
devem ser adicionados ao slice-base:
~~~go
n = append(n[:indicePivo], n[indicePivo+1:]...)
~~~
O resultado dessa operação é uma lista com todos os elementos anteriores e todos os elementos posteriores
ao `pivo`. Sendo assim uma remoção do `pivo`.

O próximo passo é particionar o slice em dois novos slices, um com os elementos menores ou iguais ao `pivo`,
e outro contendo os elementos maiores. Para isso criamos a função `particionar()`:
~~~go
menores, maiores := particionar(n, pivo)
~~~
A função `particionar()` recebe dois argumentos, um slice de inteiros (lista de menores ou maiores), e um inteiro (pivô),
e retorna dois valores, um slice de `menores[] int` contendo todos os números menores ou iguais ao `pivo`, e 
um slice de `maiores []int` contendo todos os números maiores que o `pivo`. A lógica segue o mesmo conceito anterior 
usando a função `append()`, validando o valor do `pivo` e distribuindo para os slices correspondentes:
~~~go
func particionar(numeros []int, pivo int) (menores []int, maiores []int) {
    for _, n := range numeros {
        if n <= pivo {
            menores = append(menores, n)
        } else {
            maiores = append(maiores, n)
        }
    }
    return menores, maiores
}
~~~
O ponto de atenção na função `particionar()` é o símbolo `_` (underline), conhecido como "identificador vazio".
Como já vimos o operador `range`, quando usado para iterar sobre um slice, retorna sempre dois valores, o índice, 
e o valor do elemento. Neste caso não precisamos do índice, e em GO uma váriavel declarada e não utilizada causa 
um erro de compilação. Para isso usamos esse identificador vazio.

Por fim temos o retorno da função `quicksort()`. Primeiro, chamamos `quicksort()` recursivamente para ordenar o slice 
`menores`, depois adicionamos o `pivo` ao resultado desta ordenação, e em seguida é feita outra chamada
recursiva na função `quicksort()` para ordenar o slice de `maiores`. Por fim ambas as listas são combinadas
com a função `append()`:
~~~go
return append(append(quicksort(menores), pivo), quicksort(maiores)...)
~~~

#### Execução
````
go run class_2/quicksort.go 10 30 41 53 78 12 19 22

Resultado:
[10 12 19 22 30 41 53 78]
````



