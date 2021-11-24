# 5- Criando novos tipos

Criar tipos customizados é uma ferramenta de abstração muito poderosa em linguagens de programação. Go, comparada a
linguagens puramente orientadas a objetos, tem um suporte limitado às características deste paradigma. Por exemplo não é
possível criar uma Hierarquia de tipos baseada em herança. Composição de tipos, no entanto, é suportada e encorajada em
Go.

### 5.1- Novos nomes para tipos existentes

Algumas linguagens dinâmicas como Ruby possibilitam que o programador altere o comportamento de qualquer classe ou
objeto durante a execução de um programa. Em Go isto não é possível, mas podemos estender tipos existentes através da
criação de novos tipos, por conveniência ou para melorar a legibilidade de um programa. Para demonstrar esta técnica,
vamos criar um tipo chama `ListaDeCompras` baseado em um slice de `strings`:

~~~go
type ListaDeCompras []string

func main() {
    lista := make (ListaDeCompras, 6)
    lista[0] = "Alface"
    lista[1] = "Pepino"
    lista[2] = "Azeite"
    lista[3] = "Atum"
    lista[4] = "Frango"
    lista[5] = "Chocolate"
    
    fmt.Println(lista)
}
~~~
O código anterior imprime `[Alface Pepino Azeite Atum Frango Chocolate]`. Inicialmente o tipo `ListaDeCompras`
não parece muito útil. Porém a grande vantagem em criar tipos customizados é que podemos estendê-los, algo impossível de
fazer com os tipos padrões da linguagem. Vamos estender o tipo `ListaDeCompras` e definir um método que separa os
elementos em categorias:
~~~go
func (lista ListaDeCompras) Categorizar() ([]string, []string, []string) {
    var vegetais, carnes, outros []string

    for _, e := range lista {
        switch e {
        case "Alface", "Pepino":
            vegetais = append(vegetais, e)
        case "Atum", "Frango":
            carnes = append(carnes, e)
        default:
            outros = append(outros, e)
        }
    }

    return vegetais, carnes, outros
}
~~~
Na sequencia imprimimos os valores categorizados:
~~~go
vegetais, carnes, outros := lista.Categorizar()
fmt.Println("Vegetais:", vegetais)
fmt.Println("Carnes:", carnes)
fmt.Println("Outros:", outros)
~~~
O exemplo completo está em `lista_compras.go`. Resultado:
````
Vegetais: [Alface Pepino]
Carnes: [Atum Frango]
Outros: [Azeite Chocolate]
````
Repare na forma com que o comando `switch` foi utilizado na implementação do método
`Categorizar()`, verificando múltiplos valores separados por vírgula. Também utilizamos a cláusula `default` para criar
o slice `outros`, contendo os valores não reconhecido pelas cláusulas anteriores.

### 5.2- Conversão de tipos compatíveis

Apesar do tipo `ListaDeCompras` ter sido criado com base no tipo `[]string`, na prática eles são diferentes e não são
automaticamente intercambiáveis. Desta forma não é possivel utilizar um valor `ListaDeCompras` em que um `[]string` é
esperado e vice-versa. Para exemplificar este problema, vamos realizar a conversão de tipo manualmente, no arquivo
`conversao.go`:
~~~go
package main

import "fmt"

type ListaDeCompras []string

func imprimirSlice(slice []string) {
    fmt.Println("Slice:", slice)
}

func imprimirLista(lista ListaDeCompras) {
    fmt.Println("Lista de compras:", lista)
}

func main() {
    lista := ListaDeCompras{"Alface", "Atum", "Azeite"}
    slice := []string{"Alface", "Atum", "Azeite"}
    
    imprimirSlice([]string(lista))
    imprimirLista(ListaDeCompras(slice))
}
~~~
Resultado:
````
Slice: [Alface Atum Azeite]
Lista de compras: [Alface Atum Azeite]
````

### 5.3- Criando abstrações

Anteriormente vimos como manipular slices com a função `append()`. Podemos também criar uma camada de abstração sobre os
slices e implmentar uma lista genérica que armazena valores de qualquer tipo, e criar operações para remover valores.
Primeiro dfinimos o tipo:

~~~go
type ListaGenerica []interface{}
~~~

Criamos um tipo chamado `ListaGenerica` baseado em um slice que armazena valores do tipo
`interface{}`, permitindo armazenar qualquer tipo. Agora vamos criar o método que remove os índices:

~~~go
func (lista *ListaGenerica) RemoverIndice(indice int) interface{} {
l := *lista
removido := l[indice]
*lista = append(l[0:indice], l[indice+1:]...)
return removido
}
~~~

A função `RemoverIndice` recebe a posição do índice, usa do metodo `append()` para montar a nova lsita e atribuir ao
ponteiro `*lista`. Isso tudo usando de funções ja aprendidas. Para implementar o método `RemoverInicio()`, basta apenas
chamarmos o `RemoveIndice()` passando
`0` como argumento:

