package main

import "fmt"

func main() {
	var salario float64
	var aumento float64

	fmt.Print("Infome o seu atual salario ")
	fmt.Scan(&salario)

	aumento = salario / 15
	resultado := salario + aumento
	fmt.Print("o seu novo salario sera ", resultado)

}
