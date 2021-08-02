package main

import (
    "fmt"
    "io"

    "net/http"
)

func infoHandler(w http.ResponseWriter, req *http.Request) {
    /* HTMLテキストをhttp.ResponseWriterへ書き込み */
    io.WriteString(w, `
<!DOCTYPE html>
<html lang="ja">
<head>
<meta charset="UTF-8">
<title>インフォメーション</title>
</head>
<body>
<h1>ようこそ！</h1>
</body>
</html>
`)
}

func main() {
    /* URLのパス"/info"を処理する関数を登録 */
    http.HandleFunc("/info", infoHandler)
    /* localhost:8080でサーバー処理開始 */
    http.ListenAndServe(":8080", nil)
}
