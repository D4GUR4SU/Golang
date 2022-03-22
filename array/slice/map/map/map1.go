package main

import "fmt"

func main() {
	aprovados := make(map[int]string)

	aprovados[912345678978] = "Maria"
	aprovados[512345678978] = "Pedro"
	aprovados[712345678978] = "Ana"
	fmt.Println(aprovados)

	for cpf, nome := range aprovados {
		fmt.Printf("%s (CPF: %d)\n", nome, cpf)
	}

	fmt.Println(aprovados[512345678978])
	delete(aprovados, 512345678978)
	fmt.Println(aprovados[512345678978])

}
