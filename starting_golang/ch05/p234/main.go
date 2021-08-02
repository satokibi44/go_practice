package main

import (
    "fmt"
    "encoding/json"
)

type User struct {
    Id   int    `json:"user_id"`
    Name string `json:"user_name"`
    Age  uint   `json:"user_age"`
}

func main() {
    u := User{Id: 1, Name: "Taro", Age: 32}
    bs, _ := json.Marshal(u)
    fmt.Println(string(bs))
}
/* 出力内容 */
{"user_id":1,"user_name":"Taro","user_age":32}
