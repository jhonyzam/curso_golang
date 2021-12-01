# 6- Funções

Criamos diverças funções em diferentes contextos. Criamos até uma função recursiva 
quando implementamos o _quicksort_. Também ja vimos como criar objetos e adicionar métodos
que definem seu comportamento, agora vamos detalhar um pouco mais os recursos disponíveis em
Go para escrever funções.

### 6.1- O básico

Apenas para relembrar, utilizamos a palavra chave `func` para declarar uma nova função. Por
exemplo, uma função que não recebe nenhum argumento e não retorna nenhum valor pode ser 
declarada da seguinte forma:
~~~go
func ImprimirVersao() {
    fmt.Println("1.12")
}
~~~
Argumentos são declarados de maneira similar a variáveis, portanto, especificamos primeiro
seu nome, e depois seu tipo:
~~~go
func ImprimirSaudacao(nome string) {
    fmt.Printf("Olá, %s!\n", nome)
}
~~~
Múltiplos argumentos são separados por vírgulas:
~~~go
func ImprimirDados(nome string, idade int) {
    fmt.Printf("%s, %d anos.\n", nome, idade)
}
~~~
E, caso os argumentos sejam do mesmo tipo, podemos especificá-lo um única vez:
~~~go
func ImprimirSoma(a, b int, text string) {
    fmt.Printf("%s: %d\n", texto, a + b)
}
~~~
Para retornar um valor, especificamos seu tipo ao final da declaração e utilizamos a
palavra chave `return` para retorná-lo:
~~~go
func Soma(a, b int) int {
    return a + b
}
~~~
Podemos também especificar múltiplos valores de retorno, de forma bastante similar a lista
de argumentos:
~~~go
func PrecoFinal(precoCusto float64) (float64, float64) {
    fatorLucro := 1.33
    taxaConversao := 2.34
    
    precoFinalDolar := precoCustoDolar * fatorLucro
    
    return precoFinalDolar, precoFinalDolar * taxaConversao
}
~~~

### 6.2- O Retorno nomeados

Os valores retornados por uma função podem ser nomeados no momento da declaração, fazendo
com que eles estejam disponíveis como variáveis no corpo da função. Desta forma, podemos 
utilizar a palavra chave `return` sem especificar nenum valor. Os valores atuais nas variáveis
decladas na definição da função serão retornados:
~~~go
func PrecoFinal(precoCusto float64) (precoDolar float64, precoReal float64) {
    fatorLucro := 1.33
    taxaConversao := 2.34

    precoDolar = precoCusto * fatorLucro
    precoReal =  precoDolar * taxaConversao
    
    return
}

func main() {
    precoDolar, precoReal := PrecoFinal(34.99)
    
    fmt.Printf("Preço final em dólar: %.2f\n Preço final em reais: %.2f\n", precoDolar, precoReal)
}
~~~
No exemplo anterior atribuímos os valores as variáveis `precoDolar` e `precoReal`
deretamente utilizando o operador `=`, pois estas foram declaradas na definição da função.

Nomear os valores de retorno é uma ótima maneira de documentar uma função. No exemplo
modificado, definimos que a função retorna dois valores específicos, tornando sua intenção 
mais clara. Porém, utilizar a palavra chave `return` sem especificar os valores retornados
dificulta o entendimento do código. Desta forma recomenda-se sempre especificar os valores 
retornados, mesmo que eles já estejam nomeados na assinatura da função.

Como exemplo final temos o arquivo `precos.go` com o conteúdo completo, alterando apenas
o retorno da função:
~~~go
func PrecoFinal(precoCusto float64) (precoDolar float64, precoReal float64) {
    
    //... 
	
    return precoDolar, precoReal
}
~~~
Resultado:
````
Preço final em dólar: 46.54
Preço final em reais: 108.90
````

### 6.3- Argumentos variáveis

