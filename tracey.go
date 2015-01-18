// `Tracey` is a simple library which allows for much easier function enter / exit logging
package tracey

import (
    "os"
    "fmt"
    "log"
    "regexp"
    "strings"

    "reflect"
    "runtime"
)

var SPACES_PER_TAB = 2
var DEPTH = 0
var DefaultLogger = log.New(os.Stdout, "", 0)
var EnterMessage, ExitMessage string

type Options struct {
    SpacesPerIndent int
    EnableNesting   bool
    PrintDepthValue bool

    CustomLogger    *log.Logger
    EnterMessage    string `default:"ENTER: "`
    ExitMessage     string `default:"EXIT:  "`
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

    DefaultLogger.Printf("%s%s%s\n", getDepth(), EnterMessage, fnName)
    return fnName
}

func _exit(s string) {
    _decrement()
    DefaultLogger.Printf("%s%s%s\n", getDepth(), ExitMessage, s)
}

func GetTraceFunctions(opts Options) (func(string), func(...string) string) {
    if opts.CustomLogger != nil {
        DefaultLogger = opts.CustomLogger
    }

    reflectedType := reflect.TypeOf(opts)
    EnterMessage = opts.EnterMessage
    if EnterMessage == "" {
        field, _ := reflectedType.FieldByName("EnterMessage")
        EnterMessage = field.Tag.Get("default")
    }
    ExitMessage = opts.ExitMessage
    if ExitMessage == "" {
        field, _ := reflectedType.FieldByName("ExitMessage")
        ExitMessage = field.Tag.Get("default")
    }
    return _exit, _enter
}
