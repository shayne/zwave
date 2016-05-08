package driver

import (
	"github.com/shayne/zwave/device"
	"github.com/shayne/zwave/go-openzwave"
	"github.com/shayne/zwave/go-openzwave/MF"
	"github.com/shayne/zwave/types"
)

// DeviceFactory type
type DeviceFactory func(driver types.Driver, node openzwave.Node) openzwave.Device
type libraryT map[openzwave.ProductId]DeviceFactory

var (
	library      libraryT = make(map[openzwave.ProductId]DeviceFactory)
	dimmerSwitch          = openzwave.ProductId{MF.GE, "0x3030"}
)

// Library interface
type Library interface {
	GetDeviceFactory(id openzwave.ProductId) DeviceFactory
}

// GetLibrary func
func GetLibrary() Library {
	if len(library) == 0 {
		library[dimmerSwitch] = device.DimmerFactory
	}
	return &library
}

type unsupportedDevice struct {
}

func (lib *libraryT) GetDeviceFactory(id openzwave.ProductId) DeviceFactory {
	factory, ok := (*lib)[id]
	if ok {
		return factory
	}
	return func(driver types.Driver, node openzwave.Node) openzwave.Device {
		return &unsupportedDevice{}
	}
}

func (*unsupportedDevice) NodeAdded() {
}

func (*unsupportedDevice) NodeChanged() {
}

func (*unsupportedDevice) ValueChanged(openzwave.Value) {
}

func (*unsupportedDevice) NodeRemoved() {
}
