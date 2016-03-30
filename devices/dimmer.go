package devices

// DimmerDevice yah
import (
	"fmt"

	"github.com/ninjasphere/go-openzwave"
)

// DimmerDevice yah
type DimmerDevice struct {
	Name string
	node openzwave.Node
}

// ChangeValue func
func (d *DimmerDevice) ChangeValue(value float32) {
	currentValue := d.node.GetValue(38, 1, 0)
	currentValue.SetString(fmt.Sprintf("%v", value))
}

// NodeAdded yah
func (*DimmerDevice) NodeAdded() {
	fmt.Println("Node added")
}

// NodeChanged yah
func (*DimmerDevice) NodeChanged() {
	fmt.Println("Node changed")
}

// NodeRemoved yah
func (*DimmerDevice) NodeRemoved() {
	fmt.Println("Node removed")
}

// ValueChanged yah
func (*DimmerDevice) ValueChanged(value openzwave.Value) {
	fmt.Println("Value changed")
}
