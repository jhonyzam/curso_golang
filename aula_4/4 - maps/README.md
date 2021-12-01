# 4.4 Maps
Um `map`, ou mapa, é uma coleção de pares _chave-valor_ sem nenhuma ordem definida. É a implementação em Go
de uma estrutura de dados também conhecida como _hashtable_, dicionário de dados ou array 
associativo, entre outros nomes. As chaves em um mesmo mapa devem necessariamente ser do mesmo tipo, e que estas são 
unicas, se armazenar dois valores distintos sob uma mesma chave, o primeiro valor será sobrescrito.
É possível usar o tipo `interface{}`, que possibilita armazenar qualquer valor sob uma chave, 
porém é necessário que seja validado o tipo com um _type assertion_, para evitar erros em sua
utilização.
É possível declarar mapas utilizando a form literal ou a função `make()`:
~~~go
vazio1 := map[int]string{}
vazio2 := make(map[int]string)
~~~
A quantidade de valores armazenados em um mapa é flexível e pode crescer indefinidamente durante 
a execução de um programa, sendo que podemos especificar sua capacidade inicial quando sabemos
de antemão quantos valores precisaremos armazenar. Isso pode ser importante para tornar o uso
de memória mais eficiente e evitando problemas de performace. Sempre que for possível é importante
fazer isso:
~~~go
mapaGrande := make (map[int]string, 4096)
~~~
A qualquer momento poemos inspecionar a quantidade de elementos que um mapa possui através
da função `len()`:
~~~go
capitais := map[string]string{
    "GO": "Goiânia",
    "PB": "João Pessoa",
    "PR": "Curitiba",
}

fmt.Println(len(capitais))
~~~
Resultado:
````
3
````
### 4.4.1- Populando maps
Podemos popular um mapa utilizando literais no momento da declaração ou atribuindo valores
individualmente após a declaração:
~~~go
capitais := map[string]string{
    "GO": "Goiânia",
    "PB": "João Pessoa",
    "PR": "Curitiba",
}

fmt.Println(capitais)
~~~
ou:
~~~go
capitais := map[string]string{}

capitais["GO"] = "Goiânia"
capitais["PB"] = "João Pessoa"
capitais["PR"] = "Curitiba"

fmt.Println(capitais)
~~~
Resultado:
````
map[GO:Goiânia PB:João Pessoa PR:Curitiba]
````
Vamos exemplificar com mais detalhes o que mostramos até agora, usando o exemplo em `maps.go`.
Primeiro criamos um novo tipo, chamado `Estado`:
~~~go
type Estado struct {
    nome      string    
    populacao int
    capital   string
}
~~~
Na função `main()`, populamos um mapa de estados:
~~~go
estados := make(map[string]Estado, 6)

estados["GO"] = Estado{"Goiás", 6434052, "Goiânia"}
estados["PB"] = Estado{"Paraíba", 3914418, "João Pessoa"}
estados["PR"] = Estado{"Paraná", 10997462, "Curitiba"}
estados["RN"] = Estado{"Rio Grande do Norte", 3373960, "Natal"}
estados["AM"] = Estado{"Amazonas", 3807923, "Manaus"}
estados["SE"] = Estado{"Sergipe", 2228489, "Aracaju"}
~~~
Resultado:
````
map[AM:{Amazonas 3807923 Manaus} GO:{Goiás 6434052 Goiânia} PB:{Paraíba 3914418 João Pessoa} PR:{Paraná 10997462 Curitiba} RN:{Rio Grande do Norte 3373960 Natal} SE:{Sergipe 2228489 Aracaju}]
````
### 4.4.2- Lookup: recuperando valores
A operação de recuperação de um valor em um mapa é conhecida como _lookup_ é muito similar ao índice
específico em um slice, basta especificara chave entre colchetes:
~~~go
sergipe : = estados["SE"]

