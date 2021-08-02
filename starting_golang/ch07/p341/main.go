package main

import (
    "fmt"
    "log"

    "net/http"
    "net/url"
)

func main() {
    /* POSTに送信するデータを生成 */
    vs := url.Values{}
    vs.Add("id", "1")
    vs.Add("message", "メッセージ")
    fmt.Println(vs.Encode()) // => "id=1&message=%E3%83%A1%E3%83%83%E3%82%BB%E3%83@<dtp>{lb}%BC%E3%82%B8"

    /* POSTメソッドを実行 */
    res, err := http.PostForm("https://example.com/comments/post", vs)
    if err != nil {
        log.Fatal(err)
    }
}
