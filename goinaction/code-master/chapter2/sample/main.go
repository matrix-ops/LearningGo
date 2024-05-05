package main

import (
	"fmt"
	_ "github.com/goinaction/code/chapter2/sample/matchers"
	"github.com/goinaction/code/chapter2/sample/search"
	"log"
	"os"
)

// init is called prior to main.
// init在main函数之前被调用
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

// main is the entry point for the program.
func main() {
	// Perform the search for the specified term.
	fmt.Println(search.TestVar)
	search.Run("rss")
}
