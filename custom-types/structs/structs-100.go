package main

import (
	"fmt"
)

type Monitor struct {
	Resolution	string
	Connector 	string
	Value 		float64
}

func showFields(m Monitor){
	fmt.Println("Resolution :", m.Resolution, "Connector :", m.Connector, "Value :", m.Value)
}

// func (m Monitor) showFields(){
// 	fmt.Println("Resolution :", m.Resolution, "Connector :", m.Connector, "Value :", m.Value)
// }

func main(){
	monitor := Monitor{}

	monitor.Resolution = "1080p"
	monitor.Connector = "HDMI"
	monitor.Value = 249.99

	// fmt.Println("Resolution :", monitor.Resolution, "Connector :", monitor.Connector, "Value :", monitor.Value)
	// showFields(monitor)
	// monitor.showFields()

	// monitor2 := Monitor{}
	// monitor2 := Monitor{"1080p", "HDMI", 249.99} // Literals should be in order, none should be left out.
	monitor2 := Monitor{Connector: "HDMI"} // Literals need not to be in order and can even be left out completely
	showFields(monitor2)
}