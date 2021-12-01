# 7- Concorrência com goroutines e channels

Programação concorrente é um tema bastante delicado e complexo. Se você já escreveu programas
que dependiam de múltiplas _threads_ ou mútiplos processos sabe que é uma tarefa que exige
muita atenção e paciência, especialmente quando é necessário compartilhar estado entre as 
diferentes linhas de execução.

A maior parte das linguagens de programação implementa _threads_ de alguma forma, seja através 
de um escalonador próprio ou delegando o controle para o sistema operacional hospedeiro. O
compartilhamento de estado em um programa _multi-threaded_ é normalmente implementado
através de variáveis globais e/ou compartilhadas e exige algum mecanismo de trava ou semáforo
para evitar condições de corrida (_race conditions_).

Para evitar este tipo de problema, Go implementa um modelo de concorrência baseado em 
_goroutines_ que se comunicam através de _channels_, sendo que o próprio ambiente de execução
garante que apenas uma _goroutine_ acesse um _channel_ em um determinado momento.

### 7.1- Goroutines

Uma _goroutine_ é um tipo de processo extremamente leve. Na prática, uma goroutine é muito
similiar a uma thread. No entanto, go routines são gerenciadas pelo ambiente de execução da
linguagem, que decide quando e como associálas a threads do sistema operacional. Para
iniciar uma goroutine, utilizamos a palavra-chave `go` seguida de uma chamada de função. O
ambiente de execução irá executar a função chamada sem bloquear a linha de execução principal.

Considere a seguinte função `imprimir()`, que recebe um número inteiro e o imprime três vezes
com um espaço de 200 milissegundos entre cada impressão, exemplo do código `goroutines_1.go`:
~~~go
func imprimir(n int) {
    for i := 0; i < 3; i++ {
        fmt.Printf("%d ", n)
        time.Sleep(200 * time.Millisecond)
    }
}
~~~
Agora considere a seguinte função `main()` que chama a função `imprimir()` duas vezes
com números diferentes:
~~~go
func main() {
    imprimir(2)
    imprimir(3)
}
~~~
Executando esta função obteriamos `2 2 2 3 3 3` como resultado. Vamos alterar a primeira
chamada para que seja executada em uma goroutine:
~~~go
func main() {
    go imprimir(2)
    imprimir(3)
}
~~~
Através da simples adição da plavra-chave `go`, esta nova versão executa as duas chamadas
de forma concorrente e produz um resultado diferente: `3 2 2 3 3 2`.
É importante ressaltar que as goroutines dependem da função `main()` para que continuem 
sua execução. Em outras palavras, as goroutines **morrem** quando a função `main()` finaliza
sua execução.

Podemos provar este fato com o simples programa a seguir, que inicia uma goroutine que 
dorme por 5 segundos, enqunto a função `main()` dorme por apenas 3 segundos:
~~~go
func dormir() {
    fmt.Println("Goroutine dormindo por 5 segundos...")
    time.Sleep(5 * time.Second)
    fmt.Println("Goroutine finaizada.")
}

func main() {
    go dormir()
    
    fmt.Println("Main dormindo por 3 segundos...")
    time.Sleep(3 * time.Second)
    fmt.Println("Main finaizada.")
}
~~~
O resultado da execução deste programa seria o seguinte:
````
Main dormindo por 3 segundos...
Goroutine dormindo por 5 segundos...
Main finaizada.
````
A goroutine executando a função `dormir()` nunca termina, pois a função `main()` termina
sua execução antes e o programa é finalizado.

### 7.2- Channels

A capacidade de executar diferentes goroutines concorrentemente é muito importante e abre
diversas possibilidades para a solução de problemas que exigem alta performance e melhor
eficiência no uso de recursos de processamento. Entretanto, é muito rara a situação e que 
várias goroutines são disparas independentemente, sem que haja comunicação entre elas. Os 
_channels_ foram criados como uma abstração para viabilizar esta comunicação.

