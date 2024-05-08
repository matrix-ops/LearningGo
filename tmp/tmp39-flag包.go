package main

import (
	"flag"
	"fmt"
)

var (
	httpPort = flag.Int("httpPort", 8080, "http port")
	url      = flag.String("url", "google", "www.google.com")
)

func main() {
	flag.Parse()
	fmt.Println(*url)
}
