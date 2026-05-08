package main
import "fmt"
func main() {
    var a, b int64

    fmt.Scan(&a, &b)

    if a>b {
          fmt.Printf("[ ")
        for i:=a; i>b; i-- {
            fmt.Printf("%d ", i)
        }
          fmt.Printf("]\n")
    } else if a<b {
          fmt.Printf("[ ")
        for i:=a; i<b; i++ {
            fmt.Printf("%d ", i)
        }
          fmt.Printf("]\n")
    }

}
