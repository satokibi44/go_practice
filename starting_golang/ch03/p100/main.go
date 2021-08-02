package main

import (
    "fmt"
)

func doSomething() (a int) {
    return
}

func main() {
    fmt.Println(doSomething()) // => "0"
}
