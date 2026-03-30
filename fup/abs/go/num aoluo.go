package main
import "fmt"

func main() {
    var num1, num2, num11, num22, vabs int64
    
    fmt.Scan(&num1, &num2)

        if num1<0{
            num11=num1*(-1)
            vabs=num11-num2
            println(vabs)
        } else if num2<0 {
                num22=num2*(-1)
                vabs=num1-num22
                println(vabs)
        } else {
            num11=num1*(-1)
            num22=num2*(-1)
            vabs=num11-num22
            println(vabs)
        }
    
}
