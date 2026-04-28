package main
import "fmt"
import "math"
func main() {
    var a, b, c, del, x1, x2 float64
        fmt.Scan(&a, &b, &c)

        del=math.Sqrt(b)-4*a*b

        x1=(-b+math.Sqrt(del))/(2*a)
        x2=(-b-math.Sqrt(del))/(2*a)

        if del>0 {
            fmt.Printf("%.2f\n%.2f", x1, x2)
        }
}
