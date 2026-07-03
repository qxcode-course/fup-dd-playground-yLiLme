package main
import "fmt"

func Rotatoria (letra rune, rot int) {
    var normaliza int
    normaliza = (int(letra) - 'a' + rot ) % 26

    if normaliza < 0 {
        normaliza += 26
    }
    fmt.Printf("%c\n", rune(normaliza) + 'a')
}

func main() {
    var l rune
    var rot int

    fmt.Scanf("%c\n%d", &l, &rot)
    Rotatoria(l, rot)
    
}