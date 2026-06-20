package main
import "fmt"
func main() {
    var palavras string
    var chars int
    fmt.Scan(&chars)
    var frase[]string=make([]string, chars)

    for i := range frase{
        fmt.Scanf("%s", &palavras)
        frase[i]=palavras
    }
    for i := range frase{
        if frase[i]==" "{
            continue
        }
        if frase[i]=="a" || frase[i]=="e" || frase[i]=="i" || frase[i]=="o" || frase[i]=="u"{
            frase[i]="v"
        } else {
            frase[i]="c"
        }
    }

    fmt.Println(frase)
}