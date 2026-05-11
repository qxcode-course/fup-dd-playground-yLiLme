package main
import "fmt"
func main() {
    var n int
    var p string

    fmt.Scan(&n, &p)

    fmt.Printf("[ ")
    for i:=0; i<10; i++ {
        if i==n {
            continue
        }
        if p=="d" {
            fmt.Printf("%d%s ", i, p)    
            p="e"       
        } else {
            fmt.Printf("%d%s ", i, p)
            p="d"
        } 
        if i==9 && n!=10 {
            fmt.Print("ceu ")
        }
       
    }
    fmt.Printf("]\n")
}
