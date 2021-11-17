# Slices
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

