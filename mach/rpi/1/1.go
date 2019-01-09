// Why fork 2 if you can go 1?
package main

// This is a test program only.

import (
	"fmt"
	"os"
	"time"

	"github.com/ktye/iv/apl"
	"github.com/ktye/iv/apl/numbers"
	"github.com/ktye/iv/apl/operators"
	"github.com/ktye/iv/apl/primitives"
	"github.com/ktye/sfb/fb"
)

var buf *[]uint8
var on, of []uint8

func init() {
	a()
	im, err := fb.Open("/dev/fb0")
	if err != nil {
		fatal(err)
	}
	fmt.Printf("%T\n", im)

	size := im.Bounds().Dx() * im.Bounds().Dy()
	on = make([]uint8, size)
	of = make([]uint8, size)

	// TODO: point puf to the underlying Pix struct field.
}

func main() {
	return // TODO remove

	ticker := time.NewTicker(500 * time.Millisecond)
	never := make(chan bool)
	go func(c chan<- bool) {
		var s bool
		for range ticker.C {
			toggle(&s)
		}
	}(nil)
	<-never
}

func toggle(s *bool) {
	if *s {
		copy(*buf, on)
	} else {
		copy(*buf, of)
	}
	*s = !*s
}

func fatal(err error) {
	os.Stdout.Write([]byte(err.Error()))
	os.Exit(1)
}

// Adding APL increases the size from 1.8 to 4.1 MB!
// with GOOS=linux GOARCH=arm go build (go.1.11.1)
// $ echo 41-18 | apl
// 23
func a() {
	a := apl.New(devnull{})
	numbers.Register(a)
	primitives.Register(a)
	operators.Register(a)
	a.ParseAndEval("1+1")
}

type devnull struct{}

func (n devnull) Write(p []byte) (int, error) {
	return len(p), nil
}
