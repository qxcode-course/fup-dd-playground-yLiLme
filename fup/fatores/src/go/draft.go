package main
import "fmt"
func main() {
   var n, c int
   f:=2

   fmt.Scan(&n)

    for n>1 || n>0 {
        if n%f==0 {
            n=n/f
            c++
            f++
        } else {
            f++
            c=0
            fmt.Printf("%d %d\n", f, c)
        }
        fmt.Printf("%d %d\n", f, c)
    }
  
}