Uma função pode receber um número variável de argumentos. Funções deste tipo são conhecidas
em Go como _variadic functions_. Já utilizamos algumas delas, por exemplo `fmt.Printf()` e
`append()`. Para criar uma _variadic functions_, devemos preceder o tipo do último (ou único)
argumento com reticências. Na prática, as reticências indicam que a função pode receber zero
ou mais argumentos do tipo especificado.

Vamos definir uma função que recebe uma lista variável de nomes de arquivos e criar cada um 
deles em um diretório temporário. Para isto, vamos utilizar algumas funções presentes no 
pacote `os`:
~~~go
func CriarArquivos(dirBase string, arquivos ...string) {
    for _, nome := range arquivos {
        caminhoArquivo := fmt.Sprintf("%s/%s.%s", dirBase, nome, "txt")
    
        arq, err := os.Create(caminhoArquivo)
    
        defer arq.Close()
    
        if err != nil {
            fmt.Printf("Erro ao criar arquivo %s: %v\n", nome, err)
            os.Exit(1)
        }
    
        fmt.Printf("Arquivo %s criado.\n", arq.Name())
    }
}
~~~
Note o uso das reticências na declaração do argumento `arquivos ...string`. A implementação
da função é trivial. Percorremos a lista de nomes de arquivos recebidas e, para cada nome 
de arquivo, compilamos seu caminho completo a partir do diretório especificado em `dirBase`
e utilizamos a função `os.Create()` para criá-lo em disco. Essa função recebe um caminho
completo de arquivo e cria um arquivo novo no caminho especificado. Caso o arquivo em questão já 
exista, será substituído pelo novo arquivo vazio.

Além de efetivamente criar o arquivo, a função `os.Create()` retorna um ponteiro para um
objeto do tipo `os.File` representando o descritor do arquivo criado. Este descritos pode
ser usado para relaizar outras operações no arquivo. Neste exemplo, porém, apenas chamamos
o método `Name()` para obter seu caminho completo.A função `os.Create()` também pode retornar um erro caso não seja possível criar o arquivo. Se
isto acontecer, a execução será interrompida e o programa será encerrado através da chamada 
da função `os.Exit()`.

Por fim, é improtante notar o uso da palavra chave `defer`. Ela é utilizada para instruir
o ambiente de execução Go a relaizar uma tarefa imediatamente antes de a função atual retornar.
É bastante comum utilizar `defer` para garantir que os recursos alocados pela função sejam
liberados ao final de sua execução. Neste exemplo, chamamos o método `Close()` no descritor 
do arquivo criado para ter certeza de que ele será devidamente fechado e liberado.<br />
Para exemplificar use o arquivo `arquivos.go` que usa a função `CriarArquivos()`:
~~~go
func main() {
    tmp := os.TempDir()
    
    CriarArquivos(tmp)
    CriarArquivos(tmp, "teste1")
    CriarArquivos(tmp, "teste2", "teste3", "teste4")
}
~~~

Resultado:
````
Arquivo /tmp/teste1.txt criado.
Arquivo /tmp/teste2.txt criado.
Arquivo /tmp/teste3.txt criado.
Arquivo /tmp/teste4.txt criado.
````

### 6.4- Funções anônimas

Durante o desenvolvimento de uma aplicação é muito comum encontrarmos problemas que exigem 
algum tipo de manipulação de textos (substituição de caracteres, processamento de templates, etc).

Uma das formas mais poderosas de resolver este tipo de problea é o uso de expressões regulares.
Go disponibiliza uma série de facilidades para trabalhar com elas no pacote `rexexp`, basta 
adicioná-lo à lista de imports, como no exemplo `regex_str.go`:
~~~go
package main

import (
    "fmt"
    "regexp"
)

func main() {
    texto := "Anderson tem 21 anos"
    expr := regexp.MustCompile("\\d")
    
    fmt.Println(expr.ReplaceAllString(texto, "3"))
}
~~~
Inicialmente criamos a expressão regular `expr` através da função `regexp.MustCompile()`.
Essa função recebe a expressão desejada em formato `string`, compila e retorna um ponteiro 
para o objeto do tipo `regexp.Regexp`, utilizado para interagir com a expressão compilada.

