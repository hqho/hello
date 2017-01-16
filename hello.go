package main

import (
    "fmt"

    "github.com/hqho/stringutil"
)

func main() {
    fmt.Printf(stringutil.Reverse("!oG ,olleH"))
    fmt.Printf("\n")

/*
    b := [...]string{"Penn", "Teller"}
    fmt.Println("b:", b)
*/

    letters := []string{"a", "b", "c", "d"}
    fmt.Println("letters:", letters)

    var s []byte
    s = make([]byte, 5, 5)
    t := make([]byte, len(s), (cap(s)+1)*2) // +1 in case cap(s) == 0
    copy(t, s)
    s = t
    fmt.Println("s:", s)

    a := []string{"John", "Paul"}
    b := []string{"George", "Ringo", "Pete"}
    a = append(a, b...) // equivalent to "append(a, b[0], b[1], b[2])"
    fmt.Println("a:", a)
}
