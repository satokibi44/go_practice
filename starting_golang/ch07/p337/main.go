package main

import (
    "encoding/json"
    "fmt"
    "log"
    "time"
)

type User struct {
    Id      int
    Name    string
    Email   string
    Created time.Time
}

func main() {
    /* User型に対応したJSON文字列 */
    src := `
{
  "Id":12,
  "Name":"田中花子",
  "Email":"tanaka@example.com",
  "Created":"2015-12-02T10:00:00.000000000+09:00"
}
`
    /* User型の初期化 */
    u := new(User)

    /* JSONデコード */
    err := json.Unmarshal([]byte(src), u)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("%+v\n", u)
}
