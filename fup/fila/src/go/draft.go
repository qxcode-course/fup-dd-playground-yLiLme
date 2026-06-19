package main
import (
    "fmt"
    "strings"
)

func TranformaSlice (slice[]int) string {
    sli:=make([]string, len(slice))
    var resultado string
    for i, v:= range slice {
        sli[i]=fmt.Sprintf("%d", v)
    }
    if len(slice)==0 {
        resultado="[" + " ]"
    } else {
        resultado= "[ "+strings.Join(sli, " ")+" ]"
    }   
    return resultado
}

func main() {
    var p, asp int
    fmt.Scan(&p)
    var pessoas[]int=make([]int, p)

    for i:=0; i<p; i++{
        fmt.Scan(&asp)
        pessoas[i]=asp
    }
        var sta[] int
        var dis[] int
    
    for i:=0; i<p; i++{
        if pessoas[i]%2==0{
            sta=append(sta, pessoas[i])
        } else {
            dis = append(dis, pessoas[i])
        }
    } 

        fmt.Println(TranformaSlice(dis))
        fmt.Println(TranformaSlice(sta))


}
