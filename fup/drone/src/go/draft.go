package main
import "fmt"
func main() {
    var a, b, c, h, l int64

    fmt.Scan(&a, &b, &c, &h, &l)

    if a<=h && b<=h && a<=l && b<=l {
        fmt.Println("S")
    } else if a<=h && c<=h && a<=l && c<=l {
        fmt.Println("S")
    } else if b<=h && a<=h && b<=l && a<=l {
        fmt.Println("S")
    } else if b<=h && c<=h && b<=l && c<=l {
        fmt.Println("S")
    } else if c<=h && a<=h && c<=l && a<=l {
        fmt.Println("S")
    } else if c<=h && b<=h && c<=l && b<=l {
        fmt.Println("S")
    } else {
        fmt.Println("N")
    }
}
