package main
import "fmt"
func main() {
    var frase string
    var char byte

    for {
        n, err := fmt.Scanf("%c", &char)

        if n!= 1 || err != nil || char == '\n' {
            break
        }
        if char !='\n'{
            frase+=string(char)
        }
    }
    var trecho string
    fmt.Scan(&trecho)

    count := 0 
    tamFrase, tamTrecho := len(frase), len(trecho)

    for i:=0; i<=tamFrase-tamTrecho; i++{
        if frase[i:i+tamTrecho] == trecho {
            count++
        }
    }
    fmt.Println(count)
}