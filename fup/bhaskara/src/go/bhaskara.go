package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, del, x1, x2 float64

	fmt.Scan(&a, &b, &c)

	del = math.Pow(b, 2) - 4*(a*c)

	if del > 0 {
		x1 = (-b + math.Sqrt(del)) / (2 * a)
		x2 = (-b - math.Sqrt(del)) / (2 * a)
		if x1 == 0 {
			x1 = 0
		}
		if x2 == 0 {
			x2 = 0
		}
		fmt.Printf("%.2f\n%.2f\n", x1, x2)
	} else if del == 0 {
		x1 = -b / (2 * a)
		if x1 == 0 {
			x1 = 0
		}
		fmt.Printf("%.2f\n", x1)
	} else {
		fmt.Println("nao ha raiz real")
	}
}