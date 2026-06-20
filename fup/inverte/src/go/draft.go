package main

import (
	"fmt"
	"unicode"
)
func main() {
    var l rune

    fmt.Scan(&l)

    if unicode.IsUpper(l){
        minuscula:=unicode.ToLower(l)
        fmt.Println(minuscula)
    } 
    if unicode.IsLower(l) {
        maiuscula:=unicode.ToUpper(l)
        fmt.Println(maiuscula)
    }
}