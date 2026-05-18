package main
import "fmt"
func main() {
    var n int
    fmt.Scan(&n)
    var vet[]int = make([]int, n)

    if n==0 {
        fmt.Printf("[ ]\n")
    } else {
    var nums int
    for i:=0; i<n; i++ {
        fmt.Scan(&nums)
        vet[i]=nums
    }
        fmt.Printf("[ ")
        for i := range vet {
        fmt.Printf("%d ", vet[i])
    }
        fmt.Printf("]\n")
}
}
