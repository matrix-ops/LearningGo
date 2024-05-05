package main

import (
	"encoding/json"
	"fmt"
)

type location struct {
	Name string  `json:"LocationName""`
	Lat  float64 `json:"Lat""`
	Long float64 `json:"Long"`
}

func main() {
	//var marsLocation location
	locations := []location{
		{Name: "Bradbury Landing", Lat: -4.5895, Long: 137.4417},
		{Name: "Columbia Memorial Station", Lat: -14.5684, Long: 175.472636},
		{Name: "Challenger Memorial Station", Lat: -1.9462, Long: 354.4734},
	}

	for i, _ := range locations {
		//marsLocation = locations[i]
		jsonLocation, err := json.MarshalIndent(locations[i], "", "\t")
		fmt.Println(string(jsonLocation))
		if err != nil {
			fmt.Println(err)
		}
	}
}
