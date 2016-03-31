package devices

import (
	"fmt"

	"github.com/ninjasphere/go-openzwave"
)

// DimmerDevice yah
type DimmerDevice struct {
	Name string
	Node openzwave.Node
}

// ChangeValue func
func (d *DimmerDevice) ChangeValue(value float32) {
	currentValue := d.Node.GetValue(38, 1, 0)
	currentValue.SetString(fmt.Sprintf("%v", value))
}

// NodeAdded func
func (*DimmerDevice) NodeAdded() {
}

// NodeChanged func
func (*DimmerDevice) NodeChanged() {
}

// NodeRemoved func
func (*DimmerDevice) NodeRemoved() {
}

// ValueChanged func
func (*DimmerDevice) ValueChanged(value openzwave.Value) {
}
