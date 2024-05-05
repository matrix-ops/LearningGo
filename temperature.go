package main

import "fmt"

type report struct {
	sol int
	location	location
	temperature temperature
}

type location struct {
	lat, long float64
}

type temperature struct {
	high, low celsius
}

type celsius float64

func main()  {
	location := location{lat: 2.5, long: 3.5}
	temperature := temperature{high: 3.2, low: 1.2}
	sol := 15
	report := report{sol: sol,temperature: temperature, location: location}
	fmt.Printf("%+v", report)
}
