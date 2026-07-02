package main

import (
	"fmt"
)

type Jogadas struct {
	a, b int
}

func calcPontos(jogada Jogadas) (bool, int) {
	if jogada.a < 10 || jogada.b < 10 {
		return false, 0
	}

	var d int
	if jogada.a > jogada.b {
		d = jogada.a - jogada.b
	} else {
		d = jogada.b - jogada.a
	}
	return true, d
}

func BestJogada(jogada []Jogadas) int {
	b := -1
	mP := 99999

	for i := 0; i < len(jogada); i++ {
		v, p := calcPontos(jogada[i])

		if v && p < mP {
			mP = p
			b = i
		}
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n)

	jogs := make([]Jogadas, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&jogs[i].a, &jogs[i].b)
	}

	result := BestJogada(jogs)

	if result == -1 {
		fmt.Println("sem ganhador")
	} else {
		fmt.Println(result)
	}
}
