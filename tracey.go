// `Tracey` is a simple library which allows for much easier function enter / exit logging
package tracey

import (
    "fmt"
    "regexp"
    "strings"

    "runtime"
)

var SPACES_PER_TAB = 2
var DEPTH = 0

func getDepth() string {
    return fmt.Sprintf("[%2d]%s", DEPTH, strings.Repeat(" ", DEPTH * SPACES_PER_TAB))
}
func _increment() {
    fmt.Printf("Pre increment %d\n", DEPTH)
    DEPTH += 1
}
func _decrement() {
    fmt.Printf("Pre descrement %d\n", DEPTH)
    DEPTH -= 1
    if DEPTH < 0 {
        panic("Depth is negative! Should never happen!")
    }
}

func _enter() string {
    defer _increment()
    programCounter := make([]uintptr, 10)
    runtime.Callers(2, programCounter)
    functionObject := runtime.FuncForPC(programCounter[0])

    stripFilePath := regexp.MustCompile(`^.*\.(.*)$`)
    fnName := stripFilePath.ReplaceAllString(functionObject.Name(), "$1")

    fmt.Printf("%sENTER: %s\n", getDepth(), fnName)
    return fnName
}

func _exit(s string) {
    _decrement()
    fmt.Printf("%sEXIT:  %s\n", getDepth(), s)
}

func GetTraceFunctions() (func(string), func() string) {
    return _exit, _enter
}
