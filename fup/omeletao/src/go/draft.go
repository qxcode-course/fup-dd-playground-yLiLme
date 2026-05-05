package main
import "fmt"
func main() {
    var a, b, c, d int64

    fmt.Scan(&a, &b, &c, &d)

    if a>b && a>c && a>d {
        fmt.Println(a)
    }
}
