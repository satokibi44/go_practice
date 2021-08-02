package main

import (
    "log"

    "io/ioutil"
)

func main() {
    /* foo.txtファイルの内容をすべて読み込む */
    bs, err := ioutil.ReadFile("foo.txt")
    if err != nil {
        log.Fatal(err)
    }
    /* bsは[]byte型でファイルのすべてのバイト列が入る */
}
