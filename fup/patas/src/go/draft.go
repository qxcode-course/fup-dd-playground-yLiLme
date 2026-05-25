package main
import "fmt"
func main() {
    var cb, c, a, qtani int
    var ani string

    //chico bento chuta primeiro
    fmt.Scan(&cb, &c, &a)
    var vet[] string = make([]string, a)

    for i:=0; i<a; i++ {
        fmt.Scan(&ani)
        vet[i]=ani
    }
    for i := range vet {
        if vet[i]=="v" || vet[i]=="c"{
            qtani+=4
        } else if vet[i]=="g" {
            qtani+=2
        }
    }
        fmt.Printf("%d\n", qtani)
    
    cb=cb-qtani
    c=c-qtani

    if cb<0 {
        cb=cb*(-1)
    } else if c<0 {
        c=c*(-1)
    }
    
    if cb<c {
        fmt.Println("Chico Bento")
    } else if c<cb{
        fmt.Println("Cebolinha")
    } else {
        fmt.Println("empate")
    }

}
