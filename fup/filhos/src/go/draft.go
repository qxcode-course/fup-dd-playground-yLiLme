package main
import "fmt"
func main() {
    var fn, qf, sf int64

    fmt.Scan(&fn, &qf)

    sf=fn
    for i:=sf; i<qf; i++ {
        sf+=2
        fmt.Printf("%d\n", i)
    }
}
