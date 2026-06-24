package main
import (
    "fmt"
)

func formata (vet []int) {
    if len(vet)==0{
        fmt.Println("[ ]")
    } else {

    fmt.Printf("[ ")
    for _, v:=range vet {
        fmt.Printf("%d ", v)
    }
    fmt.Print("]\n")
    }
}

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

    for i:=1; i<=talbu; i++{
        if albu[i]>1{
            VezesRep:=albu[i]-1
                for j:=0; j<VezesRep; j++ {
                    rep=append(rep, i)
                }   
        }
    }
    
    for i:=1; i<=talbu; i++ {
        if albu[i]==0{
            f=append(f, i)
        }
    }
    
    formata(rep)
    formata(f)
}
