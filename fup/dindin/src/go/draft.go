package main

import "fmt"

func main() {
	var d int
	var s, t string

	fmt.Scan(&d)

	var vetSabor []string = make([]string, d)
	var vetTurno []string = make([]string, d)

	for i := 0; i < d; i++ {
		fmt.Scan(&s, &t)
		vetSabor[i] = s
		vetTurno[i] = t
	}

	qtdChocolate := 0
	qtdLimao := 0
	qtdManha := 0
	qtdTarde := 0

	for i := 0; i < d; i++ {
		// Testando o sabor
		if vetSabor[i] == "c" {
			qtdChocolate++
		} else if vetSabor[i] == "l" {
			qtdLimao++
		}

		if vetTurno[i] == "m" {
			qtdManha++
		} else if vetTurno[i] == "t" {
			qtdTarde++
		}
	}

	if qtdChocolate > qtdLimao {
		fmt.Println("c")
	} else if qtdLimao > qtdChocolate {
		fmt.Println("l")
	} else {
		fmt.Println("empate")
	}

	if qtdManha < qtdTarde {
		fmt.Println("m")
	} else if qtdTarde < qtdManha {
		fmt.Println("t")
	} else {
		fmt.Println("empate")
	}
}