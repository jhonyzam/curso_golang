package main

import "fmt"

func soma(n, m int) int {
	return n + m
}

func main() {
	s := soma(3, 5)
	fmt.Println("A soma é:", s)
}
