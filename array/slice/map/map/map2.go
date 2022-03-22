package main

import "fmt"

func main() {

	funcESalarios := map[string]float64{
		"José João":     11325.45,
		"Pedro Jão":     15000.45,
		"Débora Rayssa": 23000.45,
	}

	funcESalarios["Théo Henrique"] = 1350.0
	delete(funcESalarios, "Inexistente")

	for nome, salario := range funcESalarios {
		fmt.Println(nome, salario)
	}
}
