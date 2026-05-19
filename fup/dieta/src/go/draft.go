package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    var vet[] int = make([]int, n)

    var cal int
    var m, cCal float64
    for i:=0; i<n; i++ {
        fmt.Scan(&cal)
        vet[i]=cal
    }
    for i:= range vet {
        cCal+=float64(vet[i])
    }
        m=cCal/float64(n)
        fmt.Printf("%.1f\n", m)
}
