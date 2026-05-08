package main
import "fmt"
func main() {
    var a, b, c int64

    fmt.Scan(&a, &b, &c)

    if b+c>a && a+c>b && b+a>c  {
        fmt.Println("True")
    } else {
        fmt.Println("False")
    }

}
