package main

import (
    "bufio"
    "fmt"
    "strings"
)

func main() {
    s := `ABC DEF
GHI JKL MNO
PQR STU VWX
YZ
`
    r := strings.NewReader(s)
    scanner := bufio.NewScanner(r)

    /* スキャン関数をbufio.ScanWordsに変更 */
    scanner.Split(bufio.ScanWords)

    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
}
