package main

import (
	"fmt"
)

func main() {
	var peso float64
	var altura float64
	var IMC float64

	fmt.Print("Qual e o seu peso? ")
	fmt.Scan(&peso)
	fmt.Print("Agora qual e a sua altura? ")
	fmt.Scan(&altura)

	IMC = peso / altura * altura

	if IMC < 18.5 {
		fmt.Print("o seu peso esta abaixo do normal ")
	} else if IMC >= 18.5 && IMC <= 24.9 {
		fmt.Print("o seu peso esta normal")
	} else if IMC >= 25 && IMC <= 29 {
		fmt.Print("voce esta acima do peso")
	} else if IMC >= 30 {
		fmt.Print("voce e a Thais Carla")
	}

}

