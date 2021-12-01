# 3- Mapas

Um `map` é uma coleção de pares chave/valor sem ordem definida, onde a chave não se repete e armazena 
um único valor, na próxima aula vamos entrar mais no detalhe do `map`. 

### 3.1- Colher estatística
O exemplo a seguir utiliza `map` para contar a frequência de palavras com as mesmas iniciais.

A função `main` apenas recebe palavras como já vimos anteriormente, e chama as funções `colherEstatisticas()`
e `imprimir()`:
~~~go
func main() {
    palavras := os.Args[1:]
    
    estatisticas := colherEstatisticas(palavras)
    
    imprimir(estatisticas)
}
~~~
Para armazenar os dados que queremos, vamos usar um `map[string]int`, a chave é do tipo `string` e o valor 
do tipo `int`. Para cada item, a chave será a letra inicial e o valor será a quantidade de repetições:
~~~go
func colherEstatisticas(palavras []string) map[string]int
~~~
Primeiro criamos o `map` que iŕa armazenar as estatísticas, para isso usamos a função `make()`, mesmo 
princípio da inicialização de slices:
~~~go
estatisticas := make(map[string]int)
~~~
Em seguida, extraímos a primeira letra (`palavra[0]`) da lista de `palavaras`, convertendo ela para maiúscula 
usando a função `ToUpper()` do pacote `strings`, assim evitamos distinção das letras:
~~~go
inicial := strings.ToUpper(string(palavra[0]))
~~~
Com a inicial em mãos, procuramos no mapa `estatisticas` uma entrada cuja chave seja igual a `inicial`.
Para fazer isso, é semelhante ao acesso de vlores em um slice, porém utilizamos a chave do valor entre 
colchetes. O resultado são dois valores, o primeiro o valor armazenado no `map`, e o segundo um `bool` indicando
se o valor existe ou não no `map`:
~~~go
contador, encontrado := estatisticas[inicial]
~~~
Feito isso, basta verificar se valor foi encontrado, se sim adiconamos mais um ao `contador`, caso
contrário será a primeira ocorrência da `inicial`. Finalizado a contagem ele efetua o retorno do
mapa `estatisticas`:
~~~go
if encontrado {
    estatisticas[inicial] = contador + 1
} else {
    estatisticas[inicial] = 1
}

return estatisticas
~~~
A função `imprimir()`, recebe o `map` de estatísticas, intera todas as entradas, imprimindo os valores
de cada inicial:
~~~go
func imprimir(estatisticas map[string]int) {
    fmt.Println("Contagem de palavras iniciadas em cada letra:")

    for inicial, contador := range estatisticas {
        fmt.Printf("%s = %d\n", inicial, contador)
    }
}
~~~

#### Execução
````
go run class_3/colherEstatistica.go Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.

Retorno:
Contagem de palavras iniciadas em cada letra:
L = 2
I = 2
D = 3
S = 2
A = 3
C = 1
E = 3
T = 1
U = 1
M = 1
````

### 3.2- Pilhas com structs
Podemos descrever structs como objetos estruturados, mais a frente estudaremos
melhor as structs. Estes objetos possuem seus atributos e seus métodos/funções.
Um exemplo seria uma struct com nome `carro`, um atributo ou field seria `portas`, e um método ou função,
agregado a essa struct seria `quantidadeDePortas()`.

No nosso exemplo vamos ter a nossa struct `Pilha`, com um slice denominado `valores` que armazena 
objetos do tipo `interface{}`. Este tipo permite com que seja armazenado qualquer tipo de dados 
válido do Go. Outro ponto a ressaltar é que o slice `valores` está declarado com a inicial em minúsculo,
dessa forma ele não é acessível em outro pacote diretamente:
~~~go
type Pilha struct {
    valores []interface{}
}
~~~
E como mencionamos é possível adicionar funções vinculadas a essa struct. Essas funções são muito semelhantes 
ao que vimos até agora, a principal diferença é o objeto receptor `pilha` do tipo `Pilha`, que deve ser 
especificado entre parênteses antes do nome da função. Assim essas funções podem acessar o slice `pilha.valores`: 
~~~go
func (pilha Pilha) Tamanho() int {
    return len(pilha.valores)
}

func (pilha Pilha) Vazia() bool {
    return pilha.Tamanho() == 0
}
~~~
As funções `Empilhar()` e `Desempilhar()` seguem o mesmo princípio de receptores, porém aqui o tipo `Pilha`
é passado como um ponteiro. Em Go, argumentos de funções e métodos são sempre passados como cópia. 
Por isso, quando precisamos alterar um valor e não recriá-lo usamos ponteiros. Basicamente, um ponteiro
contém o endereço de um valor na memória. 

A função `Empilhar()` usa o ponteiro `pilha` para adicionar objetos no slice `pilha.valores` através 
da função `append()`. Por ser um ponteiro, o slice sempre vai manter os dados na memória, e por isso apenas 
o valor é passado como argumento.

E o mesmo conceito serve para a função `Desempilhar()`, a diferença é que nessa função não possui passagem
de argumento, o valor é sempre o ultimo registro do slice:
~~~go
func (pilha *Pilha) Empilhar(valor interface{}) {
    pilha.valores = append(pilha.valores, valor)
}

func (pilha *Pilha) Desempilhar() (interface{}, error) {
    if pilha.Vazia() {
        return nil, errors.New("Pilha vazia!")
    }
    valor := pilha.valores[pilha.Tamanho()-1]
    pilha.valores = pilha.valores[:pilha.Tamanho()-1]
    
    return valor, nil
}
~~~
Com as funçãos explicadas vamos para a implementação delas na função `main`.

Primeiro, criamos um obejto a do tipo `Pilha{}`, e imprimimos o tamanho e a validação de vazia:
~~~go
pilha := Pilha{}

fmt.Println("Pilha criada com tamanho: ", pilha.Tamanho())
fmt.Println("Vazia ", pilha.Vazia())
~~~
Segungo, vamos adicionar valores na pilha, usando a função `Empilhar()`, e imprimimos seu novo
tamanho:
~~~go
pilha.Empilhar("Go")
pilha.Empilhar(2021)
pilha.Empilhar(3.1)
pilha.Empilhar("Fim")

fmt.Println("Tamanho atual ", pilha.Tamanho())
~~~
Terceiro, usamos um `for` do tipo `bool`. Como já mencionado Go possui apenas o `for` como laço de 
repetição, e neste caso ele executa enquanto o retorno de `pilha.Vazia()` for `false`. A cada repetição
a função `pilha.Desempilhar()` remove um registro da fila, assim que a `pilha` estiver vazia o laço será
completo:
~~~go
for !pilha.Vazia() {
    v, _ := pilha.Desempilhar()
    fmt.Println("Desempilhando", v)
    fmt.Println("Tamanho: ", pilha.Tamanho())
    fmt.Println("Vazia? ", pilha.Vazia())
}
~~~
Por fim, caso executar a função `pilha.Desempilhar()` sendo que já esteja vazia, imprimimos um erro:
~~~go
_, err := pilha.Desempilhar()
if err != nil {
    fmt.Println("Erro: ", err)
}
~~~
#### Execução
````
go run class_3/pilha.go

Resultado:
Pilha criada com tamanho:  0
Vazia  true
Tamanho atual:  4
Desempilhando: Fim
Tamanho:  3
Vazia?  false
Desempilhando: 3.1
Tamanho:  2
Vazia?  false
Desempilhando: 2021
Tamanho:  1
Vazia?  false
Desempilhando: Go
Tamanho:  0
Vazia?  true
Erro:  Pilha vazia!
````