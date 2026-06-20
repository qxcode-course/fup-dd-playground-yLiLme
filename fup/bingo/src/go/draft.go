package main
import "fmt"
func main() {
    var n1, n2, nums int
     fmt.Scan(&n1, &n2)
    var matriz[][]int=make([][]int, n1, n2)

    for i:=0; i<n1; i++{
        for j:=0; j<n2; j++{
            fmt.Scan(&nums)
        }
    }

    fmt.Println(matriz)
}