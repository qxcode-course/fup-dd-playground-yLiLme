package main
import "fmt"
func main() {
	var vet[]int
	var nums int

	for {
		_, err:=fmt.Scan(&nums)
		if err!=nil{
			break
		}
		vet=append(vet, nums)
	}
	fmt.Printf("[ ")
	for i:=len(vet)-1; i>=0; i-- {
		fmt.Printf("%d ",vet[i])
	}
	fmt.Printf("]\n")

}
