package main
import "fmt"
import "math"

func main() {
    var x1, y1, x2, y2, xx, yy, Dxy, XY float64

        fmt.Scan(&x1, &y1, &x2, &y2)

    xx=(x2-x1)*(x2-x1)
    yy=(y2-y1)*(y2-y1)

    Dxy=xx+yy
    XY=math.Sqrt(Dxy)

        fmt.Printf("%.2f\n", XY)

}