package main
import "fmt"
func main() {
    var c1, l1, c2, l2, d1, d2 int64

    fmt.Scan(&c1, &l1, &c2, &l2)

    d1=c1*l1
    d2=c2*l2

    if d1>d2{
        fmt.Println(d1)
    } else {
        fmt.Println(d2)
    }
}
