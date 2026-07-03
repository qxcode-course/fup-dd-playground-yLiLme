package main

import (
	"bufio"
	"fmt"
	"os"
)

func numOco(frase string, letra string) int {
	var tOco int
	for i := 0; i < len(frase); i++ {
		if frase[i] == letra[0] {
			tOco++
		}
	}
	return tOco
}

func main() {
	var f, l string
	leitor := bufio.NewScanner(os.Stdin)
	leitor.Scan()
	f = leitor.Text()
	leitor.Scan()
	l = leitor.Text()
	fmt.Println(numOco(f, l))
}
