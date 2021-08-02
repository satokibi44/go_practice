package main

import (
    "fmt"
)

const ONE = 1 // パッケージに定義された定数

func one() (int, int) {
    const TWO = 2 // 関数内に定義された定数
    return ONE, TWO
}

func main() {
    x, y := one()
    fmt.Printf("x=%d, y=%d\n", x, y) // => "x=1, y=2"
}
