package main
import "fmt"
func main() {
    var c, m, tp int

    fmt.Scan(&c)

    for {
        fmt.Scan(&m)
        tp+=m
        
        if tp>=c*2 {
            fmt.Println("hora de partir")
            break
        } else if tp==0 {
            fmt.Println("vazio")
        } else if tp>=c{
            fmt.Println("lotado")
        } else {
            fmt.Println("ainda cabe")
        }
    }


}
