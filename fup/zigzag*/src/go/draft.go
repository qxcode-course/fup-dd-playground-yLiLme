package main
import "fmt"
func main() {
    var is, fs int64

    fmt.Scan(&is, &fs)

    for i:=is; i<=fs; i++ {
        if i%3==0 && i%5==0{
            fmt.Println("zigzag")
        } else if i%5==0 {
            fmt.Println("zag")
        } else if i%3==0 {
            fmt.Println("zig")
        } else {
            fmt.Println(i)
        }
    }
}