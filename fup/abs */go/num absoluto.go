package main
import "fmt"

func main() {
    var n1, n2, vabs int64

        fmt.Scan(&n1, &n2)
 
    vabs=n1-n2

    if vabs<0 {
        vabs=vabs*(-1)
        fmt.Println(vabs)
    } else{
        fmt.Println(vabs)
    }
}
