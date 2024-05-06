package main

type Person38 struct {
	age int
	name string
}

func newPerson(age int, name string) *Person {
	p := Person{age: age, name: name}
	return &p
}

func main()  {
	p1 := newPerson(18,"zhangweilong")	
	fmt.Println(p1)
}

