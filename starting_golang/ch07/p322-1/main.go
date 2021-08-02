package main

import (
    "bufio"
    "fmt"
    "strings"
)

func main() {
    s := `XXXXX
YYYYY
ZZZZZ`
    /* 文字列からReaderを生成 */
    r := strings.NewReader(s)

    scanner := bufio.NewScanner(r)
    scanner.Scan()
    fmt.Println(scanner.Text()) // => "XXXXX"
    scanner.Scan()
    fmt.Println(scanner.Text()) // => "YYYYY"
    scanner.Scan()
    fmt.Println(scanner.Text()) // => "ZZZZZ"
}