No exemplo a expressão regular `\d` (as primeira `\` é usada como escape dentro da string)
captura qualquer caractere numérico da `string`. Com a expressão regular compilada em mãos, 
chamamos seu método `ReplaceAllString()`, especificando a `string` que deverá ser processada,
seguida da `string` que vai substituir.

Resultado:
````
Anderson tem 33 anos
````
Um outro exemplo seria transformar a primeira letra de cada palavra da frase em maiúscula.
Para esse caso, objetos do tipo `regexp.Regexp` possuem um método similar ao `ReplaceAllString()`,
chamado `ReplaceAllStringFunc()` que, em vez de uma string, aceita uma função como segundo argumento.
Ela deve, por sua vez, receber uma `string` como argumento e retornar ou `string` 
transformada. Neste caso utilizamos a função `ToUpper()` para transformar para maiúsculas,
como no exemplo a seguir:
~~~go
package main

import (
    "fmt"
    "regexp"
    "strings"
)

func main() {
    expr := regexp.MustCompile("\\b\\w")
    texto := "antonio carlos jobim"
    
    processado := expr.ReplaceAllStringFunc(
        texto,
        func(s string) string {
            return strings.ToUpper(s)
        },
    )
    
    fmt.Println(processado)
}
~~~
Resultado:
````
Antonio Carlos Jobim
````
Repare que definimos a função que transforma a string na própria chamada à função
`ReplaceAllStringFunc()`. Como esta função não tem nome são conhecidas como funções 
anônimas. Também podemos armazenar uma função anônima em uma variável. Vamos usar o exempo
anterior adicionando esta técnica no arquivo `regex_func.go`:
~~~go
package main

import (
    "fmt"
    "regexp"
    "strings"
)

func main() {
    expr := regexp.MustCompile("\\b\\w")
    
    transformadora :=
        func(s string) string {
            return strings.ToUpper(s)
        }
    
    texto := "antonio carlos jobim"
    fmt.Println(transformadora(texto))
    fmt.Println(expr.ReplaceAllStringFunc(texto, transformadora))
}
~~~
Resultado:
````
ANTONIO CARLOS JOBIM
Antonio Carlos Jobim
````
Caso queira mais detalhes sobre os recursos do pacote `regexp` visite
https://pkg.go.dev/regexp .

### 6.5- Closures
Funçẽos definidas de forma anônima dentro de outra função herdam o contexto de onde elas 
foram criadas. Para demonstrar este fato, vamos criar uma implementação de sequência de
Fibonacci utilizando uma função anônima, no exemplo `fibonacci.go`:
~~~go
package main

import "fmt"

func main() {
    a, b := 0, 1
    
    fib := func() int {
        a, b = b, a+b
    
        return a
    }
    
    for i := 0; i < 8; i++ {
        fmt.Printf("%d ", fib())
    }
}
~~~
Declaramos e inicializamos as variáveis `a` e `b` com os valores `0` e `1`, respectivamente,
dentro da própria função `main()`. Em seguida, definimos a função anônima atribuindo a 
variável `fib` que caclula a seguência de Fibonacci. A função manipula diretamente as 
variáveis `a` e `b`, atualizando seus valores a cada chamada.Esta propriedade de uma função
poder manipular o contexto onde ela foi originalmente definida é conhecida como _closure_.

Resultado:
````
1 1 2 3 5 8 13 21
````

### 6.6- Higher-Order functions

Como já mencionado, funções em Go podem retornar outras funções ou recebê-las como argumentos.
As funções que possuem estas características são conhecidas como _higher-order functions_ e são
parte fundamental de qualquer linguagem de programação funcional. Apsera de Go não ser uma
linguagem funcional pura, ela possui algumas poderozas abstraoes por este paradigma.
Como exemplo vamos usar a função que calcula a sequência de Fibonacci do exemplo anterior, 
para uma função externa chamada `GerarFinacci()`, o novo exemplo esta no arquivo `cronometro.go`:
~~~go
func GerarFibonacci(n int) func() {
    return func() {
        a, b := 0, 1
        
        fib := func() int {
            a, b = b, a+b
            
            return a
        }
        
        for i := 0; i < n; i++ {
            fmt.Printf("%d ", fib())
        }
    }
}
~~~
É importante notar que, a função `GerarFibonacci()` retorna outra função, pois definimos o
retorno como do tipo `func()`, uma função sem argumentos e sem retorno. Em seguida vamos criar
outra função chamada `Cronometrar()`, que recebe como argumento uma função qualquer e calcula
seu tempo de execução:
~~~go
func Cronometrar(funcao func()) {
    inicio := time.Now()
    
    funcao()
    
    fmt.Printf("\nTempo de execução: %s\n\n ", time.Since(inicio))
}
~~~
Primeiro armazenamos o tempo atual `time.Now()` na variável `inicio`. Em seguida, chamamos 
a função recebida como argumento. A função passada é a `func()` retornada por `GerarFibonacci()`
que efetua a sequência de fibonacci. Por fim utilizamos a função `time.Since()` para
calcular o intervalor de tempo entre o `inicio` e o tempo atual:
~~~go
func main() {
	Cronometrar(GerarFibonacci(8))
	Cronometrar(GerarFibonacci(48))
	Cronometrar(GerarFibonacci(88))
}
~~~
Resultado:
````
1 1 2 3 5 8 13 21 
Tempo de execução: 28.124µs

1 1 2 3 5 8 13 21 34 55 ... 701408733 1134903170 1836311903 2971215073 4807526976 
Tempo de execução: 99.807µs

1 1 2 3 5 8 ... 420196140727489673 679891637638612258 1100087778366101931
Tempo de execução: 179.625µs 
````
Repare na apresentação do tempo de execução de cada funçã `28.123us` para 8 números,
`99.807µs` para 48 números e `179.625µs` para 88 números. A função `String()` do tipo 
`time.Duration` (tipo de retorno de `time.Since()`) decide automaticamente a unidade de
tempos utilizada de acordo com o tamanho do período representado. Mais detalhes do pacote 
time em https://pkg.go.dev/time.

### 6.7- Tipos de função

No exemplo anterior, criamos uma função `Cronometrar()` que recebe como argumento uma outra 
função. Esta, por sua vez não recebe nenhum argumento e não retorna nenhum valor. Algumas
vezes, no entanto precisamos receber como argumento funções mais complexas e flexíveis.

Uma das formas de obter tal flexibilidade é através da definição de tipos para funções
(_function type_). Considere uma fução de agregação que recebe uma lista de valores, um valor 
inicial e uma função agregadora:
~~~go
func Agregar(valores []int, inicial int, fn func(n, m int) int) int {
    agregado := inicial
    
    for _, v := range valores {
        agregado = fn(v, agregado)
    }
    
    return agregado
}
~~~
Utilizando-a como base, podemos escrever uma função para calcular a soma de uma série numérica:
~~~go
func CalcularSoma(valores []int) int {
    soma := func(n, m int) int {
        return n + m
    }
    
    return Agregar(valores, 0, soma)
}
~~~
E também podemos escrever uma função para calcular o produto de uma série numérica:
~~~go
func CalcularProduto(valores []int) int {
    multiplicacao := func(n, m int) int {
        return n * m
    }
    
    return Agregar(valores, 1, multiplicacao)
}
~~~
Podemos agora criar a função `main()` e utilizar as duas novas funções:
~~~go

func main() {
    valores := []int{3, -2, 5, 7, 8, 22, 32, -1}
    
    fmt.Println(CalcularSoma(valores))
    fmt.Println(CalcularProduto(valores))
}
~~~
Par esta série o porgrama imprimirá os valores `74` e `1182720` para soma e o produto,
respectivamente. Repare como o trecho `func(n, m int) int` aparece três vezes e causa 
confusão especialmente quando utilizado como argumento na definição da função `Agregar()`.

Vamos extrair esta declaração comum para um tipo de função chamado `Agregadora`:
~~~go
type Agregadora func(n, m int) int
~~~
Desta forma podemos ver no arquivo `agregar.go` o exemplo muito mais legível com a definição
da função `Agregar()` para o argumento `fn`. Além disso, `CalcularSoma()` e `CalcularProduto()`
não precisam sofre alteração, pois as funções que ambas passam como argumento satisfazem
a assinatura de `Agregadora`.

### 6.8- HTTP através de funções

Neste exemplo vamos utilziar um dos recursos mais importantes da biblioteca padrão do Go, para 
escrever um servidor HTTP. A maior parte dos recursos disponíveis para a escrita de servidores
HTTP pode ser encontrada no pacote `net/http` https://pkg.go.dev/net/http.

Nosso servidor será bem simples e apenas retornará a data atual do sistema. A forma mais
fácil de escrever este servidor em Go é através da função `http.HandleFunc()`, que recebe 
dois argumentos: uma `string` que é o padrão de URL atendida por este serviço, e por
último, uma função do tipo `http.HandlerFunc`, que define a seguinte assinatura:
~~~go
type http.HandlerFunc func(http.ResponseWriter, *http.Request)
~~~
A implementação do serviço é bastante simples, segue o exemplo no arquivo `servidor_tempo.go`:
~~~go
package main

import (
    "fmt"
    "net/http"
    "time"
)

func main() {
    http.HandleFunc("/tempo", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "%s", time.Now().Format("2006-01-02 15:04:05"))
    })
    
    http.ListenAndServe(":8080", nil)
}
~~~
É importante notar que, na prática, um servidor HTTP em Go é um programa como outro qualquer,
portanto, precisa definir um pacote `main()` como ponto de partida.

Inicialmente, utilizamos a função `http.HandleFunc()` para registrá-lo, especificando
que ele atenderá requisições recebidas na URL `/tempo`. Em seguida, especificamos uma
função anônima que satisfaz a assinatura definida pelo tipo `http.HandlerFunc`, recebendo
um `http.ResponseWriter` e um ponteiro para um objeto `http.Request`.

A implementação deste código apresenta algumas novidades, como a função `fmt.Fprintf()` para
imprimir a data e hora atual como resposta da requisição HTTP. Esta função é muito similar à
`fmt.Printf`, a diferença é que esta ja escreve ja saída padrão, enquando a `Fprintf()`
recebe como primeiro argumento um valor do tipo `io.Writer` onde a saída será escrita. O valor
recebido em `w` é do tipo `http.ResponseWriter` que satisfaz a interface `io.Writer`.

Para formatar a data e hora atual, utilizamos o método `Format()` definido pelo tipo
`time.Time`. Este método recebe como argumento uma `string` que define o formato desejado.
A data utilizada no método `Format()` não é aleatória, é uma data pré estabelecida que 
reconhece o formato recebedio.

Por fim, utilizamos a função `http.ListenAndServe()` para especificar  que o servidor
deverá aceitar conexões na porta `8080`. Esta função inicia um servidor HTTP, bloqueia 
a execução do programa e delega as conexões recebidas no endereço definido para os
serviços registrados. O segundo argumento é um valor do tipo `http.Handler`, neste caso, 
foi especificado como `nil` pois já registramos um serviço através do método 
`http.HandleFun()`.

Podemos executar o servidor `servidor_tempo.go`, caso tenha sido iniciado com sucesso, nenhuma
saída sera impressa. Para testar o serviço abra seu navegador e digite 
http://localhost:8080/tempo. O resultado é uma pagina contendo a data e hora atual.