package main
import (
    "fmt"
)
func main() {
    var talbu, qtdtem, cards int
    fmt.Scan(&talbu)
    var albu[]int=make([]int, talbu+1)
    fmt.Scan(&qtdtem)

    for i:=0; i<qtdtem; i++{
        fmt.Scan(&cards)
        albu[cards]++
    }

    var rep[] int
    var f[] int 

    for i:=1; i<talbu; i++{
        if albu[i]>1{
            rep=append(rep, i)
        }
    }
    
    for i:=1; i<talbu; i++ {
        if albu[i]==0{
            f=append(f, i)
        }
    }
    
    fmt.Println(rep)
    fmt.Println(f)
}
