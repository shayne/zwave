package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/abiosoft/ishell"
	"github.com/ninjasphere/go-openzwave"
	"github.com/ninjasphere/go-openzwave/LOG_LEVEL"
	"github.com/shayne/zwave/devices"
)

var deviceMap = map[string]*devices.DimmerDevice{}

func main() {
	go func() {
		openzwave.
			BuildAPI("/etc/openzwave", "./zwave-config", "").
			AddIntOption("SaveLogLevel", LOG_LEVEL.NONE).
			AddIntOption("QueueLogLevel", LOG_LEVEL.NONE).
			AddIntOption("DumpTrigger", LOG_LEVEL.NONE).
			AddIntOption("PollInterval", 500).
			AddBoolOption("IntervalBetweenPolls", true).
			AddBoolOption("ValidateValueChanges", true).
			SetDeviceFactory(devices.DeviceFactory).
			SetEventsCallback(EventsCallback).
			Run()
	}()

	shell := ishell.New()
	shell.Println("zwave shell")

	shell.Register("dimmers", func(args ...string) (string, error) {
		out := ""
		for key, val := range deviceMap {
			out += fmt.Sprintf("%s: %#v\n", key, val)
		}
		return out, nil
	})

	shell.Register("set-dimmers", func(args ...string) (string, error) {
		if len(args) != 1 {
			return "", errors.New("usage: set-dimmers <float-value>")
		}
		value, err := strconv.ParseFloat(args[0], 32)
		if err != nil {
			return "", errors.New("usage: set-dimmers <float-value>")
		}
		for _, dimmer := range deviceMap {
			dimmer.ChangeValue(float32(value))
		}
		return "ChangeValue called", nil
	})

	shell.Start()
}
