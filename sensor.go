package main

import (
	"fmt"
	"math/rand"
	"time"
)

type keLvin float64

func fakeSensor() keLvin {
	return keLvin(rand.Intn(151) + 150)
}

func realSensor() keLvin {
	return 0
}

func measureTemperature(samples int, sensor func() keLvin) {
	for i := 0; i < samples; i++ {
		k := sensor()
		fmt.Println(k, " K")
		time.Sleep(time.Second)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	measureTemperature(3, fakeSensor)
}
