package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"Name ,auto,omitempty"`
	// 这个auto似乎并没有实际意义
	Male string `json:""`
	// 留空，表示使用和结构体同名的字段来解析json
}

func main() {
	p1 := People{
		Name: "zhangweilong",
		Male: "Man",
	}
	j1, err := json.Marshal(p1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(j1))
	fmt.Println(p1)

}
