package main

import (
	"fmt"

	"github.com/ninjasphere/go-openzwave"
	"github.com/ninjasphere/go-openzwave/LOG_LEVEL"
	"github.com/ninjasphere/go-openzwave/NT"
	"github.com/shayne/zwave/devices"
)

type zDriver struct {
	debug    bool
	zwaveAPI openzwave.API
	exit     chan int
	ready    func()
}

func newZwaveDriver(debug bool) (*zDriver, error) {
	driver := &zDriver{
		debug:    debug,
		zwaveAPI: nil,
		exit:     make(chan int, 0),
		ready:    nil,
	}

	return driver, nil
}

func (d *zDriver) setReadyCallback(cb func()) {
	d.ready = cb
}

func (d *zDriver) start() error {
	notificationCallback := func(api openzwave.API, notification openzwave.Notification) {
		switch notification.GetNotificationType().Code {
		case NT.ALL_NODES_QUERIED_SOME_DEAD:
			if d.ready != nil {
				d.ready()
			}
		case NT.ALL_NODES_QUERIED:
			if d.ready != nil {
				d.ready()
			}
		}
	}

	zwaveDeviceFactory := func(api openzwave.API, node openzwave.Node) openzwave.Device {
		d.zwaveAPI = api
		return devices.DeviceFactory(api, node)
	}

	configurator := openzwave.
		BuildAPI("/etc/openzwave", "./zwave-config", "").
		SetLogger(newZwaveLogger()).
		SetDeviceFactory(zwaveDeviceFactory).
		SetEventsCallback(eventsCallback).
		SetNotificationCallback(notificationCallback).
		AddIntOption("SaveLogLevel", LOG_LEVEL.NONE).
		AddIntOption("QueueLogLevel", LOG_LEVEL.NONE).
		AddIntOption("DumpTrigger", LOG_LEVEL.NONE).
		AddIntOption("PollInterval", 500).
		AddBoolOption("IntervalBetweenPolls", true).
		AddBoolOption("ValidateValueChanges", true)

	d.ready()

	if d.debug {
		callback := func(api openzwave.API, notification openzwave.Notification) {
			fmt.Printf("%v\n", notification)
			notificationCallback(api, notification)
		}
		configurator.SetNotificationCallback(callback)
	}

	go func() {
		d.exit <- configurator.Run()
	}()

	return nil
}

func (d *zDriver) stop() error {
	d.zwaveAPI.QuitSignal() <- 0
	d.zwaveAPI.Shutdown(0)
	return nil
}

func (d *zDriver) wait() int {
	return <-d.exit
}
