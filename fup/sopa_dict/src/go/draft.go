    package main
    import "fmt"

    func fibonacci (n int, cache map[int] int) int {
            valor , existe := cache[n]
        if existe { 
            return valor
        }
        if n == 0 {
            return 0
        }
        if n == 1 {
            return 1
        }
        resultado:=fibonacci(n-1, cache)+fibonacci(n-2, cache)
        cache[n]=resultado

        return resultado

    }

    func main() {
        var n int

        cache:=make(map[int]int)

        fmt.Scan(&n)
        fmt.Println(fibonacci(n, cache))

    }