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

func _enter(args ...interface{}) string {
    defer _increment()
    traceMessage := ""

    // Figure out the name of the caller and use that instead
    pc := make([]uintptr, 2)
    runtime.Callers(2, pc)
    fObject := runtime.FuncForPC(pc[0])
    fnName := regexp.MustCompile(`^.*\.(.*)$`).ReplaceAllString(fObject.Name(), "$1")

    if len(args) > 0 {
        if fmtStr, ok := args[0].(string); ok {
            // "$FN" will be replaced by the name of the function
            // We have a string leading args, assume its to be formatted
            traceMessage = fmt.Sprintf(fmtStr, args[1:]...)
        }
    } else {
        traceMessage = fnName;
    }

    traceMessage = regexp.MustCompile(`\$FN`).ReplaceAllString(traceMessage, fnName)

    DefaultLogger.Printf("%s%s%s\n", getDepth(), EnterMessage, traceMessage)
    return traceMessage
}

func _exit(traceMessage string) {
    _decrement()
    DefaultLogger.Printf("%s%s%s\n", getDepth(), ExitMessage, traceMessage)
}

func GetTraceFunctions(opts Options) (func(string), func(...interface{}) string) {
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
