package tracey

import (
    "testing"

    //"github.com/stretchr/testify/assert"
)

var G, O = GetTraceFunctions(Options{SpacesPerIndent: 4})

func ExampleGetTraceFunctions(test *testing.T) {
    defer G(O())
    // Outputs:
    // [ 0]ENTER: Example_GetTraceFunctions
    // [ 0]EXIT:  Example_GetTraceFunctions
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
