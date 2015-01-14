package tracey

import (
    "testing"
)

var G, O = GetTraceFunctions()

func TestGetTraceFunctions(test *testing.T) {
    defer G(O())

    Foobar(10)
}

func Foobar(i int) {
    defer G(O())

    if i > 0 {
        BarFoo(i)
    }
}

func BarFoo(i int) {
    defer G(O())

    Foobar(i - 1)
}
