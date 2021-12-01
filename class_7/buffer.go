package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 3)

	go produzirBuffer(c)

	for valor := range c {
		fmt.Println(valor)
	}
}

func produzirBuffer(c chan int) {
	c <- 1
	c <- 2
	c <- 3

	close(c)
}
