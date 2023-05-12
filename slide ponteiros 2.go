package main

import "fmt"

func inverso(ptr *bool) {
	*ptr = !(*ptr)
}

func main() {
	b := true
	inverso(&b)
	fmt.Print(b)
}
