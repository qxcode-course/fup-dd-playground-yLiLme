package main
import "fmt"
func main() {
    var n, p int
    fmt.Scan(&n)
    var fila[]int=make([]int, n)
    var par[]int

    for i:=0; i<n; i++ {
        fmt.Scan(&p)
        fila[i]=p
    }
    for i:= range fila {
        if fila[i]%2==0 {
            =fila[i]
        }
    }
}
