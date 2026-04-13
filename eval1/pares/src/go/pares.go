package main
import "fmt"
func main() {
    var a, b, i, p=0 int64
    
    fmt.Scan(&a, &b)

    if a>b {
        fmt.Println("invalido")
    } else {
        for i=0 ; a<b ; i++ {
            if a%2==0 {
                p+=a
            } 
        } if b%2==0 {
            p+=b
        }
    }
            fmt.Println(p)
}
