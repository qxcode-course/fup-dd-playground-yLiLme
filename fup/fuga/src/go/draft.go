package main
import "fmt"
func main() {
    var h, p, f, d int64

    fmt.Scan(&h, &p, &f, &d)

    for {
        f=f+d
        if f==16{
            f=0
        }
        if f==-1{
            f=15
        }
        if f==p{
            fmt.Println("N")
            break
        }
        if f==h{
            fmt.Println("S")
            break
        }
    }
}
