package main
import "fmt"
func main() {
    var a, b, cp, cb int64

    fmt.Scan(&a, &b)

    cp=a
    cb=b

    fmt.Printf("[ ")
    for i:=a; i<=b; i++ {
        fmt.Printf("%d %d ", cp, cb)
        cp+=1
        cb-=1
    } 
    fmt.Printf("]\n")

}
