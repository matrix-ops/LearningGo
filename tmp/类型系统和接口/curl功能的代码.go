package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func init() {
	if len(os.Args) != 2 {
		fmt.Println("Usage : ./curl <curl>")
		os.Exit(1)
	}
}
func main() {
	if !strings.HasPrefix(os.Args[1], "http") && !strings.HasPrefix(os.Args[1], "https") {
		os.Args[1] = "https://" + os.Args[1]
	}
	r, err := http.Get(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}
	io.Copy(os.Stdout, r.Body)
	if err := r.Body.Close(); err != nil {
		fmt.Println(err)
	}
}
