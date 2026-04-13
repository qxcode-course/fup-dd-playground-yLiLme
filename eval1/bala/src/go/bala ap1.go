package main
import "fmt"
import "math"

func main() {
    var x1, x2, y1, y2, dab, xx, yy float64

    fmt.Scan(&x1, &y1, &x2, &y2)

    xx=(x2-x1)*(x2-x1)
    yy=(y2-y1)*(y2-y1)
 
     dab=math.Sqrt(xx+yy)

    fmt.Printf("%.2f\n", dab)

}
