package main

import "fmt"

func main() {
	var preco float64
	var desconto float64

	fmt.Print("Informe o preco do produto")
	fmt.Scan(&preco)

	desconto = preco / 10
	resultado := preco - desconto
	fmt.Print("o novo preco do produto e ", resultado)
}
