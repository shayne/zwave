package devices

import (
	"fmt"

	openzwave "github.com/ninjasphere/go-openzwave"
)

// DeviceFactory yah
func DeviceFactory(api openzwave.API, node openzwave.Node) openzwave.Device {
	fmt.Printf("Device factory, node: %s %s\n", node.GetProductDescription(), node.GetNodeName())

	desc := node.GetProductDescription()
	// Only detecting dimmer types for now
	if desc.ProductType == "0x4450" {
		return &DimmerDevice{
			Name: node.GetNodeName(),
			Node: node,
		}
	}

	fmt.Printf("Unhandled node: %s\n", node)
	return nil
	// return &UnknownDevice{
	// 	Name: "Unknown Device",
	// }
}
