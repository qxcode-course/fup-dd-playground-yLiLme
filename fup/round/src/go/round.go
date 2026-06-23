package main
import "fmt" 
import "math"
func main() {
    var num float64
    var char rune

    fmt.Scanf("%c", &char)
    fmt.Scan(&num)

    if char=='c' {
        fmt.Println(math.Ceil(num))
    } else if char=='f' {
        fmt.Println(math.Floor(num))
    } else if char=='r'{
        fmt.Println(math.Round(num))
    }
}