Um _channel_ é um canal, utilizamos a função `make()`, para criar um canal capaz de trafegar 
valores to tipo `int` no seguinte exemplo:
~~~go
c := make(chan int)
~~~
Para interagir com um canal, utilizamos o operador `<-` (conhecido como _arrow operator_, ou
operador seta). A posição do canal em relação à seta indica a direção do fluxo da comunicação.
Por exemplo, para enviar valores `int` para o canal `c`, utilizamos a seguinte notação:
~~~go
c <- 33
~~~
E para receber um valor enviado para o canal `c`:
~~~go
valor := <-c
~~~
A seguir temos um exemplo simples que combina todos os passos anteiores para demonstrar
o fluxo de comunicação completo, no arquivo `channels.go`:
~~~go
func main() {
    c := make(chan int)
    
    go produzir(c)
    
    valor := <-c
    fmt.Println(valor)
}

func produzir(c chan int) {
    c <- 33
}
~~~
Inicialmente criamos um canal pra trafegar valores do tipo `int`. Em seguida, disparamos uma 
goroutine executando a função `produzir()`, que recebe um canal como argumento e simplismente
envia um número inteiro para o canal recebido.

Por padrão, operações de envio e recebimento em um canal bloqueiam até que o outro lado
esteja pronto. Este fato permite que a própria comunicação entre duas goroutines garanta a
sincronização entre elas, sem que nenhum mecanismo de travas seja necessário. Por este
motivo, a próxima linha da função `main()` que recebe um valor do canal, fará com que a 
linha de execução principal fique bloqueada até que algum valor seja enviado para o canal `c`.
Assim que o valor `33` for enviado pela função `produzir()`, a linha de execução principal
será então desbloqueada, o valor `33` será consumido, atribuído à variável `valor` e impresso
no console.

### 7.3- Buffers

Canais podem ser criados com um buffer, por exemplo:
~~~go
c := make(chan int, 5)
~~~
O canal `c` foi criado com um buffer de tamanho 5. Isto quer dizer que operações de envio
não serão bloqueadas enquanto o buffer não estiver cheio, e operações de recebimento não
serão bloqueadas enquanto o buffer não estiver vazio.
~~~go
func main() {
    c := make(chan int, 3)
    
    go produzirBuffer(c)
    
    fmt.Println(<-c, <-c, <-c, <-c)
}

func produzirBuffer(c chan int) {
    c <- 1
    c <- 2
    c <- 3
}
~~~
Criamos um canal com buffer de tamanho 3 e imediatamente disparamos uma goroutine
que envia três valores pelo canal criado. Em seguida, recebemos quatro valores do canal e os 
imprimimos no console, o resultado:
````
fatal error: all goroutines are asleep - deadlock!
````
Sim, um _deadlock_. A goroutine produtora encerrou sua execução logo após produzir os três
valores. Por causa do buffer, nenhuma das operações de envio fez com que a execução fosse 
bloqueada. No entanto ao receber um quarto valor que nunca foi produzido pelo canal, a 
linha principal ficou bloqueada. O ambiente de execução detectou o _deadlock_ e encerrou
a execução do programa com um erro. 

Para evitar isso o produtor precisa indicar de alguma forma que não enviará mais nenhum 
valor pelo canal. Para isso existe a função embutida `close()`:

~~~go
func produzirBuffer(c chan int) {
    c <- 1
    c <- 2
    c <- 3
    
    close(c)
}
~~~
É muito importante que o lado produtor feche o canal sempre que não houver mais valores
a serem produzidos. Agora, precisamos detectar que o canal foi fechado do lado consumidor.
O operador `<-` retorna sempre dois valores: o valor lido e um valor `bool` indicando se o 
valor foi lido com sucesso ou não, este valor será `false` quando o canal for fechado:

