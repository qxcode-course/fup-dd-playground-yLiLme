package main
import "fmt"
func main() {
    var c, a int

    fmt.Scan(&c, &a)
        capM:=c-1

    if a%capM==0 {
        fmt.Println(a/capM)
    } else if a%capM!=0 {
        fmt.Println((a/capM)+1)
    }

}