# Coleções

Ja vimos alguns exemplos de coleções, porém agora vamos detalhar um pouco os arrays, 
slices e maps.

### 1- Arrays
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

### 2- Slices
Os slices são uma poderasa abstraçao criada em cima de arrays que possui uma série de facilidades
amais. Diferente de um array, o slice possui tamanha variável e pode crescer indefinidamente.
Para declarar um slice, utilizamos quase a mesma sintaxe da declaração de arrays, a diferença é 
que não especificamos o tamanho.
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
retorna uma refência para o slice criado.
~~~go
func make([]T, len, cap) []T
~~~
`T` representa o tipo dos elementos do slice, `len`, o tamanhao inicial do array alocado e
`cap`  tamanho total da área de memória reservada para o crescimento do slice. O último argumento 
pode ser omitido, e neste caso Go assume o mesmo valor de tamanho.
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

### 1- Iteradores
Ja iteramos sobre slices nos exemplos anteriores, agora vamos ver as diferentes formas de
iteração disponíveis em Go
Como já vimos a única estrutura de frepetição em Go é o `for`. Em sua forma mais básica,
criamos uma condição lógoca onde loop será executado enquando a condição for verdadeira, 
semelhante à construção de um `while` em outras linguagens.
~~~go
a, b := 0, 10

for a < b {
    a += 1
}

fmt.Println(a)
~~~
O código acima representa, enquanto `a` for menor que `b`, incremente o valor de `a`. O 
resultado é a impressão do valor `10`, que é o vaor de `a` após a execução do `for`.
Também pode ser usado o `for` como cláusula de inicialização, um condição lógica e uma cláusula
de incremento.
~~~go
for i := 0; i < 10; i++ {
    // ...
}
~~~
É importante notar que a variável `i` não existeia antes do `for`, sendo assim o escopo dela
se limita ao bloco, não podendo ser acessada após a exeução fo `for`. Para isso ela precisa
ser declarada antes.
~~~go
var i int

for i = 0; i < 10; i++ {
    // ...
}

fmt.Println(i)
~~~
Desta forma, `i` continua existindo após a execução do bloco `for`.

A forma mais comum de iterar sobre slices, seria utilizando o operador `range`, como já
vimos antes.
~~~go
for indice, valor := range slice { 
    // ...
}
~~~
O operador `range` retorna o índice de cada elemento, começando em `0`, e uma cópia de cada
valor presente no slice. Quando precisamos modificar valores de um slice, ou queremos apenas
os índices, podemos apenas omitir o segundo valor, e acessar cada elemente através o índice.
~~~go
numeros := []int{1, 2, 3, 4, 5}

for i := range numero { 
    numeros[i] *= 2
}

fmt.Println(numeros)
~~~
Este código itera sobre um slice chamado `numeros`, multiplicando cada valor por 2. No final 
é impresso `[2 3 6 8 10]` como resultado.

Ao contrário, quando não precisamos dos índices dos valores podemos ignorá-los atribuindo-os
ao identificador vazio
~~~go
for _, elemento := range slice { 
    // ...
}
~~~
A última forma de utilizar o `for` é o que chamamos de _loop infinito_, que em Go é simplismente
uma cláusula `for` sem nenhuma condição. O Go assume `true` como padrão.
~~~go
for { 
    // ...
}
~~~
Para finalizar a execução, basta usar o comando `break`.

O próximo exemplo inicia um loop infinito, gera números aleatórios e sai do loop somente
quando o número gerado por divisível por 42, ou seja, quando o resto da divisão por 42 for 
0. O exemplo está pronto no arquivo `loop_infinito.go`.
~~~go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {
    rand.Seed(time.Now().UnixNano())
    n := 0
    
    for {
        n++
    
        i := rand.Intn(4200)
        fmt.Println(i)
    
        if i%42 == 0 {
            break
        }
    }
    
    fmt.Printf("Saída após %d iterações.\n", n)
}
~~~
Sempre que geramos números aleatórios, é importante configurar o vaor conhecido como _seed_(semente)
do gerador. No nosso exemplo, foi utilizado o _timestamp_ atual no formato padrão do Unix, 
para garantir que, a cada execução o gerador produza números diferentes do anterior.
Resultado:
````
1914
1063
3241
3778
2459
2604
Saída após 6 iterações.
````
Lembrando que a cada execução teremos um resultado diferente.

Por padrão como mencionado, `break` sai do loop atual. Em casos de loops aninhados, onde 
desejamos quebrar o loop externo em vez do interno, podemos nomear os blocos `for`.
~~~go
var i int

externo:
for {
    for i = 0; i < 10; i++ {
        if i == 5 {
            break externo
        }
    }
}

fmt.Println(i)
~~~
Este recurso também é muito importante quando temos, por exemplo, um bloco `switch` dentro de
um bloco `for`. O comando `break` também é usado para sair do `switch`, com isso precisamos
especificar o nome dado ao bloco `for` caso precisamos sair do loop. Vamos criar o exemplo
`loop_nomeado.go` para exemplificar.
~~~go
var i int

loop:
    for i = 0; i < 10; i++ {
        fmt.Printf("for i = %d\n", i)
    
        switch i {
        case 2, 3:
            fmt.Printf("Quebrando switch, i == %d. \n", i)
            break
        case 5:
            fmt.Println("Quebrando loop, i == 5.")
            break loop
        }
    }
    fmt.Println("Fim.")
}
~~~
Resultado:
````
for i = 0
for i = 1
for i = 2
Quebrando switch, i == 2. 
for i = 3
Quebrando switch, i == 3. 
for i = 4
for i = 5
Quebrando loop, i == 5.
Fim.
````

