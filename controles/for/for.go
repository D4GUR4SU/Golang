package main

import (
	"fmt"
	"time"
)

func main() {

	i := 1
	for i <= 10 { // não tem while em Go
		fmt.Println(i)
		i++
	}

	// for tradicional
	for i := 0; i < 20; i += 2 {
		fmt.Printf("%d", i)
	}

	// for tradicional com estruturas aninhadas
	fmt.Print("\nMisturando...")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Print("Par ")
		} else {
			fmt.Print("Ímpar ")
		}
	}

	// laço infinito
	for {
		fmt.Println("Pra sempre...")
		time.Sleep(time.Second * 5)
	}
}
