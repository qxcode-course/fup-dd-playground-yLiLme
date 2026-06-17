package main
import "fmt"
func main() {
    var nums int 
    var vet[5] int  

    for i:=0; i<5; i++ {
        fmt.Scan(&nums)
        vet[i]=nums
    }
    d:=vet[0]
    for i:=0; i<5; i++{
        if vet[i]<d{
            d=vet[i]
        }
        
    }
        fmt.Println(d)
}
