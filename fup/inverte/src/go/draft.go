package main

import (
	"fmt"
	"strings"
)

func inver (letra rune) {
    var ver string
    if letra >= 'a' && letra <= 'z' {
        ver = strings.ToUpper(string(letra))
        fmt.Println(ver)
    } else if letra >= 'A' && letra <= 'Z' {
        ver = strings.ToLower(string(letra))
        fmt.Println(ver)
    } else {
        fmt.Printf("%c\n", letra)
    }
}

func main() {
    var l rune

    fmt.Scanf("%c", &l)
    inver(l)
}