// fooパッケージの概要説明。
package foo

const (
    // ○○の最大値
    MAX = 100
)

const (
    A = iota
    B
    C
)

// T型の定義
type T struct {
    Filed1 int
    field2 string
}

// *T型に定義された1番目のメソッド
func (t *T) Method1() {
}

// *T型に定義された2番目のメソッド
func (t *T) Method2() {
}

// 型コンストラクタ
func New() *T {
    return &T{}
}
