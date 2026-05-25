package main
import "fmt"
func main() {
    var n, nums int
    fmt.Scan(&n)
    var vet[]int=make([]int, n)

    for i:=0; i<n; i++ {
        fmt.Scan(&nums)
        vet[i]=nums
    } 
        fmt.Printf("[ ")
    for i:=n-1; i>=0; i-- {
        fmt.Printf("%d ", vet[i])
    }
        fmt.Printf("]\n")
}