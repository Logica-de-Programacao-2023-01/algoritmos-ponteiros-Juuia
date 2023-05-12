package main

import "fmt"

type produto struct {
	nome    string
	preco   float64
	estoque int
}

func iproduto(p *produto, novopreco float64) {
	p.preco = novopreco
}

func main() {
	produto := produto{
		nome:    "cropped",
		preco:   49.70,
		estoque: 12,
	}
	fmt.Println("Preço antigo", produto.preco)
	novopreco := 59.70
	iproduto(&produto, novopreco)
	fmt.Println("Preço novo", produto.preco)
}
