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

type Options struct {
    SpacesPerIndent int
    EnableNesting   bool
    PrintDepthValue bool
}

func getDepth() string {
    return fmt.Sprintf("[%2d]%s", DEPTH, strings.Repeat(" ", DEPTH * SPACES_PER_TAB))
}

func _increment() {
    DEPTH += 1
}

func _decrement() {
    DEPTH -= 1
    if DEPTH < 0 {
        panic("Depth is negative! Should never happen!")
    }
}

func _enter(ss ...string) string {
    defer _increment()

    fnName := ""
    if len(ss) == 0 {
        programCounter := make([]uintptr, 10)
        runtime.Callers(2, programCounter)
        functionObject := runtime.FuncForPC(programCounter[0])

        stripFilePath := regexp.MustCompile(`^.*\.(.*)$`)
        fnName = stripFilePath.ReplaceAllString(functionObject.Name(), "$1")
    } else {
        fnName = ss[0]
    }

    fmt.Printf("%sENTER: %s\n", getDepth(), fnName)
    return fnName
}

func _exit(s string) {
    _decrement()
    fmt.Printf("%sEXIT:  %s\n", getDepth(), s)
}

func GetTraceFunctions(opts Options) (func(string), func(...string) string) {
    fmt.Printf("%v", opts)
    return _exit, _enter
}
