# go-tracey

[![Build Status](https://travis-ci.org/sabhiram/go-tracey.svg?branch=master)](https://travis-ci.org/sabhiram/go-tracey) [![Coverage Status](https://coveralls.io/repos/sabhiram/go-tracey/badge.svg?branch=master)](https://coveralls.io/r/sabhiram/go-tracey?branch=master)

A simple function tracer for Go

## Install

```sh
go get github.com/sabhiram/go-tracey
```

## Sample Usage

```sh
$go get github.com/sabhiram/go-tracey
$touch foo.go
```

*file: foo.go*
```go
package main

import (
    "github.com/sabhiram/go-tracey"
)

// Setup global enter exit trace functions (default options)
var G, O = tracey.GetTraceFunctions(tracey.Options{})

func foo() {
    defer G(O())
    bar()
}

func bar() {
    defer G(O())
    baz()
}

func baz() {
    defer G(O())
}

func main() {
    defer G(O())

    foo()
    bar()
    baz()
}
```

```sh
$go run foo.go
[ 0]ENTER: main
[ 1]  ENTER: foo
[ 2]    ENTER: bar
[ 3]      ENTER: baz
[ 3]      EXIT:  baz
[ 2]    EXIT:  bar
[ 1]  EXIT:  foo
[ 1]  ENTER: bar
[ 2]    ENTER: baz
[ 2]    EXIT:  baz
[ 1]  EXIT:  bar
[ 1]  ENTER: baz
[ 1]  EXIT:  baz
[ 0]EXIT:  main
```
