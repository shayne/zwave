package device

import (
	"github.com/shayne/zwave/go-openzwave"
	"github.com/shayne/zwave/types"
)

// Device struct
type Device struct {
	Driver types.Driver
	// Info      *model.Device
	// SendEvent func(event string, payload interface{}) error
	Node openzwave.Node
}

// func (device *Device) GetDriver() ninja.Driver {
// 	return device.Driver.Ninja()
// }

// func (device *Device) GetDeviceInfo() *model.Device {
// 	return device.Info
// }

// func (device *Device) SetEventHandler(sendEvent func(event string, payload interface{}) error) {
// 	device.SendEvent = sendEvent
// }

// Init func
func (device *Device) Init(driver types.Driver, node openzwave.Node) {
	device.Driver = driver
	device.Node = node
	// device.Info = &model.Device{}

	productID := node.GetProductId()
	productDescription := node.GetProductDescription()

	sigs := make(map[string]string)

	sigs["zwave:manufacturerId"] = productID.ManufacturerId
	sigs["zwave:productId"] = productID.ProductId
	sigs["zwave:manufacturerName"] = productDescription.ManufacturerName
	sigs["zwave:productName"] = productDescription.ProductName
	sigs["zwave:productType"] = productDescription.ProductType

	// device.Info.Signatures = &sigs

	//
	// This naming scheme won't survive reconfigurations of the network
	// where the network has two devices of the same type.
	//
	// So, we will need to investigate generating a unique token that
	// gets stored in the device. When this scheme is implemented we
	// should update the naming scheme used here from v0 to v1.
	//
	// device.Info.NaturalIDType = "ninja.zwave.v0"
	// device.Info.NaturalID = fmt.Sprintf(
	// 	"%08x:%03d:%s:%s",
	// 	node.GetHomeId(),
	// 	node.GetId(),
	// 	productId.ManufacturerId,
	// 	productId.ProductId)
	//
	// device.Info.Name = &productDescription.ProductName

	// initialize brightness from the current level

}
