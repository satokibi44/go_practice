package main

import (
    "fmt"
    "sync"
    "time"
)

var st struct{ A, B, C int }

/* ミューテックスを保持するパッケージ変数 */
var mutex *sync.Mutex

func UpdateAndPrint(n int) {
    /* ロック */
    mutex.Lock()

    st.A = n
    time.Sleep(time.Microsecond)
    st.B = n
    time.Sleep(time.Microsecond)
    st.C = n
    time.Sleep(time.Microsecond)
    fmt.Println(st.A, st.B, st.C)

    /* アンロック */
    mutex.Unlock()
}

func main() {
    /* ミューテックスの生成 */
    mutex = new(sync.Mutex)

    for i := 0; i < 5; i++ {
        go func() {
            for i := 0; i < 1000; i++ {
                UpdateAndPrint(i)
            }
        }()
    }
    for {
    }
}
