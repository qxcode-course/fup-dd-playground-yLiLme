package main
import "fmt"
func main() {
    var chars int
    var letra rune
    
    fmt.Scanln(&chars)
    frase:=make([]string, chars)

    for i:= range frase {
        fmt.Scanf("%c", &letra)
        if letra=='\n'{
            fmt.Scanf("%c", letra)
        }
        frase[i]=string(letra)
    }
    fmt.Println(frase)
}

