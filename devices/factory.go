package devices

import (
	"fmt"

	"github.com/ninjasphere/go-openzwave"
)

// DeviceFactory yah
func DeviceFactory(api openzwave.API, node openzwave.Node) openzwave.Device {
	fmt.Println("myDeviceFactory called")

	if node.GetId() == 2 {
		return &DimmerDevice{
			Name: "Hallway dimmer",
			node: node,
		}
	}

	return &UnknownDevice{}
}
