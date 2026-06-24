package main
import "fmt"
func main() {
    var n, c int
        fmt.Scan(&n)
    var animais[]int=make([]int, n)

    for i:=0; i<n; i++{
        fmt.Scan(animais[i])
    }
    for i:=range animais {
        if animais[i]+animais[i+1]==0{
            c++
        }
    }
        fmt.Println(c)
}
