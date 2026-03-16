package main
import "fmt"
func main() {
    var p1, p2, p3 float64
    var v1, v2, v3, din, sobrou float64
    
fmt.Scan(&p1, &p2, &p3)
fmt.Scan(&v1, &v2, &v3)
fmt.Scan(&din)

sobrou=din-((p1*v1)+(p2*v2)+(p3*v3))

fmt.Printf("%.2f\n", sobrou)


}
