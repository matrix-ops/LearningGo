package main

import (
	"TestPackage"
	"fmt"
)

func main() {
	fmt.Println(TestPackage.TestVar)
	// var f1 TestPackage.FileTest
	// f1.fd =1
	// 由于FileTest内嵌的fileTest类型并未被导出，所以此处无法直接在TestPackage包外访问fileTest的字段
	// }
}
