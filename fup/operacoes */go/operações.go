package main
import "fmt"

func main() {
  var a, b, s, sub, m, rd int64
  var d float64
	fmt.Scan(&a, &b)

	s=a+b
	sub=a-b
	m=a*b
	d=float64(a)/float64(b)
	rd=a%b

	fmt.Printf("%d\n%d\n%d\n%.2f\n%d\n", s, sub, m, d, rd)

}
