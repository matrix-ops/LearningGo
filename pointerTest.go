package main

import (
	"fmt"
	"time"
)

type stats struct {
	level             int
	endurance, health int
}

type character struct {
	name  string
	stats stats
}

func reclassity(p *[]string) {
	slice1 := *p
	*p = slice1[0:1]
	// 显式指向切片的指针只能修改切片的长度、容量和起始偏移量
}

func levelUp(s *stats) {
	s.level++
	s.endurance = 42 + (14 * s.level)
	s.health = 5 * s.endurance
}

func addAddress(p string) string {
	return p + " 住在中国"
}

type person1 struct {
	name string
	age  int
}
type person2 struct {
	name string
	age  int
}

func (p1 *person1) addAge() {
	p1.age++
}

func (p2 person2) addAge() person2 {
	p2.age++
	return p2
}

func main() {
	var p *string
	zhang := "weilong"
	p = &zhang
	fmt.Println(p, &zhang)
	// 直接打印p是打印zhang的内存地址值
	fmt.Println(*p)
	// 解引用打印的是内存地址值里面的东西
	zhang = addAddress(*p)
	fmt.Println(zhang)

	p3 := []string{"Earth", "Venus", "Mercury"}
	reclassity(&p3)
	fmt.Println("这是p2", p3)

	p1 := &person1{
		name: "zhangweilong",
		age:  23,
	}

	p2 := person2{
		name: "zhangweilong",
		age:  23,
	}
	p1.addAge()
	fmt.Println("指针操作结构体:", *p1)
	// 如果不解引用，打印的是p1的内存地址
	p2 = p2.addAge()
	fmt.Println("值传递操作结构体，但是有返回值:", p2)
	// 如果不重新赋值给p2，仅执行p2.addAge()不会对p2自身产生任何影响
	const layout = "Mon, Jan 2, 2006"
	day := time.Now()
	day = day.Add(24 * time.Hour)
	fmt.Println(day)
	var var1 int
	fmt.Println(&var1)
	var1++
	fmt.Println(var1)
	fmt.Println(&var1)
	stats1 := stats{
		level:     1,
		endurance: 2,
		health:    3,
	}
	fmt.Println(stats1.level)
	fmt.Println(stats1)

	player := character{name: "motherFucker"}
	levelUp(&player.stats)
	// levelUp函数没有返回任何东西，在它内部直接针对player结构体变量进行了修改
	fmt.Printf("%+v", player.stats)
	t1 := stats{
		level:  1,
		health: 1,
	}
	fmt.Println(t1)
	t1p := &t1
	fmt.Println(t1p)
	t2p := &stats{level: 1, health: 1}
	fmt.Println(t2p)
	fmt.Println(t1p == t2p)
	fmt.Println(*t1p)
	fmt.Println(*t2p)

	var a [8]int
	a[0] = 1
	a[1] = 2
	ap := &a
	fmt.Println(ap)
	a1 := &[8]int{0, 1, 2, 3, 4}
	fmt.Println(a1)
	s1 := &[3]string{"Mercury", "Venus", "Earth"}
	fmt.Println(*s1)
	*s1 = [3]string{"Earth", "Venus", "Mercury"}
	fmt.Println(*s1)
}
