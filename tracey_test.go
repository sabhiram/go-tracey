package tracey

import (
    //"fmt"
    "os"
    "log"
    "testing"

    //"github.com/stretchr/testify/assert"
)

var Warn = log.New(os.Stdout, "Custom:", 0)
var G, O = GetTraceFunctions(Options{SpacesPerIndent: 4, CustomLogger: Warn, EnterMessage: "e_n_t_e_r: "})

func TestGetTraceFunctions(test *testing.T) {
    defer G(O())


    // Outputs:
    // [ 0]ENTER: Example_GetTraceFunctions
    // [ 0]EXIT:  Example_GetTraceFunctions

    func(s string) {
        defer G(O("FIRST FUNC"))

        func(s string) {
            defer G(O())
        }("Another str")

    }("Test string")

    var v = func(s string) {
        defer G(O("FN_V"))
    }
    v("Yeehaw")
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
