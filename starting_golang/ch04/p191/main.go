package main

import (
    "fmt"
    "time" // ウエイトのために使用
)

func receive(name string, ch <-chan int) {
    for {
        i, ok := <-ch
        if ok == false {
            /* 受信できなくなったら終了 */
            break
        }
        fmt.Println(name, i)
    }
    fmt.Println(name + " is done.")
}

func main() {
    ch := make(chan int, 20)

    go receive("1st goroutine", ch)
    go receive("2nd goroutine", ch)
    go receive("3rd goroutine", ch)

    i := 0
    for i < 100 {
        ch <- i
        i++
    }
    close(ch)

    /* ゴルーチンの完了を3秒待つ */
    time.Sleep(3 * time.Second)
}
