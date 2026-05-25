package main
import "fmt"

type Aluno struct{
    nome string
    n1, n2, n3 float64
}
func média (n1, n2, n3 float64) float64{
    var m float64
    m=(n1+n2+n3)/3
    return m
}

func main() {
    var n int
    fmt.Scan(&n)
    var alu[]int=make([]int, n)

    
}
