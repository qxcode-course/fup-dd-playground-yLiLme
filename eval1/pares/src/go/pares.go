package main
import "fmt"

func main() {
    var a, b, i, p int64
    p=0

    fmt.Scan(&a, &b)
 
    if a<b {
        for i=0 ; a<b ; i++ {
            if a%2==0 {
                p+=a
            } 
        } 
        if b%2==0 {
            p+=b
        }
    } else {
        fmt.Println("invalido")
    } 
            fmt.Println(p)
}
