package tracey

import (
	"bytes"
	"log"

	"github.com/stretchr/testify/assert"
	"testing"
)

// Define a custom logger which is just a string buffer
// so we can validate that we are building the appropriate
// trace messages
var TestBuffer bytes.Buffer
var BufLogger = log.New(&TestBuffer, "", 0)

func GetTestBuffer() string {
	return TestBuffer.String()
}

func ResetTestBuffer() {
	// This prepends a newline in front of the buffer after a reset
	// since our validation logic below opens the ` on the previous line
	// to make the validation logic easier to read, and copy in and out
	TestBuffer.Reset()
	BufLogger.Printf("\n")
}

func TestBasicUsage(test *testing.T) {
	ResetTestBuffer()
	G, O := New(&Options{CustomLogger: BufLogger})

	second := func() {
		defer G(O("SECOND"))
	}
	first := func() {
		defer G(O("FIRST"))
		second()
	}
	first()

	assert.Equal(test, GetTestBuffer(), `
[ 0]ENTER: FIRST
[ 1]  ENTER: SECOND
[ 1]  EXIT:  SECOND
[ 0]EXIT:  FIRST
`)
}

func TestDisableTracing(test *testing.T) {
	ResetTestBuffer()
	G, O := New(&Options{CustomLogger: BufLogger, DisableTracing: true})

	second := func() {
		defer G(O("SECOND"))
	}
	first := func() {
		defer G(O("FIRST"))
		second()
	}
	first()

	assert.Equal(test, GetTestBuffer(), "\n")
}

func TestCustomEnterExit(test *testing.T) {
	ResetTestBuffer()
	G, O := New(&Options{CustomLogger: BufLogger, EnterMessage: "enter: ", ExitMessage: "exit:  "})

	second := func() {
		defer G(O("SECOND"))
	}
	first := func() {
		defer G(O("FIRST"))
		second()
	}
	first()

	assert.Equal(test, GetTestBuffer(), `
[ 0]enter: FIRST
[ 1]  enter: SECOND
[ 1]  exit:  SECOND
[ 0]exit:  FIRST
`)
}

func TestDisableNesting(test *testing.T) {
	ResetTestBuffer()
	G, O := New(&Options{CustomLogger: BufLogger, DisableNesting: true})

	second := func() {
		defer G(O("SECOND"))
	}
	first := func() {
		defer G(O("FIRST"))
		second()
	}
	first()

	assert.Equal(test, GetTestBuffer(), `
[ 0]ENTER: FIRST
[ 1]ENTER: SECOND
[ 1]EXIT:  SECOND
[ 0]EXIT:  FIRST
`)
}

func TestCustomSpacesPerIndent(test *testing.T) {
	ResetTestBuffer()
	G, O := New(&Options{CustomLogger: BufLogger, SpacesPerIndent: 3})

	second := func() {
		defer G(O("SECOND"))
	}
	first := func() {
		defer G(O("FIRST"))
		second()
	}
	first()

	assert.Equal(test, GetTestBuffer(), `
[ 0]ENTER: FIRST
[ 1]   ENTER: SECOND
[ 1]   EXIT:  SECOND
[ 0]EXIT:  FIRST
`)
}

func TestDisableDepthValue(test *testing.T) {
	ResetTestBuffer()
	G, O := New(&Options{CustomLogger: BufLogger, DisableDepthValue: true})

	second := func() {
		defer G(O("SECOND"))
	}
	first := func() {
		defer G(O("FIRST"))
		second()
	}
	first()

	assert.Equal(test, GetTestBuffer(), `
ENTER: FIRST
  ENTER: SECOND
  EXIT:  SECOND
EXIT:  FIRST
`)
}

// Helper function - part of "TestUnspecifiedFunctionName"
func foobar() {
	G, O := New(&Options{CustomLogger: BufLogger})
	defer G(O())
}
func TestUnspecifiedFunctionName(test *testing.T) {
	ResetTestBuffer()

	// Call another named function
	foobar()

	assert.Equal(test, GetTestBuffer(), `
[ 0]ENTER: foobar
[ 0]EXIT:  foobar
`)
}

// Negative tests
func TestMoreExitsThanEntersMustPanic(test *testing.T) {
	G, _ := New(&Options{CustomLogger: BufLogger})
	assert.Panics(test, func() {
		G("")
	}, "Calling exit without enter should panic")
}

// Examples
func ExampleNew_noOptions() {
	G, O := New(nil)

	second := func() {
		defer G(O("SECOND"))
	}
	first := func() {
		defer G(O("FIRST"))
		second()
	}
	first()

	// Output:
	// [ 0]ENTER: FIRST
	// [ 1]  ENTER: SECOND
	// [ 1]  EXIT:  SECOND
	// [ 0]EXIT:  FIRST
}

func ExampleNew_customMessage() {
	G, O := New(&Options{EnterMessage: "en - ", ExitMessage: "ex - "})

	second := func() {
		defer G(O("SECOND"))
	}
	first := func() {
		defer G(O("FIRST"))
		second()
	}
	first()

	// Output:
	// [ 0]en - FIRST
	// [ 1]  en - SECOND
	// [ 1]  ex - SECOND
	// [ 0]ex - FIRST
}

func ExampleNew_changeIndentLevel() {
	G, O := New(&Options{SpacesPerIndent: 1})

	second := func() {
		defer G(O("SECOND"))
	}
	first := func() {
		defer G(O("FIRST"))
		second()
	}
	first()

	// Output:
	// [ 0]ENTER: FIRST
	// [ 1] ENTER: SECOND
	// [ 1] EXIT:  SECOND
	// [ 0]EXIT:  FIRST
}
