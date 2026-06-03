package main
import "fmt"
func main() {
    var d, sd, td int
    var s, t string
    fmt.Scan(&d)
    var vet[]string=make([]string, d)
    var vet1[]string=make([]string, d)

    for i:= range vet {
        fmt.Scanf("%s%s", &s, &t)
        vet[i]=s
        vet1[i]=t
    }
    for i:= range vet {
        if vet[i]=="c" {
            sd+=sd+1
        }
        
    }
}
