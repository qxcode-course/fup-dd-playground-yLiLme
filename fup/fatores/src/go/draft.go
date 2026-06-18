package main
import "fmt"
func main() {
    var n int
    div:=2

    fmt.Scan(&n)

    for n>1 {
        c:=0
        for n%div==0 {
            c++
            n=n/div
        }
        if c>0{
            fmt.Printf("%d %d\n", div, c)
        }
        div=div+1
    }

}
