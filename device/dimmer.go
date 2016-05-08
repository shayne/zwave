package device

import (
	"fmt"
	"time"

	"github.com/shayne/zwave/channels"
	"github.com/shayne/zwave/go-openzwave"
	"github.com/shayne/zwave/go-openzwave/CC"
	"github.com/shayne/zwave/types"
	"github.com/shayne/zwave/utils"
)

const (
	maxDeviceBrightness = 100             // by experiment, a level of 100 does not work for this device
	maxDelay            = time.Second * 5 // maximum delay for apply calls
)

var (
	multiLevelSwitch = openzwave.ValueID{CC.SWITCH_MULTILEVEL, 1, 0}
)

// DimmerDevice yah
type DimmerDevice struct {
	Device

	brightnessChannel *channels.BrightnessChannel

	// brightness is a cache of the current brightness when the device is switched off.
	// It is updated from the device on a confirmed attempt to adjust the level to a non-zero value
	brightness uint8

	refresh chan struct{} // used to wait for confirmation of updates after a level change

	emitter utils.Emitter
}

// DimmerFactory func
func DimmerFactory(driver types.Driver, node openzwave.Node) openzwave.Device {
	device := &DimmerDevice{}

	device.Init(driver, node)

	var ok bool

	device.brightness, ok = device.Node.GetValueWithId(multiLevelSwitch).GetUint8()
	if !ok || device.brightness == 0 {
		// we have to reset brightness to 100 since we apply brightness when
		// we switch it on

		//
		// one implication of this is that if the controller is removed
		// then replaced, while the light is off, the original brightness
		// will be lost
		//

		device.brightness = maxDeviceBrightness
	}

	device.refresh = make(chan struct{}, 0)

	device.emitter = utils.Filter(
		func(level utils.Equatable) {
			// device.unconditionalSendLightState(level.(*utils.WrappedUint8).Unwrap())
		},
		30*time.Second)

	return device
}

// NodeAdded func
func (d *DimmerDevice) NodeAdded() {
	node := d.Node
	conn := d.Driver.Connection()
	d.brightnessChannel = channels.NewBrightnessChannel(d)

	err := conn.ExportChannel(d, d.brightnessChannel, "brightness")
	if err != nil {
		fmt.Printf("failed to export brightness channel for %v: %s", node, err)
	}

	d.Node.GetValueWithId(multiLevelSwitch).SetPollingState(true)
}

// NodeChanged func
func (d *DimmerDevice) NodeChanged() {
}

// NodeRemoved func
func (*DimmerDevice) NodeRemoved() {
}

// ValueChanged func
func (d *DimmerDevice) ValueChanged(v openzwave.Value) {
	switch v.Id() {
	case multiLevelSwitch:
		select {
		case d.refresh <- struct{}{}:
		default:
			d.sendLightState()
		}
	}
}

// SetBrightness func
func (d *DimmerDevice) SetBrightness(state float64) error {
	var err error
	if state < 0 {
		state = 0
	} else if state > 1.0 {
		state = 1.0
	}
	level, ok := d.Node.GetValueWithId(multiLevelSwitch).GetUint8()
	fmt.Printf("Current level: %d\n", level)
	if ok {
		newLevel := uint8(state * maxDeviceBrightness)
		// if level > 0 {
		err = d.setDeviceLevel(newLevel)
		// } else {
		// 	d.brightness = newLevel // to be applied when device is switched on
		// 	d.emitter.Reset()
		// }
	} else {
		err = fmt.Errorf("Unable to apply brightness - get failed.")
	}
	return err
}

//
// Issue a set against the OpenZWave API, then wait until the refreshed
// value matches the requested level or until a timeout, issuing refreshes
// as required.
//
func (d *DimmerDevice) setDeviceLevel(level uint8) error {

	val := d.Node.GetValueWithId(multiLevelSwitch)

	if level >= maxDeviceBrightness {
		// aeon will reject attempts to set the level to exactly 100
		level = maxDeviceBrightness - 1
	}

	if !val.SetUint8(level) {
		return fmt.Errorf("Failed to set level to %d - set failed", level)
	}
	timer := time.NewTimer(maxDelay)

	// loop until timeout or until refresh yields expected level

	for {
		if !val.Refresh() {
			return fmt.Errorf("Failed to set required level to %d - refresh failed", level)
		}
		select {
		case timeout := <-timer.C:
			_ = timeout
			if level != 0 {
				d.brightness = level
			}
			return fmt.Errorf("Failed to set required level to %d - timeout", level)
		case refreshed := <-d.refresh:
			_ = refreshed
			current, ok := val.GetUint8()
			if ok && current == level {
				if level != 0 {
					d.brightness = level
				}
				d.emitter.Reset()
				return nil
			}
		}
	}
}

//
// This call is used to reflect notifications about the current
// state of the light back to towards the ninja network
//
func (d *DimmerDevice) sendLightState() {
	level, ok := d.Node.GetValueWithId(multiLevelSwitch).GetUint8()
	if ok {
		//
		// Emit the current state, but filter out levels that don't change
		// within a specified period.
		//
		d.emitter.Emit(utils.WrapUint8(level))
	}
}

func (d *DimmerDevice) unconditionalSendLightState(level uint8) {
	if d.brightness == maxDeviceBrightness-1 {
		d.brightness = 100
	}

	// onOff := level != 0
	brightness := float64(d.brightness) / maxDeviceBrightness

	// d.onOffChannel.SendState(onOff)
	d.brightnessChannel.SendState(brightness)
}
