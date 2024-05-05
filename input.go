package main

import "fmt"

func main() {

	yesNo := "fuckyou"
	var launch bool
	fmt.Println(launch)
	switch yesNo {
	case "true", "yes", "1":
		launch = true
	case "false", "no", "0":
		launch = false
	default:
		fmt.Println("值有误")
	}
}
