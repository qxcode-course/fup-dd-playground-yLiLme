package main
import "fmt"

func mmcDeFator (a, b int) int {
	if a==0||b==0{
		return 0
	}

	c:=1
	div:=2

	for a>1 || b>1 {
		deu:=false
		if a%div==0{
			a=a/div
			deu=true
		}
		if b%div==0{
			b=b/div
			deu=true
		}
		if deu {
			c=c*div
		} else {
			div++
		}
		
	}
		return c
}

func main() {
    var n1, n2 int

	fmt.Scan(&n1, &n2)
	fmt.Println(mmcDeFator(n1, n2))
    
}