~~~go
func (lista *ListaGenerica) RemoverInicio() interface{} {
return lista.RemoverIndice(0)
}
~~~

Isto funciona porque, subtituindo o `indice` na operação que altera a lista por `0`, obtemos a
expressão `append(l[0:0], l[1:]...)`, ou seja adicionamos os elementos da lista iniciando no índice 1 à lita vazia
retornada pela operação `l[0:0]`. Da mesma forma implementamos o método de remover o vaor final da lista, passando o
último índice da lista como argumento:
~~~go
func (lista *ListaGenerica) RemoverFim() interface{} {
    return lista.RemoverIndice(len(*lista) - 1)
}
~~~
O que ocorre aqui pe o inverso do que vimos anteriormente, assumindo que o último índice presente na lista seja 5,
obtemos a expressão `append(l[0:5], l[6:])`, ou seja, adicionamos a lista vazia retornada por `l[6:]` onde o índice 5
foi removido. O programa completo esta no exemplo `lista_generica.go`.

Resultado:

````
Lista original: 
[1 Café 42 true 23 Bola 3.14 false]

Removendo do início: 1, após remoção:
[Café 42 true 23 Bola 3.14 false]
Removendo do fim: false, após remoção:
[Café 42 true 23 Bola 3.14]
Removendo do índice 3: 23, após remoção:
[Café 42 true Bola 3.14]
Removendo do índice 0: Café, após remoção:
[42 true Bola 3.14]
Removendo do último índice 0: 3.14, após remoção:
[42 true Bola]
````

### 5.2- Structs

Ja explicamos anteriormente uma struct, e agora vamos detalhar um pocuo mais. Um struct é uma coleção de váriaveis que
forma um novo tipo. Elas são importantes para agrupar dados relacionados criando a noção de registros. Se estivéssemos
desenvolvendo um programa para extrair estatísticas de um arquivo de texto, poderíamos definir a struct `Arquivo` desta
forma:
~~~go
type Arquivo struct {
nome        string
tamanho     float64
caracteres  int
palavras    int
linhas      int
}
~~~
Com a struct definida, podemos criar uma instância especificando o valor dos campos na ordem em que eles foram definidos
e separados por vírgulas:
~~~go
arquivo := Arquivo{"artigo.txt", 12.68, 12986, 1862, 220}

fmt.Println(arquivo)
~~~
Desta forma teríamos o valor `{artigo.txt 12.68 12986 1862 220}` impresso. Por questões de clareza e organização pode
ser especificado os nomes dos campos ao inicializar uma struct, isso pode ser útil principalmente quando nem todos os
campos precisam ter valores atribuidos:
~~~go
codigoFonte := Arquivo{tamanho: 1.12, nome: "programa.go"}

fmt.Println(codigoFonte)
~~~
Este código imprime `{programa.go 1.12 0 0 0}` como resultado. Os valores não especificados
(`caracteres`, `palavras` e `linhas`) foram inicializados om 0, valor inicial para o tipo `int`.

Para acessar os valores em uma struct utilizamos o operador `.` parado do nome da variável e o nome do campo acessado:
~~~go
fmt.Printf("%s\t%.2fKB\n", arquivo.nome, arquivo.tamanho)
fmt.Printf("%s\t%.2fKB\n", codigoFonte.nome, codigoFonte.tamanho)
~~~

Resultado:

````
arquivo.txt	    12.68KB
programa.go	    1.12KB
````
Em programas maiores é comum inicializar uma struct e armazenar um ponteiro para ela numa variável que será manipulada
porteriormente. Para isso utilizamos o operador `&`, que retorna um ponteiro para a struct criada. Por conveniência, ao
manipular uma struct podemos omitir o operador `*`, como podemos ver a seguir:
~~~go
ponteiroArquivo := &Arquivo{tamanho: 7.29, nome: "arquivo.txt"}

fmt.Printf("%s\t%.2fKB\n", ponteiroArquivo.nome, ponteiroArquivo.tamanho)
~~~
Em lingaugens orientadas a objetos, definimos tipos usando classes, que reunem dados e métodos na classe. Em Go,
utilizamos structs para definir novos tipos. Assim como fizemos anteriormente quando definimos tipos que estendem os
tipos padrão da linguagem, podemos também definir métodos cujo receptores são structs, ficando muito semelhante à uma
classe. Vamos enriquecer nosso tipo `Arquivo` criando dois métodos, `TamanhoMedioDePalavra()` e
`MediaDePalavraPorLinha()`:

~~~go
func (arq *Arquivo) TamanhoMedioDePalavra() float64 {
    return float64(arq.caracteres) / float64(arq.palavras)
}

func (arq *Arquivo) MediaDePalavraPorLinha() float64 {
    return float64(arq.palavras) / float64(arq.linhas)
}

