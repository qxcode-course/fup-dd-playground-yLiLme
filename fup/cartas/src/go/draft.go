package main
import (
    "fmt"
    "strings"
)
func main() {
    var n, nums int
    fmt.Scan(&n)
    var cartas[]string = make([]string, n)

    for i := 0; i < n; i++ {
        fmt.Scan(&nums)
        if nums==1{
            cartas[i]="A"
        } else if nums==11{
            cartas[i]="J"
        } else if nums==12{
            cartas[i]="Q"
        } else if nums==13{
            cartas[i]="K"
        } else {
            numtext:=fmt.Sprintf("%d", nums)
            cartas[i]=numtext
        }
    }

    
    cards:="["+strings.Join(cartas, ", ")+"]"
    fmt.Println(cards)
    
}