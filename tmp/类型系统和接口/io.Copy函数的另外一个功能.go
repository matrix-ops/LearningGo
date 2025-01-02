package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
)

func main() {
	var b bytes.Buffer
	b.Write([]byte("hello"))
	fmt.Fprintf(&b, " world")
	// 这里之所以写成指针形式, 是因为此处需要直接修改b
	// Fprintf并没有规定第一个参数应该是指针类型还是值类型
	// Fprintf只规定了第一个参数得是io.Writer接口
	// bytes.Buffer实现了Write()方法，因此符合io.Writer这个接口要求。
	// 但是bytes.Buffer在实现Write方法的时候指定了指针作为接受者
	// 因此只有bytes.Buffer类型的指针才满足io.Write方法
	// func (b *Buffer) Write(p []byte) (n int, err error)
	// func (b *Buffer) Read(p []byte) (n int, err error)
	io.Copy(os.Stdout, &b)
}