func main() {
    arquivo := Arquivo{"artigo.txt", 12.68, 12986, 1862, 220}
    
    fmt.Printf("Média de palavras por linha: %.2f\n", arquivo.MediaDePalavraPorLinha())
    fmt.Printf("Tamanho médio de palavra: %.2f\n", arquivo.TamanhoMedioDePalavra())
}

~~~
Resultado:
````
Média de palavras por linha: 8.46
Tamanho médio de palavra: 6.97
````
Convertemos `int` para `float64` no momento da divisão pois queremos que o resultado seja um número decimal. Sempre que
dividir dois `ints` o resultado produzido também será um `int`
e a parte decimal do número é truncada.

### 5.2- Interfaces

Uma interface é a definição de um conunto de métodos comuns a um ou mais tipo. É o que permite a criação de tipos
polimórficos em Go. Java possui um conceito muito parecido, também chamado de interface. A grande diferença é que em Go,
um tipo implementa uma interface implicitamente, bastando que este tipo defina todos os métodos desta interface, não
havendo necessidade de palavras chave como _implements_,
_extends_ etc.

Nesse exemplo vamos criar uma interface chamada `Operacao`, que define um único método `Calcular()`:
~~~go
type Operacao interface {
    Calcular()    int
}
~~~
Podemos agora criar um tipo chamado `Soma` contendo dois operandos, e implementar o método
`Calcular()`. Para facilitar a leitura dos resultados, vamos implementar o método `String()`. Este método é invocado
automaticamente quando queremos apresentar um valor utilizando o pacore `fmt`. Para implementá-lo, vamos utilizar a
função `Sprintf()` do próprio pacote
`fmt`, que retorna uma `string` formatada:
~~~go
type Soma struct {
    operando1, operando2 int
}

func (s Soma) Calcular() int {
    return s.operando1 + s.operando2
}

func (s Soma) String() string {
    return fmt.Sprintf("%d + %d", s.operando1, s.operando2)
}
~~~
Para utilizar este tipo na prática, poderia se simplismente declarado uma variável do tipo
`Soma`. Entretando podemos ver que a assinatura do método `Calcular()` do tipo `Soma` satisfaz a definição da
interface `Operacao`, e isto é suficiente para dizer que a `Soma` é uma
`Operacao`. Desta forma, podemos atribuir um valor `Soma` a uma variável definida como sendo do tipo `Operacao`:
~~~go
var soma Operacao
soma = Soma{10, 20}

fmt.Printf("%v = %d\n", soma, soma.Calcular())
~~~
Repare na forma que é utilizado o marcador `%v` para obter a representação `string` de uma
`Soma` neste caso `10 + 20`. Assim executando este código termos o resultado `10 + 20 = 30`. Para detalhar melhor vamos
implementar mais um tipo chamado `Subtracao`, que também implementa a interface `Operacao`, o exemplo completo esta
em `interface.go`.

Resultado:

````
10 + 20 = 30
30 - 15 = 15
10 - 50 = -40
5 + 2 = 7
Valor acumulado = 12
````

### 5.3- Duk typing e polimorfismo

_Duck typing_ (tipagem pato) é a capacidade de um sistema de tipos de determinar a semântica de um dado tipo baseado em
seus métodos e não em sua hierarquia. O nome tem origem do chamado duck test, se faz "quack" como um pato e anda como um
pato, então provavelmente é um pato. Isso é mais comum em linguagens dinamicamente tipadas como Ruby ou Python, porém
como Go não permite herança mas os tipos implementam interfaces implicitamente, na prática considera-se que Go tem uma
forma de _duck typing_.

Para facilitar o entendimento, vamos usar o exemplo anterior, porém dentro do arquivo
`duck_typing.go` e criar uma nova função com os dados antes no `main()`:
~~~go
func acumular(operacoes []Operacao) int {
    acumulador := 0
    for _, op := range operacoes {
        valor := op.Calcular()
        fmt.Printf("%v = %d\n", op, valor)
        acumulador += valor
    }
    return acumulador
}
~~~
Assim é possível reutilizar a função acumular passando como argumento um slice, de qualquer objeto que
implemente `Calcular()`, retornando um `int`. E adicionamos o novo tipo `Idade`:
~~~go
type Idade struct {
    anoNascimento int
}

func (i Idade) Calcular() int {
    return time.Now().Year() - i.anoNascimento
}

func (i Idade) String() string {
    return fmt.Sprintf("Idade desde %d", i.anoNascimento)
}

func main() {
    // ...	
	
    idades := make([]Operacao, 3)
    idades[0] = Idade{1969}
    idades[1] = Idade{1977}
    idades[2] = Idade{2001}
    
    fmt.Println("Idades acumuladas =", acumular(idades))
}
~~~
Resultado:
````
10 + 20 = 30
30 - 15 = 15
10 - 50 = -40
5 + 2 = 7
Valor acumulado = 12
Idade desde 1969 = 52
Idade desde 1977 = 44
Idade desde 2001 = 20
Idades acumuladas = 116
````