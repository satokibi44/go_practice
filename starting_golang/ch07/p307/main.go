package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    /* 標準入力をソースにしたスキャナの生成 */
    scanner := bufio.NewScanner(os.Stdin)

    /* 入力のスキャンが成功する限り繰り返すループ */
    for scanner.Scan() {
        fmt.Println(scanner.Text()) // スキャン内容を文字列で出力
    }

    /* スキャンにエラーが発生した場合の処理 */
    if err := scanner.Err(); err != nil {
        fmt.Fprintln(os.Stderr, "読み込みエラー:", err)
    }
}