fmt.Println(sergipe)
~~~
Desta forma a variável `sergipe` receberia o `Estado` presente no mapa e o valor `{Sergipe 2228489 Aracaju}`
seria impresso.
Mas oq acontece se tentarmos acessar um valor de um estado que não esta no mapa:
~~~go
fmt.Println(estados["SP"])
~~~
O valor `{ 0 }` é impresso. Note que existe um espaço em branco antes e outro depois do valor.
Muitas vezes isso pode causar comportamentos inesperados em um programa, e para evitar pode ser
feito um teste se a chave existe ou não:
~~~go
saoPaulo, encontrado = estados["SP"]
if encontrado {
    fmt.Println(saoPaulo)
}
~~~
O segundo valor retornado pela operação de lookup é um `bool` que receberá o valor `true`
caso a chave esteja presente no map, ou `false` caso contrário.

### 4.4.3- Atualizando valores
Para atualizar valores existentes em um mapa, utilizamos a mesma sintaxe da inserção de
um novo valor, isso por que as chaves são únicas então o valor é atualizado:
~~~go
idades := map[string]int{
    "João": 37,
    "Ricardo": 26,
    "Jhony": 30,
}

idades["Jhony"] = 31

fmt.Println(idades["Jhony"])
~~~
Resultado:
````
31
````
### 4.4.4- Removendo valores
Podemos remover valores presentes em um mapa utilizando a função `delete()`, por exemplo se
quisermos remover o estado do Amazonas do mapa de estados utilizado anteriormente:
~~~go
delete(estados, "AM")
~~~
### 4.4.5- Iterando sobre maps
Podemos utilizar o operador `range` para iterar sobre todas as entradas de um mapa, o 
exemplo abaixo usa o mapa de estados utilizado anteriormente:
~~~go
for sigla, estado := range estados {
    fmt.Printf("%s (%s) possui %d habitantes.\n", estado.nome, sigla, estado.populacao)
}
~~~
Resultado:
````
Sergipe (SE) possui 2228489 habitantes.
Goiás (GO) possui 6434052 habitantes.
Paraíba (PB) possui 3914418 habitantes.
Paraná (PR) possui 10997462 habitantes.
Rio Grande do Norte (RN) possui 3373960 habitantes.
Amazonas (AM) possui 3807923 habitantes.
````
Ja mencionamos antes que a ordem dos elementos no mapa não é garantida, no exemplo abaixo:
~~~go
quadrados := make (map[int]int, 15)

for i := 1; i <= 15; i++ {
    quadrados[i] = i * i
}

for n, quadrado := range quadrados {    
    fmt.Printf("%d^2 = %d\n", n, quadrado)
}
~~~
Resultado:
````
6^2 = 36
10^2 = 100
15^2 = 225
1^2 = 1
3^2 = 9
4^2 = 16
11^2 = 121
13^2 = 169
5^2 = 25
8^2 = 64
12^2 = 144
14^2 = 196
2^2 = 4
7^2 = 49
9^2 = 81
````
Caso execute novamente a ordenação pode ser alterada.

Muitas vezes precisamos apresentar os dados seguinte uma ordem definida. Em Go, a forma
recomendada é manter uma estrutura de dados separada contendo as chaves ordenadas, iterando
sobre esta estrutura e obtendo os valores correspondetes no mapa. Vamos usar o exemplo anterior 
para aplicar esta técnica usando as facilidades do pacote `sort`, usar o arquivo `map_ordenado.go`:
~~~go
package main

import (
    "fmt"
    "sort"
)

func main() {
    quadrados := make(map[int]int, 15)

    for i := 1; i <= 15; i++ {
        quadrados[i] = i * i
    }
    
    numeros := make([]int, 0, len(quadrados))
    
    for n := range quadrados {
        numeros = append(numeros, n)
    }
    sort.Ints(numeros)
    
    for _, numero := range numeros {
        quadrado := quadrados[numero]
        fmt.Printf("%d^2 = %d\n", numero, quadrado)
    }
}
~~~
Resultado:
````
1^2 = 1
2^2 = 4
3^2 = 9
4^2 = 16
5^2 = 25
6^2 = 36
7^2 = 49
8^2 = 64
9^2 = 81
10^2 = 100
11^2 = 121
12^2 = 144
13^2 = 169
14^2 = 196
15^2 = 225
````
Com o uso da função `sort.Ints()`  ordenamos o slice com as cahves do mapa garantindo assim
o mesmo resultado de saída.