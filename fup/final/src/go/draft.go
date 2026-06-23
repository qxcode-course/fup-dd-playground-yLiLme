package main
import "fmt"
func main() {
    var nt1, nt2, ntf, m int64
    fmt.Scan(&nt1, &nt2, &ntf)

    m=(nt1+nt2)/2

    if m>=7 {
        fmt.Println("aprovado")
    } else if m<4 {
        fmt.Println("reprovado")
    } else {
        m=(ntf+m)/2
            if m>=5 {
                fmt.Println("aprovado na final")
            }else {
                fmt.Println("reprovado na final")
            }
    }
}
