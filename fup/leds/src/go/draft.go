package main

import (
	"bufio"
	"fmt"
	"os"
	"text/scanner"
)
func main() {
	var n, tl int
	fmt.Scan(&n)

	scanner:=bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for i:=0; i<n; i++{
		scanner.Scan()
		numled:=scanner.Text()
	}
	for i:=0; i<n; i++{
		if numled[i]==1{
			
		}
	}
	

}