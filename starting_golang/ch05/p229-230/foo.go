package foo

type T struct {
    Field1 int // 公開フィールド
    field2 int // 非公開フィールド
}

/* 公開メソッド */
func (t *T) Method1() int {
    return t.Field1
}

/* 非公開メソッド */
func (t *T) method2() int {
    return t.field2
}
