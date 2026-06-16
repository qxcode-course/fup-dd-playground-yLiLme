package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)

    e := make([]int, 70)
    d := make([]int, 70)

    for i:=0; i<n; i++ {
        var tam int
        var l string

        fmt.Scan(&tam, &l)

        if l=="E"{
            e[tam]++
        } else {
            d[tam]++
        }
    }

    p:=0

    for i:=0; i<70; i++ {
        if e[i]<d[i] {
            p+= e[i]
        } else {
            p+= d[i]
        }
    }
    
        fmt.Println(p)
}
