package main
import "fmt"
    //a>b>c
    //b>c>a
    //c>a>b
func main() {
    var a, b, c int64

    fmt.Scan(&a, &b, &c)
    
    if a>b && b>c || c>b && b>a {
        fmt.Println(b)
    } else if b>c && c>a || a>c && c>b{
        fmt.Println(c)
    } else if c>a && a>b || b>a && a>c {
        fmt.Println(a)
    }
}
