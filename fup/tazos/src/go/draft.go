package main
import "fmt"
func main() {
    var qtdtaz, tazos int
    fmt.Scan(&qtdtaz)
    var taz[]int= make([]int, qtdtaz)

    for i:=1; i<qtdtaz; i++ {
        fmt.Scan(&tazos)
        taz[tazos]++
    }
    
    var rep[] int

    for i:=0; i<qtdtaz; i++{
        if taz[i]>1{
            rep=append(rep, i)
        }
    }

    fmt.Println(rep)

}