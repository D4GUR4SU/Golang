package main

import "fmt"

func main() {
	funcsPorLetra := map[string]map[string]float64{
		"G": {
			"Gabriela Silva": 154126.89,
			"Gugga Pereira":  12534.67,
		},
		"J": {
			"José João": 176298.00,
		},
		"P": {
			"Pedro Junior": 1200.00,
		},
	}

	delete(funcsPorLetra, "P")

	for letra, funcs := range funcsPorLetra {
		fmt.Println(letra, funcs)
	}
}
