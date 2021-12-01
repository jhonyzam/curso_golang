package main

import (
	"fmt"
)

func main() {
	c := make(chan int)

	go produzir(c)

	valor := <-c
	fmt.Println(valor)
}

func produzir(c chan int) {
	c <- 33
}
