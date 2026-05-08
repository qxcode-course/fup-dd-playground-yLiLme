package main
import "fmt"
import "math"
//programa que calcula se uma raiz é real
func main() {
    var a, b, c, del, x1, x2 float64

    fmt.Scan(&a, &b, &c)

        del=math.Pow(b,2)-4*(a*c)

    if del>0 {
        x1=(-b+math.Sqrt(del))/(2*a) //positivo
        x2=(-b-+math.Sqrt(del))/(2*a) //negativo
        fmt.Printf("%.2f\n%.2f\n", x1, x2)
    } else if del==0 {
        x1=(-b+math.Sqrt(del))/(2*a) //positivo
        x2=(-b-+math.Sqrt(del))/(2*a) //negativo
        if x1==0 {
            fmt.Printf("%.2f\n", x2)
        } else {
            fmt.Printf("%.2f\n", x1)
        }
    } else {
        fmt.Println("nao ha raiz real")
    }


}