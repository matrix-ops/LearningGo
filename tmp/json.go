package main

import (
	"encoding/json"
	"fmt"
)

type People struct {
	Name string `json:"PeopleName,omitempty"`
	// 将Name字段映射为PeopleName字段，同时如果此字段的值为空或者不存在此字段，那么不解析到JSON里面
	Male string `json:""`
	// 留空，表示使用和结构体同名的字段来解析json
}

func main() {
	p1 := People{
		Name: "zhangweilong",
		Male: "Man",
	}
	p2 := People{
		Name: "ZhangTingTing",
	}
	// Male 没有使用omitempty,即使为空，仍然解析成了json
	p3 := People{
		Male: "Women",
	}
	j1, err := json.Marshal(p1)
	j2, err := json.Marshal(p2)
	j3, err := json.Marshal(p3)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(j1))
	// {"Name ":"zhangweilong","Male":"Man"}
	fmt.Println(string(j2))
	// {"Name ":"ZhangTingTing","Male":""}
	fmt.Println(string(j3))
	// {"Male":"Women"}
	// 由于Name字段是omitempty，所以在p3没有定义Name字段的情况下json里面不会出现Name字段
	fmt.Println(p1)

}
