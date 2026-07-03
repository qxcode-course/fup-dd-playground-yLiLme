package main
import "fmt"
func main() {
    var frase string
    var char byte

    for {
        n, _ := fmt.Scanf("%c", char)

        if n != 1 || char == '\n' {
            break
        }
        if char != '\r' {
            frase += string(char)
        }
    }

    var antiga, nova string
    fmt.Scan(&antiga, &nova)
    tamFrase := len(frase)
    tamAntiga := len(antiga)

    for i:=0; i<tamFrase; i++ {
        if i+tamAntiga <= tamFrase && frase[i:i+tamAntiga] == antiga {
            fmt.Print(nova)
            i += tamAntiga
        } else {
            fmt.Printf("%c", frase[i])
            i++
        }
    }
        fmt.Println()
}