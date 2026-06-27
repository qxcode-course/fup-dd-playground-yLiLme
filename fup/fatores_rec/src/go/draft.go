package main
import "fmt"

func calcFatores (n int, div int, resp map[int]int) int {
    if n==1{
        return 1
    }
    if n%div==0{
        resp[div]+=1
        return calcFatores(n, div, resp)
    }
    if n%div!=0{
        div++
        return calcFatores(n, div, resp)
    }

        return calcFatores(n, div, resp)
}

func main() {
    var num int
    div:=2
    
    cache:=make(map[int]int)

    fmt.Scan(&num)
    calcFatores(num, div, cache)

    for i:= range cache {
        fmt.Printf("%d\n", cache[i])
    }

}