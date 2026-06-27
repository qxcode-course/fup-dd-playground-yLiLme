package main
import "fmt"
func main() {
    var nome string
    var age int64

    fmt.Scanf("%s", &nome)
    fmt.Scan(&age)

    if age<12 {
        fmt.Printf("%s eh crianca\n", nome)
    } else if age>=12 && age<18 {
        fmt.Printf("%s eh jovem\n", nome)
    } else if age>=18 && age<65 {
        fmt.Printf("%s eh adulto\n", nome)
    } else if age>=65 && age<1000 {
        fmt.Printf("%s eh idoso\n", nome)
    } else {
        fmt.Printf("%s eh mumia\n", nome)
    }
}
