package main
import "fmt"
/*
11, 12 e 13 == 10 pontos
1 == 11 || o necessário pra fazer 21 (que seja menos de 11 e se a soma de todas as cartas der
mais de 21 o as vai valer 1)
o resto vale o resto
*/
func main() {
	var n, nums int
	fmt.Scan(&n)
	var cartas[]int=make([]int, n)

	for i := range cartas {
		fmt.Scan(&nums)
		cartas[i]=nums
	}
			
	c:=0
	qtda:=0
	
	for i:=range cartas{
		if cartas[i]==1{
			c+=11
			qtda++
		} else if cartas[i]==11 || cartas[i]==12 || cartas[i]==13 {
			c+=10
		} else {
			c+=cartas[i]
		}
	}
	for c>21 && qtda>0{
		c-=10
		qtda--
	}

	fmt.Println(c)

}
