package main

import (
	"fmt"
	"os"
	"strconv"
)

const valorDolarPorReal = 5.54

func main() {

	if len(os.Args) < 3 {
		fmt.Println("Uso: conversor <moeda> <valores>")
		os.Exit(1)
	}

	moedaOrigem := os.Args[1]
	valoresOrigem := os.Args[2:len(os.Args)]

	var moedaDestino string

	if moedaOrigem == "real" {
		moedaDestino = "dolar"
	} else if moedaOrigem == "dolar" {
		moedaDestino = "real"
	} else {
		fmt.Printf("%s não é uma moeda conhecida!", moedaDestino)
		os.Exit(1)
	}

	for i, v := range valoresOrigem {
		valorOrigem, err := strconv.ParseFloat(v, 64)
		if err != nil {
			fmt.Printf(
				"O valor %s na posição %d não é um número válido!\n", v, i)
			os.Exit(1)
		}

		var valorDestino float64

		if moedaOrigem == "real" {
			valorDestino = valorOrigem * valorDolarPorReal
		} else {
			valorDestino = valorOrigem / valorDolarPorReal
		}

		fmt.Printf("%.2f %s = %.2f %s\n", valorOrigem, moedaOrigem, valorDestino, moedaDestino)
	}

}
