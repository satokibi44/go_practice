package main

import "foo"

foo.MAX             // fooパッケージの定数MAXを参照
foo.internal_const  // コンパイルエラー

foo.FooFunc(5)      // fooパッケージの関数FooFuncを参照
foo.internalFunc(5) // コンパイルエラー
