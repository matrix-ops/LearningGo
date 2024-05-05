package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type JsonStruct struct {
	Name    string `json:"Name"`
	Address string `json:"Address"`
	age     int
}

func main() {
	j1 := JsonStruct{
		Name:    "zhangweilong",
		Address: "GuangDong GuangZhou",
		age:     19,
	}
	fmt.Println(j1)
	data, err := json.Marshal(j1)
	// 将结构体序列化为json
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%s\n", data)
	fmt.Println(reflect.TypeOf(data))
	// 无符号八位数整型切片

}
