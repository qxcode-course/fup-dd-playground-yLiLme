package main
import "fmt"

func main() {
    var t, h, m, s, rs, rt int64

    fmt.Scan(&t)

    h=t/3600
    rt=t%3600 //o que sobra pra m e s
    m=rt/60
    rs=rt%60
    s=rs

    fmt.Printf("%d:%d:%d\n", h, m, s)

}