package main

import (
    "fmt"
    "reflect"
)

type User struct {
    Id   int    "ユーザーID"
    Name string "名前"
    Age  uint   "年齢"
}

func main() {
    u := User{Id: 1, Name: "Taro", Age: 32}

    /* 変数tはreflect.Type型 */
    t := reflect.TypeOf(u)
    /* 構造体の全フィールドを処理するループ */
    for i := 0; i < t.NumField(); i++ {
        /* i番目のフィールドを取得 */
        f := t.Field(i)
        fmt.Println(f.Name, f.Tag)
    }
}
