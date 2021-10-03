package main

import (
	"fmt"
	"go-101/parts"
)

type Part interface {
	Specs()	string
	Price()	string
}

func Summary (part Part) string {
	return part.Specs() + "\n" + part.Price()
}

func main(){
	catalog := [] Part {}
	
	catalog = append(catalog, parts.Monitor{Resolution: "1080p", Connector: "HDMI", Value: 249.99})
	catalog = append(catalog, parts.HardDrive{Type: "SSD", Connector: "SATA", Value: 149.99})
	catalog = append(catalog, parts.Keyboard{Keys: 108, Switches: "Mechanical", Value: 99.9})

	for _, part := range catalog {
		fmt.Println(Summary(part))
		fmt.Println("---------------------------")
	}
}