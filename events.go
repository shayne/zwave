package main

import (
	"fmt"

	"github.com/ninjasphere/go-openzwave"
	"github.com/shayne/zwave/devices"
)

// EventsCallback yah
func EventsCallback(api openzwave.API, event openzwave.Event) {
	switch event.(type) {
	case *openzwave.NodeAvailable:
		fmt.Println("NodeAvailable!")
		node := event.GetNode()
		device := node.GetDevice()
		// Hack to detect dimmer and set its value to 0
		if dimmer, ok := device.(*devices.DimmerDevice); ok {
			dimmer.ChangeValue(0)
		}
	case *openzwave.NodeChanged:
		fmt.Println("NodeChanged!")
	case *openzwave.NodeUnavailable:
		fmt.Println("NodeUnavailable!")
	}
}
