package main

import "fmt"

func swap(ptr *int, ptra *int) {
	var a int = *ptra
	*ptra = *ptr
	*ptr = a
}

func main() {
	a := 10
	b := 20

	fmt.Println("Antes da troca: a=", a, ", b=", b)
	swap(&a, &b)
	fmt.Println("Depois da troca: a=", a, ", b=", b)
}
