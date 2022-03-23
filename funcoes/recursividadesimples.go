package main

import "fmt"

// Uma solução melhor seria... utilizando uint para retornar numeros positivos
func fatorial(n uint) uint {
	switch {
	case n == 0:
		return 1
	default:
		return n * fatorial(n-1)
	}
}

func main() {
	resultado := fatorial(6)
	fmt.Println(resultado)
}
