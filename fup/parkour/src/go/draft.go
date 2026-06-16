package main
import "fmt"
func main() {
    var n, nums int
    fmt.Scan(&n)
    var vet []int = make([]int, n)

    for i:=0; i<n; i++ {
        fmt.Scan(&nums)
        vet[i]=nums
    }

    parkas := 0
    for i:=0; i<n-1; i++ {
        diferença:=vet[i]-vet[i+1]
        if diferença<0{
            diferença=diferença*(-1)
            if diferença>1 {
                parkas++
            }
        } else if diferença>1 {
            parkas++
        }
    }
        fmt.Println(parkas)
}