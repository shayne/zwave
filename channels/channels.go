package channels

import (
	"fmt"

	"github.com/shayne/zwave/types"
)

// BrightnessChannel type
type BrightnessChannel struct {
	device types.DimmerDevice
}

// SendState func
func (ch *BrightnessChannel) SendState(v float64) error {
	fmt.Printf("BrightnessChannel#SendState: %f\n", v)
	return ch.device.SetBrightness(v)
}

// NewBrightnessChannel func
func NewBrightnessChannel(d types.DimmerDevice) *BrightnessChannel {
	return &BrightnessChannel{
		device: d,
	}
}
