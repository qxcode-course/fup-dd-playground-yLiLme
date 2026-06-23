package main

import (
	"bufio"
	"fmt"
	"os"
)
func main() {
    scanner:=bufio.NewScanner(os.Stdin)
    
    if scanner.Scan(){
        frase:=scanner.Text()
        tam:=len(frase)

        for i:=tam-1; i>=0; i--{
            fmt.Printf("%c", frase[i])
        }
        fmt.Println()
    }
}

