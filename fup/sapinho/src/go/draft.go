package main
import "fmt"
func main() {
    var p, s, e, altu int

    fmt.Scan(&p, &s, &e)

    for i:=altu; i<p; i++ {
        zero:=altu
        altu=altu+s
        if altu>=p{
            fmt.Printf("%d saiu\n",zero)
            break
        } else {
            fmt.Printf("%d %d\n", zero, altu)
        }
        altu=altu-e
    } 
}
