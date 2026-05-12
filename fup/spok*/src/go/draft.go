package main
import "fmt"

func verificaPalindromo (num int64) int64 {
    if num%10==1 {
        return 1
    } else if num==0 {
        return 1
    } else {
        return 0
    }
}

func main() {
    var id int64

    fmt.Scan(&id)

    fmt.Println(verificaPalindromo(id))
}
