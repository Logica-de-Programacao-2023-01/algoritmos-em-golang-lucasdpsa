package main

import "fmt"

func main() {
	var x int64
	var y int64
	var w int64

	fmt.Print("digite um numero ")
	fmt.Scan(&x)

	y = x - 1
	w = x + 1
	fmt.Print("o sucessor de ", x, " vai ser ", w, " e o antecessor sera ", y)

}
