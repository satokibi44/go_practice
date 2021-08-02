package main

t := &foo.T{}
t.Method1()  // OK
t.Field1     // OK
t.method2()  // コンパイルエラー
t.field2     // コンパイルエラー
