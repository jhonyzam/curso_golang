# Arrays
Arrays em go são coleções indexadas de valores do mesmo tipo e de tamanho fixo e invariável. 
O primeiro elemento do array possui indice `0`, e o último elemento é sempre `len(array) - 1`.
Para declarar um array, podemos utilizar um das seguintes formas:

~~~go
var a [3]int

numeros := [5]int{1, 2, 3, 4, 5}

primos := [...]int{2, 3, 5, 7, 11, 13}

nomes := [2]string{}
~~~
O array `a` foi declarado como `[3]int`, ou uma lista de três números inteiros. O tamanho
do array deve sempre ser especificado na declaração e faz parte do tipo do arrya, logo
`[3]int` e `[5]int` são considerados tipos diferentes.

O valor de `a` ao ser impresso foi `[0 0 0 ]`. Quando um array é declarado, seus elementos ganham
automaticamento um valor inicial conhecido como _zero value_ em Go. Este valor varia de acordo
com o tipo de dados declarado:

* `false` para valores do tipo `bool`
* 0 para `ints`
* 0.0 para `floats`
* "" (ou `string` vazia) para `strings`
* `nil` para ponteiros, funções, interfaces, slices, maps e channels

Muitas vezes vamos saber qual os valores contidos no array, para facilicar a vaida do programador
pode-se substituir o tamanho do array  declarado por _ellipsis_ (reticências), instruindo o
compilador a calcular o tamanho do array com base na quantidade de elementos declarados como
fizemos no array `primos`.

O tamanho de um array pode ser obtido através do uso da função `len()`. Por exemplo, a execução
do código a seguir nos daria `3 5 6 2` como resultado:
~~~go
fmt.Println(len(a), len(numeros), len(primos), len(nomes))
~~~
Existe também a possibilidade de criar arrays cujos valores são também arrays, ou arrays
multidimencionais.
~~~go
var multiA [2][2]int
multiA[0][0], multiA[0][1] = 3, 5
multiA[1][0], multiA[1][1] = 7, 2
fmt.Println("Array multiA: ", multiA)

multiB := [2][2]int{{2, 13}, {-1, 6}}
fmt.Println("Array multiB: ", multiB)
~~~
Este código tem como resultado:
````
Multi A: [[3 5] [7 -2]]
Multi B: [[2 13] [-1 6]]
````
Os arrays têm sua importância e seu papel, mas não são muito flexíveis, é possível aumentar o
seu tamanhao mas exige muito trabalho manual e verificação de limites, cópia e geração de novos
arrays.
Em Go é mais comum o uso de slices, a própria biblioteca padrão utiliza slices em vez de arrays
como parâmetros e tipos de retornos.