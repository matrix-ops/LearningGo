package main

import (
	"encoding/json"
	"fmt"
)

type locationStruct struct {
	Name string `json:"jiaoshenmen"`
	Lat  float64
	Long float64
}

func main() {
	locations := []locationStruct{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5685, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}
	var l locationStruct
	l = locations[0]
	fmt.Println("locations切片的第一个字段是:")
	fmt.Printf("%+v\n", l)

	for i := 0; i <= 2; i++ {
		bytes, _ := json.MarshalIndent(locations[i], "", "")
		fmt.Println(string(bytes))
		fmt.Println(bytes)
	}
}
