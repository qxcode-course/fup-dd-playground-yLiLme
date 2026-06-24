package main
import "fmt"

func ehPrimo (n, d int) int{
    if d==1 {
        return 1
    } else if n%d==0 {
        return  0
    } else if n<=1{
        return 0
    } 

    return ehPrimo(n, d-1)

}

func main() {
    var num int

    fmt.Scan(&num)
    fmt.Println(ehPrimo(num, num-1))


}