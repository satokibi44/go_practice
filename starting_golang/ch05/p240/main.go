package main

t := foo.NewI()
t.Method1() // OK
t.method2() // コンパイルエラー
