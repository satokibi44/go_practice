package foo

type I interface {
    Method1() string
    method2() string
}

type T struct{}

func (t *T) Method1() string {
    return "Method1()"
}

func (t *T) method2() string {
    return "method2()"
}

func NewI() I {
    return &T{}
}
