package main
import "fmt"

func ehVogal (c byte) bool {
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'A' || 
    c == 'E' || c == 'I' || c == 'O' || c == 'U' 

}

func main() {
    var p, resul string

    if n, _ := fmt.Scan(&p); n == 1 {
        resul=p
    }

    for {
        n, _ := fmt.Scan(&p)
        if n!=1 {
            break
        }
        tamRes := len(resul)

        if tamRes > 0 && ehVogal(resul[tamRes-1]) && ehVogal(p[0]) {
            resul= resul[:tamRes-1]
            resul+=p
        } else {
            resul += " " + p
        }    
    }

        fmt.Println(resul)

}