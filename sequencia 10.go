package main

import "fmt"

func main() {
	var peso float64
	fmt.Print("Qual seu peso em kg? ")
	fmt.Scan(&peso)

	resultado := peso * 2.2846
	fmt.Print("Seu peso em libras é ", resultado)

}
