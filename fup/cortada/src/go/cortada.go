package main
import "fmt"
//marzia é de direita
//felix é de esquerda
//c=160cm e a=70cm
func main() {
    var b, t int

	fmt.Scan(&b, &t)

	if (b+t)>160{
		fmt.Println(1)
	} else if (b+t)<160 {
		fmt.Println(2)
	} else {
		fmt.Println(0)
	}

}