~~~go
func main() {
    c := make(chan int, 3)
    
    go produzirBuffer(c)
    
    for {
        valor, ok := <-c
        if ok {
            fmt.Println(valor)
        } else {
            break
        }
    }
}
~~~
Resultado:
````
1
2
3
````
Desta forma resolvemos o problema do _deadlock_, porém este processo de checar o canal pode 
ser melhorado usando o operador `range` para ler os valores de um canal. Desta forma conforme
eles são produzidos os valores ja são lidos automaticamente, o código completo do exemplo está 
em: `buffer.go`:
~~~go
func main() {
    c := make(chan int, 3)
    
    go produzirBuffer(c)
    
    for valor := range c {
        fmt.Println(valor)
    }
}
~~~
Resultado:
````
1
2
3
````
Assim o código fica mais claro e temos o mesmo resultado.

### 7.4- Controlando a direção do fluxo

Por padrão, a comunicação em um canal é bidirecional. Algumas vezes, porém, desejamos 
controlar a direção do fluxo quando passamos um canal como argumento para outra função, 
ou quando temos uma função que retorne um canal.

Vamos utilizar o exemplo anterior vamos alterar a função `produzir()` para definir a
direção da comunicação:
~~~go

func produzirBuffer(c chan<- int) {
    // ...
}
~~~
Repare que, ao receber o canal como argumento, definimos que  função poderá somente enviar
valores para o canal (`chan<-`). Desta forma, caso a função tente receber valores pelo 
mesmo canal, causará o erro:
````
invalid operation: <-c (receive from send-only type chan<- int)
````
De maneira similar, podemos definir um canal _read-only_ (somente leitura):
~~~go
func consumirBuffer(c <-chan int) {
    // ...
}
~~~
Como Go é uma linguagem fortemente tipada e os canais são tipados, esse erro é causado em 
tempo de compilação.

### 7.5- Select
É muito comum encontrar um cenário em que uma goroutine precisa interagir com últiplos canais
de comunicação. Por exemplo, uma função pode disparar goroutines que escrevem valores em 
canais diferentes, e a função original depende dos valores de todos estes canais para 
produzir seu resultado final.

Para evitar que a execução de uma goroutine seja bloqueada esperando por operações em
algum dos canais dos quais ela depende, Go fornece um comando chamado `select`, que é
muito semelhante ao `switch`.Sua forma geral é:
~~~go
select {
case v1 := <-canal1:
	// ...
case v2 := <-canal2:
    // ... 
default:
    // ... 
}
~~~
Caso exista algum valor a ser lido no `canal1`, ele será atribuído à variável `v1` e o
boco associado ao primeiro comando `case` será executado. Caso contrário, o `canal2` será
checado por valores recebidos e, em caso positivo, seu valor será atribuído à variável `v2`
e o bloco associado ao segundo comando `case` será executado. Se nenhum dos canais 
especificados possuírem valores prontos para serem lidos, o bloco `default` será executado.

Uma das formas mais comuns do uso do `select` é dentro de um laço que controla a comunicação
com as goroutines. Vamos escrever um programa que, dada uma lista de números, separa-os em 
duas listas de pares e ímpares.

Primeiro, vamos criar a função `separar()`, que recebe a lista de números e três canais: 
`i`, `pp` e `pronto`, todos unidirecionais:
~~~go
func separar(nums []int, i, p chan<- int, pronto chan<- bool) {
    for _, n := range nums {
        if n%2 == 0 {
            p <- n
        } else {
            i <- n
        }
    }
    pronto <- true
}
~~~
Para cada número presente em `nums`, verificamos se é par e, em caso positivo, enviamos o
número para o canal `p`, caso seja um número ímpar, ele será enviado para o canal `i`. Ao
final da iteração, enviamos o valor `true` para o canal `pronto`, indicando o fim do 
processamento. Agora na função `main()` criamos os canais e listas de números:
~~~go
i, p := make(chan int), make(chan int)
pronto := make(chan bool)

nums := []int{1, 23, 42, 5, 8, 6, 7, 4, 99, 100}
~~~
Em seguida, disparamos uma _goroutine_ para separar os números:
~~~go
go separar(nums, i, p, pronto)
~~~
Com tudo preparado, precisamos coletar os resultados evniados para cada canal. Desta forma 
criamos duas listas `impares` e `pares`, populando-as com um laço até que o valor da variável
`fim` seja verdadeiro:
~~~go
var impares, pares []int
fim := false

