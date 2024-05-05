 常量的默认类型是big.Int

strconv.Atoi函数的作用并不是说你给定一个ASCII上的字符，它帮你转换成Int类型，而是把一个可能在ASCII上有转义作用的数字原封不动转成Int类型。

strconv.Itoa函数的作用则刚好相反，它将一个数字直接转换成字符串，而不是转换成这个数字对应的字符值

range函数用于返回数组、切片、通道和集合中的元素，第一个返回值是索引，第二个是值







华氏度转摄氏度

F-32/1.8

摄氏度转华氏度

C+32*1.8



###### 方法被调用的时候前面需要加上它所属于的类型名

```go
package main
import "fmt"
type kelvin float64
type celsius float64
func (k kelvin) kelvinToCelsius() celsius {
    return celsius(k - 273.15)
}

func main() {
    var k kelvin = 10.00
    fmt.Println(k.kelvinToCelsius())
}
```

###### 方法的意义

属于同一类型的变量，都能使用属于此类型的方法



###### 把函数赋值给变量

```go
package main
import "fmt"
type kelvin float64
type celsius float64

func kelvinToCelsius(k kelvin) celsius {
    return celsius(k - 273.15)
}

func main() {
    var kToC func(k kelvin) celsius
    // 在不使用类型推断的情况下，变量的类型声明必须和函数的形参和返回值一致
    kToC = kelvinToCelsius
    fmt.Println(kToC(10.00))
}
```



###### 定义函数类型

```go
type kelvinToCelsius func() celsius
// 与上同，如果kelvin对应的函数原型中具有形参和返回值，此处也要写上
type kelvinToCelsius func(k kelvin) celsius
// 就像这样
```



###### strings.Join

原型：

```go
func Join(s []string sep)  string	
```

将字符串切片中的所有元素合并成一个字符串，可指定sep作为元素间的分隔符



Sprintf返回一个格式化字符串

Sprintf函数不会立即打印它的输出，它可以被赋值给一个变量，当使用其他的方式打印这个变量时，Sprintf函数中的内容将被打印

Fprintln将字符串写入一个io.Writer类型，可以是stdout



接口

并且某个类型实现了接口定义的所有方法，这个类型的值就可以被赋值给接口类型的变量

接口类型的变量，它的值，可以是实现了这个接口类型的其他类型的值

一个符合条件的变量，可以被赋值给接口，不一定要字面量

只要这个变量的类型，实现了接口声明中的所有方法，注意方法签名还要一样



error接口的意义在于，一个类型，只要它有一个Error() string方法，那它的变量就能被当做error返回值来用。
