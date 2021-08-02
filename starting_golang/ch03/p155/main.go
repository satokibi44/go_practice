package main

import (
    "fmt"
)

var S = ""

func init() {
    S = S + "A"
}

func init() {
    S = S + "B"
}

func init() {
    S = S + "C"
}

func main() {
    /* init関数は定義された順序で実行される */
    fmt.Println(S) // => "ABC"
}
