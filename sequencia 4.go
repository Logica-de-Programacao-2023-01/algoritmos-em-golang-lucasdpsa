package main

import "fmt"

func main() {
	var x float64
	var y float64
	var w float64

	fmt.Print("digite um numero ")
	fmt.Scan(&x)
	fmt.Print("digite um numero... de novo ")
	fmt.Scan(&y)
	fmt.Print("digite mais um numero.... e o ultimo ... prometo ")
	fmt.Scan(&w)

	resultado := (2*x + 3*y + 5*w) / 10
	fmt.Print("a media ponderada e ", resultado)
}
