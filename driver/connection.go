package driver

import (
	"fmt"

	"github.com/shayne/zwave/types"
)

type connection struct {
}

func (c *connection) ExportChannel(d types.Device, ch types.Channel, id string) error {
	fmt.Printf("ExportChannel: %v %v %v\n", d, ch, id)
	return nil
}
