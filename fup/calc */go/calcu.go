package main
import "fmt"
func main() {
    var n1, n2, s, su, m, d int64
    var op rune

    fmt.Scan(&n1, &n2)
    fmt.Scanf("%c", &op)

    switch op {
    case '+':
        s=n1+n2
        fmt.Println(s)
    case '-':
        su=n1-n2
        fmt.Println(su)
    case '*':
        m=n1*n2
        fmt.Println(m)
    default:
        d=n1/n2
        fmt.Println(d)
    }
}
