package main
import "fmt"
func main() {
    var a, b, s int64

    fmt.Scan(&a, &b)

    if b>=a {
        for i:=a; i<=b; i++ {
            if i%2==0 && i%3==0 {
                s+=i
            }   
        }
            fmt.Println(s)
    } else {
        fmt.Println("invalido")
    }    
}
