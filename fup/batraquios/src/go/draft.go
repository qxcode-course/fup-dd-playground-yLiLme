package main
import "fmt"
func main() {
    var n1, m1, n2, m2, ct int
    fmt.Scan(&n1, &m1)
    var vetn [] int = make([]int, n1) 
    var vetm [] int = make([]int, m1)

    for i:=0; i<n1; i++ {
        fmt.Scan(&n2)
        vetn[i]=n2
    }
    for i:=0; i<m1; i++ {
        fmt.Scan(&m2)
        vetm[i]=m2
    }
    for i := range vetm {
        if vetn[i]==vetm[i] {
            ct+=1
        }
    }
        if ct>0 {
            fmt.Println("sim")
        } else {
            fmt.Println("nao")
        }
}
