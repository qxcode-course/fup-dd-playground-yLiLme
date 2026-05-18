package main
import "fmt"
func main() {
    var n int
    fmt.Scan(n)
    var vet[]int = make([]int, n)

    var nums int
    for i:=0; i<=n; i++ {
         fmt.Scan(&nums)
         vet[i]=nums
         fmt.Println(vet[i])
    }
    

}
