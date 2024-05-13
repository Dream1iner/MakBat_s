package main

import (
	"fmt"
	"strconv"
	"time"

	"github.com/distatus/battery"
	"github.com/getlantern/systray"
)

var batChargeStr string

func onReady() {

	go func() {

		for {
			getData()
			systray.SetTitle(batChargeStr)
			time.Sleep(5 * time.Second)
		}
	}()

	mQuit := systray.AddMenuItem("Quit", "Quit the application")
	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func onExit() {

}

func getData() {
	batteries, err := battery.GetAll()

	if err != nil {
		fmt.Println("Could not get battery info:", err)
		return
	}

	for _, battery := range batteries {
		batCharge := int(float64(battery.Current) / float64(battery.Full) * 100)
		newBatChargeStr := strconv.Itoa(batCharge) + "%"
		// fmt.Println("Battery state:", battery.State.String()) // Debug
		if battery.State.String() == "Charging" {
			// fmt.Println("Battery is charging") // Debug
			newBatChargeStr = "⚡ " + newBatChargeStr
		}
		if battery.State.String() == "Full" {
			newBatChargeStr = "Full"
		}

		if newBatChargeStr != batChargeStr {
			// fmt.Println("Battery charge updated:", newBatChargeStr) // Debug
			batChargeStr = newBatChargeStr
			//newBatChargeStr = batChargeStr // Update newBatChargeStr

		} else {
			// fmt.Println("Battery charge unchanged") // Debug

		}
	}
}

func main() {

	systray.Run(onReady, onExit)

}