for !fim {
    select {
    case n := <-i:
        impares = append(impares, n)
    case n := <-p:
        pares = append(pares, n)
    case fim = <-pronto:
    }
}
~~~
Repare que a última cláusula `case:` só é executada quando a goroutine enviar o valor
`true` para o canal `pronto` e, assim a variável `fim` assume o valor `true`, atendendo
a condição de saída do laço.

Por fim, imprimimos os valore separados por lista:
~~~go
fmt.Printf("Ímpares: %v | Pares %v\n", impares, pares)
~~~
Resultado:
````
Ímpares: [1 23 5 7 99] | Pares [42 8 6 4 100]
````

### 7.6- Temporizadores e timeouts

Controlar a execução de múltiplas goroutines através de um _select_ é uma técnica bastante
simples e muito utilizada em Go. Porém, existem casos em que a execução de uma determinada
tarefa precisa aontecer em um príodo limitado de tempo.

O pacote `time` fornece uma função `After()` que ajuda a resolver estas casos. Sua 
assinatura é:
~~~go
func After(d Duration) <-chan Time
~~~
Esta função convenientemente retorna um canal. Assim, podemos facilmente utilizá-la dentro
de um _select_ para controlar o tempo de execução de uma goroutine, e tomar alguma
atitude caso ela não produza resultados dentro do tempo esperado.

Vamos simular esta situação com uma goroutine que dorme por 5 segundos e sinaliza o final
de sua execução enviando o valor `true` para um canal, exemplo `timeout.go`:
~~~go
func executar(c chan<- bool) {
    time.Sleep(5 * time.Second)
    c <- true
}

func main() {
    c := make(chan bool, 1)
    
    go executar(c)
    
    fmt.Println("Esperando...")
    
    fim := false
    for !fim {
        select {
        case fim = <-c:
            fmt.Println("Fim!")
        case <-time.After(2 * time.Second):
            fmt.Println("Timeout!")
            fim = true
        }
    }
}
~~~
Podemos notar a condição `case <-time.After(2 * time.Second):` dentro do `select` com 
isso montamos um mecanismode _timeoutr_ utilizando o `time.After()`, como podemos ver na função
`executar()` o `time.Sleep` está de 5 segundos, logo o timeout é apresentado.
Resultado:
````
Esperando...
Timeout!
````
Caso o `time.Sleep()` fosse de 1 segundo o resultado seria:
````
Esperando...
Fim!
````

### 7.7- Sincronizando múltiplas goroutines

Anteriormente, vimos um exemplo de como esperar que uma goroutine finalize a sua execução
através da utilização de um canal. No entanto, quando precisamos esperar pela execução de 
múltiplas goroutines, controlar manualmente quantas delas já terminaram pode ser uma 
tarefa sujeita a falhas. Para isso Go fornece um tipo chamado `WaitGroup`, presente no 
pacote `sync`, vamos ver o exemplo `sincronizador.go`:
~~~go
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

func main() {
    inicio := time.Now()
    
    rand.Seed(inicio.UnixNano())
    
    var controle sync.WaitGroup
    
    for i := 0; i < 5; i++ {
        controle.Add(1)
        go executarTimeSleep(&controle)
    }
    
    controle.Wait()
    
    fmt.Printf("Finalizando em %s.\n", time.Since(inicio))
}
~~~
Inicialmente, armazenamos o _timestamp__ atual e configuramos a semente para a geração
de números aleatórios, garantindo que números diferentes sejam gerados a cada execução.

Em seguida, criamos uma variável chamada `controle` do tipo `sync.WaitGroup` que
utilizaremos para sincronizar a execução das goroutines. Antes de disparar cada goroutine,
chamamos o método `control.Add(1)`, indicando que uma nova goroutine deverá ser sincronizada,
e então iniciamos sua execução chamando `executar()`, passando um ponteiro da variável controle.
Lembrando que argumentos em Go são passados como cópia, logo é importante que seja passado como 
ponteiro, pois o mesmo `WaitGroup` vai ser usado por todas as goroutines.

