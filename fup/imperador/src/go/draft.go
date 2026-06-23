package main
import "fmt"
func main() {
    var n int
    var combatentes rune
    fmt.Scan(&n)
    var arena[][]int=make([][]int, n)
    g:=2
    c:=1

    for i:=0; i<n; i++ {
        for j:=0; j<n; j++{
            fmt.Scanf("%c",&combatentes)
            arena[i][j]=int(combatentes)
        }
    }
    for i:=0; i<n; i++ {
        for j:=0; j<n; j++{
            
        }
    }
}