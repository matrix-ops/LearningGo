package main

import "fmt"

type celsiusType float64
type kelvinType float64

func kelvinToCelsius(k float64) float64 {
	k -= 273.15
	return k
}

func celsiusToFahrenheit(c float64) float64 {
	f := (c * 9.0 / 5.0) + 32.0
	return f
}

func kelvinToFahrenheit(k float64) float64 {
	k -= 459.67
	return k
}

func celsiusToKelvin(c celsiusType) kelvinType {
	k := c + 273.15
	return kelvinType(k)
}

func main() {
	var celsius celsiusType = 127.00
	kelvin := celsiusToKelvin(celsius)
	fmt.Println(kelvin)
	//kelvin := 294.0
	//kelvin1 := 233.0
	//kelvin2 := 0.0
	//celsius := kelvinToCelsius(kelvin)
	//celsius1 := kelvinToCelsius(kelvin1)
	//fahrenheit := celsiusToFahrenheit(celsius)
	//fahrenheit1 := kelvinToFahrenheit(kelvin2)
	//fmt.Printf("摄氏度: %.2f℃, %.2f℃\n", celsius, celsius1)
	//fmt.Printf("%.2f\n", fahrenheit)
	//fmt.Printf("%.2f\n", fahrenheit1)
}
