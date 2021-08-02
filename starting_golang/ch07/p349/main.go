package main

import (
    "fmt"
    "io"

    "crypto/md5"
)

func main() {
    h := md5.New()
    io.WriteString(h, "ABCDE")

    fmt.Println(h.Sum(nil))
    // => "[46 205 222 57 89 5 29 145 63 97 177 69 121 234 19 109]"

    /* MD5ハッシュ値から16進数の文字列を生成 */
    s := fmt.Sprintf("%x", h.Sum(nil))
    fmt.Println(s)
    // => "2ecdde3959051d913f61b14579ea136d"
}
