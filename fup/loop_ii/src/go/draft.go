package main
import "fmt"
func main() {
    var a, b int64

    fmt.Scan(&a, &b)

      fmt.Print("[")
    for i:=a; i<b; i++ {        
        fmt.Printf(" %d", i)
    }
      fmt.Print(" ]\n")
}
