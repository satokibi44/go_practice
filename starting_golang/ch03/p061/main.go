package main

import (
    "fmt"
)

/* nはパッケージ変数 */
var n = 100

func main() {
    /* パッケージ変数nの値を+1して表示 */
    n = n + 1
    fmt.Printf("n=%d\n", n)
}
