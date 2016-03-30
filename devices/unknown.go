package devices

import (
	"fmt"

	"github.com/ninjasphere/go-openzwave"
)

// UnknownDevice yah
type UnknownDevice struct {
	Name string
	node openzwave.Node
}

// NodeAdded yah
func (*UnknownDevice) NodeAdded() {
	fmt.Println("Node added")
}

// NodeChanged yah
func (*UnknownDevice) NodeChanged() {
	fmt.Println("Node changed")
}

// NodeRemoved yah
func (*UnknownDevice) NodeRemoved() {
	fmt.Println("Node removed")
}

// ValueChanged yah
func (*UnknownDevice) ValueChanged(value openzwave.Value) {
	fmt.Println("Value changed")
}
