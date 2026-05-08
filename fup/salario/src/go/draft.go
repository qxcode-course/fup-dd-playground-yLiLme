package main
import "fmt"
func main() {
    var sl, nsl float64

    fmt.Scan(&sl)

    if sl<=1000 {
        nsl=(sl*20/(100))+sl
        fmt.Printf("%.2f\n",nsl)
    } else if sl>1000 && sl<=1500 {
        nsl=(sl*15/(100))+sl
        fmt.Printf("%.2f\n",nsl)
    } else if sl>1500 && sl<=2000 {
        nsl=(sl*10/(100))+sl
        fmt.Printf("%.2f\n",nsl)
    } else {
        nsl=(sl*5/(100))+sl
        fmt.Printf("%.2f\n",nsl)
    }
}
