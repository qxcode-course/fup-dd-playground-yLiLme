package main
import "fmt"
func main() {
    var c, b, g, m, count int

    fmt.Scan(&c, &b, &g, &m)

    count=b+g+m

    if count%c==0 {
        fmt.Println(count/c)
    } else if count%c!=0 {
        fmt.Println((count/c)+1)
    } 

}
