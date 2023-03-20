package main

import "fmt"

func main() {
	var dias uint64
	const diaria = 2.10

	fmt.Print("Quantos dias voce trabalhou?")
	fmt.Scan(&dias)

	resultado := float64(dias) * diaria
	fmt.Print("o seu salario sera ", resultado)
}
