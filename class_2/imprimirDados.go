package main

import (
	"fmt"
)

func imprimirDados(nome string, idade int) {
	fmt.Printf("%s tem %d anos.", nome, idade)
}

func main() {
	imprimirDados("Jhonatan", 31)
}
