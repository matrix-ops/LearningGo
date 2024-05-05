package main

import "fmt"

type celsius float64
type fahrenheit float64
type kelvin float64

func (f fahrenheit) celsius() celsius {
	f = (f - 32) / 1.8
	return celsius(f)
}

func (k kelvin) celsius() celsius {
	k = k - 273.15
	return celsius(k)
}

func (c celsius) fahrenheit() fahrenheit {
	c = (c + 32) * 1.8
	return fahrenheit(c)
}

func (k kelvin) fahrenheit() fahrenheit {
	k = k*1.8 - 459.67
	return fahrenheit(k)
}

func (c celsius) kelvin() kelvin {
	c = c + 273.15
	return kelvin(c)
}

func (f fahrenheit) kelvin() kelvin {
	f = (f + 459.67) / 1.8
	return kelvin(f)
}
func main() {
	var k kelvin = 100
	f := k.fahrenheit()
	fmt.Println(f)
}
