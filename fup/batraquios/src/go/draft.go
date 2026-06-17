package main
import "fmt"
func main() {
    var n1, m1, n2, m2 int
    
    fmt.Scan(&n1)
    var vetn [] int = make([]int, n1) 
    
    for i:=0; i<n1; i++ {
        fmt.Scan(&n2)
        vetn[i]=n2
        
    }

    fmt.Scan(&m1)
    var vetm [] int = make([]int, m1)

    for i:=0; i<m1; i++ {
        fmt.Scan(&m2)
        vetm[i]=m2
    }
        var aall=true
    for i:=0; i<n1; i++ {
        var anum=false
        for j:=0; j<m1; j++{
            if vetn[i]==vetm[j] {
                anum=true
                break
            }
        }
        if anum==false {
            aall=false
            break
        }
    }
        if aall==true{
            fmt.Println("sim")
        } else {
            fmt.Println("nao")
        }
}
