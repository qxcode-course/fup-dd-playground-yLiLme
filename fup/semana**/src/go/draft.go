package main
import "fmt"
func main() {
    var d, h int64
        fmt.Scan(&d, &h)

    if d>=2 && d<=6 {
        if h>=8 && h<=11 || h>=14 && h<=17{
            fmt.Println("SIM")
        } else {
            fmt.Println("NAO")
        }
    } else if d==7 { 
            if h>=8 && h<=11 {
                fmt.Println("SIM")
            } else {
                fmt.Println("NAO")
            }
    } else {
        fmt.Println("NAO")
    }
}
