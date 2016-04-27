package devices

import "github.com/shayne/zwave/go-openzwave"

// UnknownDevice yah
type UnknownDevice struct {
	Name string
}

// NodeAdded func
func (*UnknownDevice) NodeAdded() {
}

// NodeChanged func
func (*UnknownDevice) NodeChanged() {
}

// NodeRemoved func
func (*UnknownDevice) NodeRemoved() {
}

// ValueChanged func
func (*UnknownDevice) ValueChanged(value openzwave.Value) {
}
