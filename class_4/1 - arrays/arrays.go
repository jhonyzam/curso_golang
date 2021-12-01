package main

import "fmt"

func main() {
	var a [3]int
	fmt.Println("Array a: ", a)

	numeros := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array numeros: ", numeros)

	primos := [...]int{2, 3, 5, 7, 11, 13}
	fmt.Println("Array primos: ", primos)

	nomes := [2]string{}
	fmt.Println("Array nomes: ", nomes)

	var multiA [2][2]int
	multiA[0][0], multiA[0][1] = 3, 5
	multiA[1][0], multiA[1][1] = 7, 2
	fmt.Println("Array multiA: ", multiA)

	multiB := [2][2]int{{2, 13}, {-1, 6}}
	fmt.Println("Array multiB: ", multiB)
}
