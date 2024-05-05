package main

import "fmt"

type location struct {
	lat, long float64
}

func (l location) String() string  {
	return fmt.Sprintf("%v %v", l.lat, l.long)
}

func main()  {
	curiosity := location{lat: 1.23, long: 4.56}
	fmt.Println(curiosity.String())
}