Na sequência chamamos o método `controle.Wait()`, que nesta caso bloqueia a execução da 
função `main()` até que todas as goroutines tenha finalizado.
~~~go
func executarTimeSleep(controle *sync.WaitGroup) {
    defer controle.Done()
    
    duracao := time.Duration(1+rand.Intn(5)) * time.Second
    fmt.Printf("Dormindo por %s... \n", duracao)
    time.Sleep(duracao)
}
~~~
Na função `executarTimeSleep()` garantimos que o método `controle.Done()` será chamado quando
a função terminar sua execução, notificando o `WaitGroup` deste fato. Após criamos um 
valor do tipo `time.Duration` baseado em um número aleatório entre 1 e 5, e chamamos a 
função `time.Sleep()` para que a goroutine atual durma por esse tempo.
Resultado:
````
Dormindo por 5s... 
Dormindo por 4s... 
Dormindo por 5s... 
Dormindo por 2s... 
Dormindo por 2s... 
Finalizando em 5.000208796s.
````
Vamos demonstrar um segundo exemplo utilizando um conceito de _threads_ muito legal, usando
o exemplo `mult_threads.go`. A intenção desse código é simular uma captura de dados, e em 
paralelo uma execução de regras.

O importante de ressaltar nesse código é o seguinte:
~~~go
var wg sync.WaitGroup
qtThreads := 1
threads := make(chan bool, qtThreads)

a := []string{"valor1", "valor2", "valor3", "valor4"}
ac := make(chan *string, len(a))
acReturn := make(chan *[]string, 1)

for _, str := range a {
    wg.Add(1)
    go CapturarDados(str, &wg, &threads, ac)
}

go ExecutarDados(ac, acReturn)

wg.Wait()
~~~
Temos uma variável chamada `qtThreads` ela vai ser usada para estruturarmos quantas 
goroutines queremos em paralelo. Criamos um slice `a` de `string`, para simular uma base de dados
e um channel `ac` para compartilhamento entre as goroutines, e o channel `acReturn` apenas
para imprimir o valor final.

Na sequencia montamos um loop percorrendo todos os valores do slice e executando a função
`CapturarDados()`:
~~~go
func CapturarDados(s string, wg *sync.WaitGroup, threads *chan bool, ac chan *string) {
    defer func() {
        wg.Done()
        <-*threads
    }()
    
    *threads <- true
    
    time.Sleep(4 * time.Second)
    ac <- &s
}
~~~
Esta função é pra simular uma busca em uma base dados, ou algo do tipo. O mais legal é o
uso da variável `threads`, pois vamos disparar varias goroutines mas elas vão ficar em espera
e executar conforme esse canal for sendo liberado.

Por fim temos função `ExecutarDados()`, que vai ficar validando se o canal esta aberto
populando os valores para iprimir ao final:
~~~go
func ExecutarDados(ac chan *string, acReturn chan *[]string) {
    var ac2 []string
    cc := 0
    for {
        a, open := <-ac
        if open {
            ac2 = append(ac2, *a)
            cc++
            fmt.Printf("channel %d \n", cc)
        } else {
            acReturn <- &ac2
            break
        }
    }
}
~~~
Resultado 1 thread:
````
Capturando...
channel 1 
Capturando...
channel 2 
Capturando...
channel 3 
Capturando...
channel 4 
&[valor1 valor2 valor3 valor4]
FINALIZOU
````
Resultado 2 thread:
````
Capturando...
Capturando...
channel 1 
channel 2 
Capturando...
channel 3 
channel 4 
Capturando...
&[valor2 valor1 valor3 valor4]
FINALIZOU
````
Repare como no resultado com 1 thread, ele executa sequencial, enquanto no exemplo
com 2 threads ele dispara duas execuções juntas, e depois segue a fila conforme esta
libere espaço de acordo com as threads em paralelo.