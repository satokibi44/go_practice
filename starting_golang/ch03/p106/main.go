package main

import (
    "fmt"
)

func later() func(string) string {
    // 1つ前に与えられた文字列を保存するための変数
    var store string
    // 引数に文字列をとり文字列を返す関数を返す
    return func(next string) string {
        s := store
        store = next
        return s
    }
}

func main() {
    f := later()

    fmt.Println(f("Golang"))   // => ""
    fmt.Println(f("is"))       // => "Golang"
    fmt.Println(f("awesome!")) // => "is"
}
