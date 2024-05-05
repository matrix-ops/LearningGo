package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var moneyArray = [3]int{5, 10, 25}
	var totalMoney int = 0
	var Money int = 0
	rand.Seed(time.Now().UnixNano())
	for totalMoney < 2000 {
		Money = moneyArray[rand.Intn(len(moneyArray))]
		totalMoney += Money
		totalMoneyFloat := float64(totalMoney)
		fmt.Printf("%.2f\n", totalMoneyFloat/100.00)
	}
}
