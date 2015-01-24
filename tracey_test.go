package tracey

import (
    "log"
    "bytes"

    "testing"
    "github.com/stretchr/testify/assert"
)

// Define a custom logger which is just a string buffer
// so we can validate that we are building the appropriate
// trace messages
var TestBuffer bytes.Buffer
var BufLogger = log.New(&TestBuffer, "", 0)

func GetTestBuffer() string {
    return "\n" + TestBuffer.String()
}
func ResetTestBuffer() {
    TestBuffer.Reset()
}

func TestBasicUsage(test *testing.T) {
    ResetTestBuffer()
    G, O := GetTraceFunctions(Options{ CustomLogger: BufLogger })

    second := func() {
        defer G(O("SECOND"))
    }
    first := func() {
        defer G(O("FIRST"))
        second()
    }
    first()

    assert.Equal(test, GetTestBuffer(),`
[ 0]ENTER: FIRST
[ 1]  ENTER: SECOND
[ 1]  EXIT:  SECOND
[ 0]EXIT:  FIRST
`)
}

func TestCustomEnterExit(test *testing.T) {
    // Define a custom logger which is just a string buffer
    // so we can validate it :)
    ResetTestBuffer()
    G, O := GetTraceFunctions(Options{ CustomLogger: BufLogger, EnterMessage: "enter: ", ExitMessage: "exit:  " })

    second := func() {
        defer G(O("SECOND"))
    }
    first := func() {
        defer G(O("FIRST"))
        second()
    }
    first()

    assert.Equal(test, GetTestBuffer(),`
[ 0]enter: FIRST
[ 1]  enter: SECOND
[ 1]  exit:  SECOND
[ 0]exit:  FIRST
`)
}
