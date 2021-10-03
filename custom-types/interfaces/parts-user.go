package main

import (
	"fmt"
	"go-101/parts"
)

func MonitorSummary(monitor parts.Monitor) string {
	return monitor.Specs() + "\n" + monitor.Price()
}

func HardDriveSummary(hardDrive parts.HardDrive) string {
	return hardDrive.Specs() + "\n" + hardDrive.Price()
}

func KeyboardSummary(keyboard parts.Keyboard) string {
	return keyboard.Specs() + "\n" + keyboard.Price()
}

func main(){
	monitor := parts.Monitor{
		Resolution:"1080p",
		Connector: "HDMI",
		Value: 249.99,
	}

	fmt.Println(MonitorSummary(monitor))
	fmt.Println("---------------------------")

	hardDrive := parts.HardDrive{
		Type: "SSD",
		Connector: "SATA",
		Value: 149.99,
	}

	fmt.Println(HardDriveSummary(hardDrive))
	fmt.Println("---------------------------")

	keyboard := parts.Keyboard{
		Keys: 108,
		Switches: "Mechanical",
		Value: 99.9,
	}

	fmt.Println(KeyboardSummary(keyboard))
	fmt.Println("---------------------------")
}