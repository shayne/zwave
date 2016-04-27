package main

import (
	"fmt"

	"github.com/shayne/zwave/go-openzwave"
	"github.com/shayne/zwave/devices"
)

func eventsCallback(api openzwave.API, event openzwave.Event) {
	switch event.(type) {
	case *openzwave.NodeAvailable:
		fmt.Println("NodeAvailable!")
		node := event.GetNode()
		device := node.GetDevice()
		// Add node to any device that supports it
		if dimmer, ok := device.(*devices.DimmerDevice); ok {
			deviceMap[dimmer.Node.GetNodeName()] = dimmer
		}
	case *openzwave.NodeChanged:
		fmt.Println("NodeChanged!")
	case *openzwave.NodeUnavailable:
		fmt.Println("NodeUnavailable!")
	}
}
