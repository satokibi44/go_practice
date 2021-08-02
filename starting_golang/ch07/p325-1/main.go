package main

import (
    "log"
    "os"

    "io/ioutil"
)

func main() {
    f, err := os.Open("foo.txt")
    if err != nil {
        log.Fatal(err)
    }

    /* foo.txtの入力をすべて読み込む */
    bs, err := ioutil.ReadAll(f)
    if err != nil {
        log.Fatal(err)
    }
    /* bsは[]byte型でファイルのすべてのバイト列が入る */
}
