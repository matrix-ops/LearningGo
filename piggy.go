package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var totalMoney float64 = 0.0
	var Money float64 = 0.0
	var moneyArray = [3]float64{0.05, 0.1, 0.25}
	rand.Seed(time.Now().Unix())
	for totalMoney < 20.00 {
		Money = moneyArray[rand.Intn(len(moneyArray))]
		totalMoney += Money
		fmt.Printf("%.2f\n", totalMoney)
	}
}
