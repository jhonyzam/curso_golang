# Simplificando o GO


Os exemplos são baseados no livro: "**Programando em Go - Crie aplicações com a linguagem do Google**"

###Como instalar o go 
* https://golang.org/doc/install


###Primeiro programa em go

Após a intalação concluída, basta criar um arquivo `ola.go` e digitar:

~~~go
package main

import "fmt"

func main() {
	fmt.Println("Olá, mundo")
}
~~~

Salve o arquivo vá até o diretori de instalação onde o arquivo foi criado e execute no console:
````
go run ola.go
````

###The Go Playground
Caso preferir não baixar o go, pode usar o playground para executar e se divertir

* https://play.golang.org/p/LaiIWrR2qWu