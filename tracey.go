// `Tracey` is a simple library which allows for much easier function enter / exit logging
package tracey

import (
    "os"
    "fmt"
    "log"
    "regexp"
    "strings"
    "strconv"

    "reflect"
    "runtime"
)

type Options struct {
    DisableNesting      bool
    DisableDepthValue   bool

    CustomLogger        *log.Logger

    SpacesPerIndent     int    `default:"2"`
    EnterMessage        string `default:"ENTER: "`
    ExitMessage         string `default:"EXIT:  "`

    currentDepth        int
}
var options Options

func getDepth() string {
    spaces := strings.Repeat(" ", options.currentDepth * options.SpacesPerIndent)
    if !options.DisableDepthValue {
        return fmt.Sprintf("[%2d]%s", options.currentDepth, spaces)
    }
    return spaces
}

func _increment() {
    options.currentDepth += 1
}

func _decrement() {
    options.currentDepth -= 1
    if options.currentDepth < 0 {
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

    options.CustomLogger.Printf("%s%s%s\n", getDepth(), options.EnterMessage, traceMessage)
    return traceMessage
}

func _exit(s string) {
    _decrement()
    options.CustomLogger.Printf("%s%s%s\n", getDepth(), options.ExitMessage, s)
}

func GetTraceFunctions(opts Options) (func(string), func(...interface{}) string) {
    options = opts

    if options.CustomLogger == nil {
        options.CustomLogger = log.New(os.Stdout, "", 0)
    }

    reflectedType := reflect.TypeOf(options)
    if options.EnterMessage == "" {
        field, _ := reflectedType.FieldByName("EnterMessage")
        options.EnterMessage = field.Tag.Get("default")
    }
    if options.ExitMessage == "" {
        field, _ := reflectedType.FieldByName("ExitMessage")
        options.ExitMessage = field.Tag.Get("default")
    }

    if options.DisableNesting {
        options.SpacesPerIndent = 0
    } else if options.SpacesPerIndent == 0 {
        field, _ := reflectedType.FieldByName("SpacesPerIndent")
        options.SpacesPerIndent, _ = strconv.Atoi(field.Tag.Get("default"))
    }

    return _exit, _enter
}
