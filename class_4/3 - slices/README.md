# 4.3 Slices
Os slices são uma poderasa abstração criada em cima de arrays que possui uma série de facilidades
amais. Diferente de um array, o slice possui tamanha variável e pode crescer indefinidamente.
Para declarar um slice, utilizamos quase a mesma sintaxe da declaração de arrays, a diferença é 
que não especificamos o tamanho:
~~~go
var a []int
primos := []int{2, 3, 5, 7, 11, 13}
nomes := []string{}

fmt.Println(a, primos, nomes)
~~~
Resultado:
````
[] [2 3 5 7 11 13] []
````
Como `a` e `nomes` não possuem valor, ambos foram impressos com vazio.
Arrays e slices em Go possuem duas propriedades importantes tamanho `len()` e capacidade `cap()`.
Um slice pode ser criado também atravès da função `make()`, que internamente aloca um array e 
retorna uma refência para o slice criado:
~~~go
func make([]T, len, cap) []T
~~~
`T` representa o tipo dos elementos do slice, `len`, o tamanhao inicial do array alocado e
`cap`  tamanho total da área de memória reservada para o crescimento do slice. O último argumento 
pode ser omitido, e neste caso Go assume o mesmo valor de tamanho:
~~~go
b := make ([]int, 10)
fmt.Println(b, len(b), cap(b))

c := make ([]int, 10, 20)
fmt.Println(c, len(c), cap(c))
~~~
Resultado:
````
[0 0 0 0 0 0 0 0 0 0] 10 10
[0 0 0 0 0 0 0 0 0 0] 10 20
````
A vantagem de usar slices através da função `make()` é que quando usados como argumentos
ou no retorno de funções, estes são passados por referência e não por cópia. Isto faz om que
as chamadas sejam muito mais eficientes, pois o tamanho da referência será sempre o mesmo,
independente do tamanho do slice.

### 4.3.1 Fatiando slices
Para fatiar ou extrais uma parte de um slice ou de um array, utilizamos a seguinte forma:
~~~go
novoSlice := slice[inicio : fim]
~~~
Sendo que ` <= inicio <= fim <= len(slice)`. Qualquer combinação de índices incial e final que não
atenda a esta regra resulta em um erro de compilação (_slice bounds out of range_), significa que 
foi usado um valor de índice fora dos limites do slice.
Por questões práticas, podemos omitir o índice inicial e/ou final. Se o índice inicial for
omitido, `0` é assumido como padao, bem como `len(slice)-1` é assumido como valor padrão para o 
índice final:
~~~go
fib := []int{1, 1, 2, 3, 5, 8, 13}
fmt.Println(fib)
fmt.Println(fib[:3])
fmt.Println(fib[2:])
fmt.Println(fib[:])
~~~
Resultado:
````
[1 1 2 3 5 8 13]
[1 1 2]
[2 3 5 8 13]
[1 1 2 3 5 8 13] 
````
Sabemos que, quando um slice é criado, um array é alocado internamente. Ao fatiarmos este
slice e o atribuírmos a um nova variável, temos um novo slice que compartilha o mesmo
array interno do original. Isto quer dizer que, quando um elemento comum aos dois slices
é modificado, esta modificação é refletinado no outro:
~~~go
original := []int{1, 2, 3, 4, 5}
fmt.Println("Original:", original)

novo := original[1:3]
fmt.Println("Novo:", novo)

original[2] = 13

fmt.Println("Original pós modificação:", original)
fmt.Println("Novo pós modificação:", novo)
~~~
Resultado:
````
Original: [1 2 3 4 5]
Novo: [2 3]
Original pós modificação: [1 2 13 4 5]
Novo pós modificação: [2 13]
````
Veja que o valor de `original[2]` foi alterado para 13, o conteúdo dos dois slices foi 
modificado, isto acontece para qualquer slice criado fatiando um outro slice. Isto também é 
válido para um array, mas sempre que fatiar uma array o resultado é um slice, nunca outro
array.

### 4.3.2 Inserindo valores
Todas as operações realizadas sobre slices são baseadas na função `append()`. Já vimos alguns
exemplos de como usá-la.
Para inserir um novo valor ao final de um slice, utilizamos `append()` em sua forma mais
básica:
~~~go
s := make([]int, 0)
s = append(s, 23)

fmt.Println(s)
~~~
Resultado:
````
[23]
````
E para inserir um novo valor no começo de um slice, precisamos inicialmente criar um novo
slice contendo o valor que desejamos, e depois adicionar todos os elementos do slice inicial
ao recém criado:
básica:
~~~go
s := []int{23, 24, 25}
n := []int{22} 
s = append(n, s...)

fmt.Println(s)
~~~
Resultado:
````
[22 23 24 25]
````
Também é possível inserir um ou mais valores em qualquer posição do slice. Para isso, 
precisamos fatiar o slice até a posição onde desajamos inserir os novos valores:
~~~go
s := []int{11, 12, 13, 16, 17, 18}
v := []int{14, 15} 
s = append(s[:3], append(v, s[3:]...)...)

fmt.Println(s)
~~~
Resultado:
````
[11 12 13 14 15 16 17 18]
````

### 4.3.3 Removendo valores
Para remover valores do começo de um slice não precisamos da função `append()`. Basta fatiar
o slice ignorando os índices dos elementos que desejamos remover, e atribuir o novo slice
à mesma variável:
~~~go
s := []int{20, 30, 40, 50, 60}
s = s[1:]

fmt.Println(s)
~~~
Resultado:
````
[30 40 50 60]
````
Da mesma forma, para remover valores do final de um slice, fatiamos este ignorando os 
índices dos elementos finais.
~~~go
s := []int{20, 30, 40, 50, 60}
s = s[:3]

fmt.Println(s)
~~~
Resultado:
````
[20 30 40]
````
Por fim, para remover valores do meio de um slice, precisamos utilziar a função `append()`,
utilizando duas fatias do slice original como argumentos.
~~~go
s := []int{10, 20, 30, 40, 50, 60}
s = append(s[:2], s[4:]...)

fmt.Println(s)
~~~
Resultado:
````
[10 20 50 60]
````

### 4.3.4 Copiando slices
Até agora sempre que manipulamos slices, utilizamos a função `append()` para modificar
o slice original. Porém muitas vezes precisamos manter o estado do slice original intacto
e manipular uma cópia dele. Para isso usamos a função `copy()` com a seguinte assinatura:
~~~go
func copy(destino, origem []Tipo) int
~~~
Para criar uma cópia de um slice, chamamos a função `copy()` da seguinte forma:
~~~go
numeros := []int{1, 2, 3, 4, 5}
dobros := make([]int, len(numeros))

copy(dobros, numeros)

for i := range dobros {
    dobros[i] *= 2
}

fmt.Println(numeros)
fmt.Println(dobros)
~~~
Resultado:
````
[1 2 3 4 5]
[2 4 6 8 10]
````