package main

import (
	"os"

	"github.com/ninjasphere/go-openzwave"
	"github.com/ninjasphere/go-openzwave/LOG_LEVEL"
	"github.com/shayne/zwave/devices"
)

func main() {
	os.Exit(
		openzwave.
			BuildAPI("/etc/openzwave", "", "").
			AddIntOption("SaveLogLevel", LOG_LEVEL.NONE).
			AddIntOption("QueueLogLevel", LOG_LEVEL.NONE).
			AddIntOption("DumpTrigger", LOG_LEVEL.NONE).
			AddIntOption("PollInterval", 500).
			AddBoolOption("IntervalBetweenPolls", true).
			AddBoolOption("ValidateValueChanges", true).
			SetDeviceFactory(devices.DeviceFactory).
			SetEventsCallback(EventsCallback).
			Run())
}
