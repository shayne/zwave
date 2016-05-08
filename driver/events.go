package driver

import (
	"fmt"

	"github.com/shayne/zwave/device"
	"github.com/shayne/zwave/go-openzwave"
)

func eventsCallback(deviceMap map[interface{}]interface{}) func(api openzwave.API, event openzwave.Event) {
	return func(api openzwave.API, event openzwave.Event) {
		switch event.(type) {
		case *openzwave.NodeAvailable:
			fmt.Println("NodeAvailable!")
			node := event.GetNode()
			d := node.GetDevice()
			// Add node to any device that supports it
			if dimmer, ok := d.(*device.DimmerDevice); ok {
				deviceMap[dimmer.Node.GetNodeName()] = dimmer
			}
		case *openzwave.NodeChanged:
			fmt.Println("NodeChanged!")
		case *openzwave.NodeUnavailable:
			fmt.Println("NodeUnavailable!")
		}
	}
}
