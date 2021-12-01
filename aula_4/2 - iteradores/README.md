# 4.2 Iteradores
Como já vimos a única estrutura de repetição em Go é o `for`. Em sua forma mais básica,
criamos uma condição lógoca onde loop será executado enquando a condição for verdadeira, 
semelhante à construção de um `while` em outras linguagens:
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
de incremento:
~~~go
for i := 0; i < 10; i++ {
    // ...
}
~~~
É importante notar que a variável `i` não existeia antes do `for`, sendo assim o escopo dela
se limita ao bloco, não podendo ser acessada após a exeução fo `for`. Para isso ela precisa
ser declarada antes:
~~~go
var i int

for i = 0; i < 10; i++ {
    // ...
}

fmt.Println(i)
~~~
Desta forma, `i` continua existindo após a execução do bloco `for`.

A forma mais comum de iterar sobre slices, seria utilizando o operador `range`, como já
vimos antes:
~~~go
for indice, valor := range slice { 
    // ...
}
~~~
O operador `range` retorna o índice de cada elemento, começando em `0`, e uma cópia de cada
valor presente no slice. Quando precisamos modificar valores de um slice, ou queremos apenas
os índices, podemos apenas omitir o segundo valor, e acessar cada elemente através o índice:
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
ao identificador vazio:
~~~go
for _, elemento := range slice { 
    // ...
}
~~~
A última forma de utilizar o `for` é o que chamamos de _loop infinito_, que em Go é simplismente
uma cláusula `for` sem nenhuma condição. O Go assume `true` como padrão:
~~~go
for { 
    // ...
}
~~~
Para finalizar a execução, basta usar o comando `break`.

O próximo exemplo inicia um loop infinito, gera números aleatórios e sai do loop somente
quando o número gerado por divisível por 42, ou seja, quando o resto da divisão por 42 for 
0. O exemplo está pronto no arquivo `loop_infinito.go`:
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
desejamos quebrar o loop externo em vez do interno, podemos nomear os blocos `for`:
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
`loop_nomeado.go` para exemplificar:
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

