package main

import "fmt"

func trocar(p1, p2 int) (segundo, primeiro int) {
	primeiro = p1
	segundo = p2
	return // retorno limpo
}

func main() {
	x, y := trocar(3, 2)
	fmt.Println(x, y)

	segundo, primeiro := trocar(7, 1)
	fmt.Println(segundo, primeiro)
}
