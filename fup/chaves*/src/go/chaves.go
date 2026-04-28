package main
import "fmt"
func main() {
    var n int64

    fmt.Scan(&n)

    if n>0 {
        fmt.Println("positivo")
    } else if n<0 {
        fmt.Println("negativo")
    } else {
        fmt.Println("nulo")
    }
}
