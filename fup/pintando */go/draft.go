package main
import "fmt"
import "math"

func main() {
    var a, l1, l2, l3, p float64

    fmt.Scan(&l1, &l2, &l3)

    p=(l1+l2+l3)/2
    a=math.Sqrt(p*(p-l1)*(p-l2)*(p-l3))

    fmt.Printf("%.2f\n", a)

}
