package main

import "fmt"

func main() {
	var anos uint64
	const dias = 365

	fmt.Print("Qual e a sua idade? ")
	fmt.Scan(&anos)

	resultado := anos * dias
	fmt.Print("a conversao dos seus anos em dias vai ser ", resultado)

}
