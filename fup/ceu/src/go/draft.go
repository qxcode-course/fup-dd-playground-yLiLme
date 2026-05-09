package main
import "fmt"
func main() {
    var p int64

    fmt.Scan(&p)

        fmt.Print("[ ")
    for i:=0; i<10; i++ {
        if p==int64(i){
            continue
        }
        fmt.Printf("%d ", i)
    }
        fmt.Printf("ceu ]\n")
}
 