package main

import "fmt"

func main() {
	var array1 [3]*string
	array2 := [3]*string{new(string), new(string), new(string)}
	fmt.Println(array1)
	*array2[0] = "zhangweilong"
	*array2[1] = "Blue"
	*array2[2] = "Red"
	array1 = array2
	fmt.Println(*array1[0])
}
