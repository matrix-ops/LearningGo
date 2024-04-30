package main

import (
	"fmt"
	"io"
	"os"
)

type safeWriter struct {
	w   io.Writer
	err error
}

func (sw *safeWriter) Writeln(s string) {
	if sw.err != nil {
		return
	}
	_, sw.err = fmt.Fprintln(sw.w, s)
}

func proverbs(name string) error {
	f, err := os.Create(name)
	if err != nil {
		return err
	}
	sw := safeWriter{w: f}
	sw.Writeln("Errors are values.")
	sw.Writeln("Don’t just check errors, handle them gracefully.")
	sw.Writeln("Don't panic.")
	sw.Writeln("Make the zero value useful.")
	sw.Writeln("The bigger the interface, the weaker the abstraction.")
	sw.Writeln("interface{} says nothing.")
	sw.Writeln("Gofmt's style is no one's favorite, yet gofmt is everyone's favorite.")
	sw.Writeln("Documentation is for users.")
	sw.Writeln("A little copying is better than a little dependency.")
	sw.Writeln("Clear is better than clever.")
	sw.Writeln("Concurrency is not parallelism.")

	sw.Writeln("Don’t communicate by sharing memory, share memory by communicating.")
	sw.Writeln("Channels orchestrate; mutexes serialize.")
	return sw.err
}

func main() {
	err := proverbs("./TestProverbs")
	if err != nil {
		fmt.Println(err)
	}
}
