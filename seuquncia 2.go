package main

import "fmt"

func main() {
	var x int64

	fmt.Print("digite um numero, pfv ")
	fmt.Scan(&x)

	y := x * 2
	w := x * 3
	z := x * 4

	fmt.Print("o dobro, triplo e o quadruplo de ", x, " sao, respectivamente ", y, " , ", w, " e ", z)
}
