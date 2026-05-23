package main
import "fmt"
func main() {
    var n, p, s, r int
    fmt.Scan(&n)
    var vet[]int=make([]int, n)

    for i:=0; i<n; i++{
        fmt.Scan(&p)
        vet[i]=p
    }
    for i:= range vet{
        if vet[i]%2==0 || vet[i]==0{
            r+=vet[i]
        } else {
            s+=vet[i]
        }
    }
        if r>s{
            fmt.Println("rebeldes")
        } else if s>r{
            fmt.Println("soldados")
        } else {
            fmt.Println("empate")
        }
}
