package main
import "fmt"

func PrintaLimite ( frase []string, inicio, lim int ) []string {
    var newFrase []string=make([]string, lim)
    
    for i:=inicio; i<lim; i++ {
        newFrase[i]=frase[i]
    }

    return newFrase

}

func main() {
    var frahsi[]string=nil
    var ini, limi int

        fmt.Scan(&ini, &limi)
    for {
        fmt.Scanf("%s!=\n", &frahsi)
    }
    
    fmt.Println(PrintaLimite(frahsi, ini, limi))


}