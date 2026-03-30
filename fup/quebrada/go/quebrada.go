package main
import "fmt"

func main() {
    var a, b, d, rd int64
    var rf float64

    fmt.Scan(&a, &b)
	
    d=a/b
    rd=a%b
    rf=float64(a)/float64(b)

    fmt.Printf("%d\n%d\n%.2f", d, rd, rf)
}

