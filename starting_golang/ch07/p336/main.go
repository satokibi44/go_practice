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
    /* User構造体の初期化 */
    u := new(User)
    u.Id = 1
    u.Name = "山田太郎"
    u.Email = "yamada@example.com"
    u.Created = time.Now()

    /* JSONエンコード */
    bs, err := json.Marshal(u)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(string(bs))
}
