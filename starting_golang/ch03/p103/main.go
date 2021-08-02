package main

import (
    "fmt"
)

func main() {
    fmt.Printf("%#v\n", func(x, y int) int { return x + y })
    fmt.Printf("%#v\n", func(x, y int) int { return x + y }(2, 3))
}
