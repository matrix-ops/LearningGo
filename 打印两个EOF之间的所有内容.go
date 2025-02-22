package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func processFile(filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("无法打开文件 %s: %v", filePath, err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)

	var (
		eofCount   int
		startPrint bool
		builder    strings.Builder
	)

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatalf("读取文件出错 %s: %v", filePath, err)
		}

		if strings.TrimSpace(line) == "EOF" {
			eofCount++
		}

		if eofCount == 1 {
			startPrint = true
		} else if eofCount == 2 {
			startPrint = false
			break
		}

		if startPrint {
			builder.WriteString(line)
		}

		if err == io.EOF {
			break
		}
	}

	if eofCount < 2 {
		fmt.Printf("文件 %s 中没有找到两个 EOF\n", filePath)
	} else {
		fmt.Printf("文件 %s 中从第一个 EOF 到第二个 EOF 之间的内容:\n%s\n", filePath, builder.String())
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("用法: go run program.go <文件路径1> <文件路径2> ...")
		return
	}

	for _, filePath := range os.Args[1:] {
		processFile(filePath)
	}
}
