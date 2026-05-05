package main
import "fmt"
func main() {
    var m, f1, f2, f3 int64

    fmt.Scan(&m, &f1, &f2)

    f3=m-(f1+f2)
    
    if f1>f2 && f1>f3 {
        fmt.Println(f1)
    }
}
