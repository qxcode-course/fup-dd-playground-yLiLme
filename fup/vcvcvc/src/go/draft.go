package main
import "fmt"
func main() {
    var char byte

    for {
        n, err := fmt.Scanf("%c", &char)

        if err != nil || n != 1 || char == '\n' {
            break
        }

        if char == '\r' {
            continue
        }

        if char == ' ' {
            fmt.Print(" ")
        } else {
            letra := char
            if letra >= 'A' && letra <= 'Z' {
                letra = letra + 32 
            }

            if letra == 'a' || letra == 'e' || letra == 'i' || letra == 'o' || letra == 'u' {
                fmt.Print("v")
            } else if letra >= 'a' && letra <= 'z' {
                fmt.Print("c")
            }
        }
    }
        fmt.Println()
}