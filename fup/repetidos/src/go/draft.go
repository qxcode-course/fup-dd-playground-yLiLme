package main
import "fmt"
func main() {
    var p, n, nums, ct int 

    fmt.Scan(&p, &n)
    var vet [] int = make([]int, n)
    
    for i:=0; i<n; i++ {
        fmt.Scan(&nums)
        vet[i]=nums
    }
    for i := range vet {
        if vet[i]==p {
            ct+=1
        }
    }
        fmt.Println(ct)
}
