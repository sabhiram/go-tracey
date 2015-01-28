# go-tracey

[![Build Status](https://travis-ci.org/sabhiram/go-tracey.svg?branch=master)](https://travis-ci.org/sabhiram/go-tracey) [![Coverage Status](https://coveralls.io/repos/sabhiram/go-tracey/badge.svg?branch=master)](https://coveralls.io/r/sabhiram/go-tracey?branch=master)

A simple function tracer for Go

## Install

```sh
go get github.com/sabhiram/go-tracey
```

## Basic Usage

*file: foo.go*
```go
package main

import (
    "github.com/sabhiram/go-tracey"
)

// Setup global enter exit trace functions (default options)
var G, O = tracey.New(nil)
// Or...
// var Exit, Enter = tracey.New(nil)

func foo(i int) {
    // $FN will get replaced with the function's name
    defer G(O("$FN(%d)", i))
    if i != 0 {
        foo(i - 1)
    }
}

func main() {
    defer G(O())
    foo(2)
}
```

```sh
go run foo.go 
[ 0]ENTER: main
[ 1]  ENTER: foo(2)
[ 2]    ENTER: foo(1)
[ 3]      ENTER: foo(0)
[ 3]      EXIT:  foo(0)
[ 2]    EXIT:  foo(1)
[ 1]  EXIT:  foo(2)
[ 0]EXIT:  main
```

## Configurable Tracey Options

```go
// These options represent the various settings which tracey exposes.
// A pointer to this structure is expected to be passed into the
// `tracey.New(...)` function below.
type Options struct {

    // Setting "DisableTracing" to "true" will cause tracey to return
    // no-op'd functions for both exit() and enter(). The default value
    // for this is "false" which enables tracing.
    DisableTracing      bool

    // Setting the "CustomLogger" to nil will cause tracey to log to
    // os.Stdout. Otherwise, this is a pointer to an object as returned
    // from `log.New(...)`.
    CustomLogger        *log.Logger

    // Setting "DisableDepthValue" to "true" will cause tracey to not
    // prepend the printed function's depth to enter() and exit() messages.
    // The default value is "false", which logs the depth value.
    DisableDepthValue   bool

    // Setting "DisableNesting" to "true" will cause tracey to not indent
    // any messages from nested functions. The default value is "false"
    // which enables nesting by prepending "SpacesPerIndent" number of
    // spaces per level nested.
    DisableNesting      bool
    SpacesPerIndent     int    `default:"2"`

    // Setting "EnterMessage" or "ExitMessage" will override the default
    // value of "Enter: " and "EXIT:  " respectively.
    EnterMessage        string `default:"ENTER: "`
    ExitMessage         string `default:"EXIT:  "`

}
```

## Want to help out?

I appreciate any and all feedback. There is no "real" coding standard. I am still finding my feet in Go and am not sure what I like and abhor yet.

Please run the tests and check out the examples for more details. All tests will be run against TravisCI and then coverage results will be forwarded to coveralls.io. 

To run tests:
```sh
$cd go-tracey
$go test -v
```

[Link to tests for the lazy](https://github.com/sabhiram/go-tracey/blob/master/tracey_test.go)
