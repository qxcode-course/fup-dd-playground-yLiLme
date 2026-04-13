package main
import "fmt"
func main() {
    var nome string
    var age int64

    fmt.Scanf("%c", &nome)
    fmt.Scan(&age)

    if nome=="crianca" && age<12 {
        fmt.Printf("%c eh crianca", nome)
    }
}
