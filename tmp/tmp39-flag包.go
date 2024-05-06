package main

import "flag"

var (
	httpPort = flag.Int("httpPort", 8080, "http port")
	help     = flag.Bool("help", false, "show help")
)

func main() {
	flag.Parse()
}
