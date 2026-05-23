package main
import "fmt"
func main() {
    var nvet,  ani, ct int64
    fmt.Scan(&nvet)
    var vet [] int = make([]int, nvet)

    for i:=0; i<nvet; i++ {
        fmt.Scan(&ani)
        vet[i]=ani
    }
    for i:= range vet {
        if {
            ct+=1
        }
    }
    fmt.Println(ct)
}
