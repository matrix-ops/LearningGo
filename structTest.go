package main

import (
	"fmt"
	// "encoding/json"
)

type Employee struct {
	ID      int
	Name    string
	Address string
}

func main() {
	var zhangweilong Employee
	var EmployeeZhang *Employee = &zhangweilong
	zhangweilong.ID = 1
	zhangweilong.Name = "zhangweilong"
	zhangweilong.Address = "GuangDong GuangZhou LiWan ZhongNan"
	x := &zhangweilong.Address
	*x = "ZhongGuo " + *x
	zhangweilong.Address = *x
	fmt.Println(zhangweilong.Address)
	fmt.Println(EmployeeZhang.Address)
	fmt.Printf("%+v\n", zhangweilong)
	// 会输出两行一模一样的地址信息
	// str, _ := json.Marshal(zhangweilong)
	// fmt.Println(string(str))

	var huangyuting Employee = Employee{
		12,
		"huangyuting",
		"GuangDong  GuangZhou",
	}
	fmt.Printf("%+v", huangyuting)
}
