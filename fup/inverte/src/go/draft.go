package main

import (
	"fmt"
	"unicode"
)
func main() {
    var l string

    fmt.Scanf("%s", &l)
    ll:=rune(l[0])

    if unicode.IsUpper(ll){
        minuscula:=unicode.ToLower(ll)
        fmt.Printf("%c\n", minuscula)
    } 
    if unicode.IsLower(ll) {
        maiuscula:=unicode.ToUpper(ll)
        fmt.Printf("%c\n", maiuscula)
    }
    if !unicode.IsLetter(ll){
        fmt.Println(ll)
    }
}