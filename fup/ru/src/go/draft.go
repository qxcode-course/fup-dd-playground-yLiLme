package main

import (
	"bufio"
	"fmt"
	"os"
)

func tratamento(frase string) {
	for i := 0; i < len(frase); i++ {
		if frase[i] == 'a' || frase[i] == 'e' || frase[i] == 'i' || frase[i] == 'o' || frase[i] == 'u' {
			fmt.Printf("%c", frase[i])
		}
	}
	fmt.Println()
	for i := 0; i<len(frase); i++ {
		if frase[i] != 'a' && frase[i] != 'e' && frase[i] != 'i' && frase[i] != 'o' && frase[i] != 'u' && frase[i] != ' ' {
			fmt.Printf("%c", frase[i])
		}
	}
    fmt.Println()

}

func main() {
	var frase string

	leitor := bufio.NewScanner(os.Stdin)
	leitor.Scan()
	frase = leitor.Text()
	tratamento(frase)
}
