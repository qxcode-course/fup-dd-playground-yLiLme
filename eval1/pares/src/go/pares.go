package main
import "fmt"

func main() {
    var a, b, p int

    fmt.Scan(&a, &b)
 
    if a>b {
        fmt.Println("invalido")
    } else {
        for i:=a; i<=b; i++ {
            if i%2==0 {
                p+=i
            }
        }
            fmt.Println(p)
    }

}
