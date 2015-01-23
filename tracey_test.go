package tracey

import (
    "fmt"
    "os"
    "log"
    "testing"

    //"github.com/stretchr/testify/assert"
)

var Warn = log.New(os.Stdout, "Custom:", 0)
var G, O = GetTraceFunctions(Options{SpacesPerIndent: 4, CustomLogger: Warn, EnterMessage: "e_n_t_e_r: "})

func Foo(a ...interface{}) string {
    fmtStr, ok := a[0].(string)
    if ok {
        return fmt.Sprintf(fmtStr, a[1:]...)
    }
    return "ERROR"
}

func TestFoo(test *testing.T) {
    fmt.Printf("%s\n", Foo("Test--"))
}

func TestGetTraceFunctions(test *testing.T) {
    defer G(O())

    // Outputs:
    // [ 0]ENTER: Example_GetTraceFunctions
    // [ 0]EXIT:  Example_GetTraceFunctions

    func(s string) {
        defer G(O(" --> $FN <-- "))

        func(s string) {
            defer G(O())
        }("Another str")

    }("Test string")

    var v = func(s string) {
        defer G(O("FN_V %s <-- is my value", s))
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
