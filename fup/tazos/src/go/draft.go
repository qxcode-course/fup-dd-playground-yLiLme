package main
import "fmt"

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
    var qtdtaz int
    fmt.Scan(&qtdtaz)
    var taz[]int= make([]int, qtdtaz)

    for i:=0; i<qtdtaz; i++ {
        fmt.Scan(&taz[i])
    }
    
    f:=make(map[int]int)
    mrep:=0

    for _, v:=range taz {
        f[v]++
        if f[v]>mrep{
            mrep=f[v]
        }
    }
    
    var win []int

    for _, v:=range taz{
        if f[v]==mrep{
            add:=false
            for _, w:=range win {
                if w==v{
                    add=true
                    break
                }
            }
            if !add{
                win=append(win, v)
            }
        }
    }
        formata(win)
}