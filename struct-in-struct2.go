package main

import "fmt"

type report struct {
	sol int
	temperature
	location
}

type temperature struct {
	hight, low celsius
}

type location struct {
	lat, long float64
}

type celsius float64

func main()  {
	bradbury := location{-4.5895, 137.4417}
	t := temperature{-1.0, -78.0}
	report := report{15,t,bradbury}
	fmt.Printf("%+v", report)
}
