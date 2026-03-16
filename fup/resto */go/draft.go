package main
import "fmt"

func main(){
    var a, b, d, resto int32
    fmt.Scan(&a, &b)
    d=a/b
    resto=a%b
    fmt.Println(d, resto)

}