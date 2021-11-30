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
