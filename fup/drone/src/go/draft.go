package main

import "fmt"

func main() {
	var a, b, c, h, l int
	fmt.Scan(&a)
	fmt.Scan(&b)
	fmt.Scan(&c)
	fmt.Scan(&h)
	fmt.Scan(&l)

	if (a <= h && b <= l) || (a <= l && b <= h) ||
		(a <= h && c <= l) || (a <= l && c <= h) ||
		(b <= h && c <= l) || (b <= l && c <= h) {
		fmt.Println("S")
	} else {
		fmt.Println("N")
	}
}