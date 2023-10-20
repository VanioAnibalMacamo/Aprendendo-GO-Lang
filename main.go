package main

import (
	"fmt"
)

var a string

func main() {
	resultado := soma(1, 1)
	fmt.Printf("%T", resultado)
}

func soma(a int, b int) int {
	return a + b
}
