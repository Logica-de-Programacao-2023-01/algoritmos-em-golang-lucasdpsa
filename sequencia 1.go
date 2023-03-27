package main

import "fmt"

func main() {

	var x int64
	var y int64
	var w int64

	fmt.Print("digite um numero, pfv ")
	fmt.Scan(&x)
	fmt.Print("digite outro numero ")
	fmt.Scan(&y)
	fmt.Print("agora digite mais outro... e o ultimo... pfv ")
	fmt.Scan(&w)

	resultado := x + y + w
	fmt.Print("a soma dos numeros que voce digitou agorinha pouco e ", resultado)

}
