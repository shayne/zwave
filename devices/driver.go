package devices

import (
	"fmt"

	"github.com/shayne/zwave/go-openzwave"
	"github.com/shayne/zwave/go-openzwave/LOG_LEVEL"
	"github.com/shayne/zwave/go-openzwave/NT"

	"github.com/shayne/zwave/logger"
)

// ZDriver type
type ZDriver struct {
	config   *ZDriverCfg
	zwaveAPI openzwave.API
	exit     chan int
}

// ZDriverCfg type
type ZDriverCfg struct {
	DeviceMap map[interface{}]interface{}
	Debug     bool
}

// NewZwaveDriver func
func NewZwaveDriver(config *ZDriverCfg) (*ZDriver, error) {
	driver := &ZDriver{
		config:   config,
		zwaveAPI: nil,
		exit:     make(chan int, 0),
	}

	return driver, nil
}

// Start func
func (d *ZDriver) Start() error {
	notificationCallback := func(api openzwave.API, notification openzwave.Notification) {
		switch notification.GetNotificationType().Code {
		case NT.ALL_NODES_QUERIED_SOME_DEAD:
		case NT.ALL_NODES_QUERIED:
		}
	}

	zwaveDeviceFactory := func(api openzwave.API, node openzwave.Node) openzwave.Device {
		d.zwaveAPI = api
		return DeviceFactory(api, node)
	}

	configurator := openzwave.
		BuildAPI("/etc/openzwave", "./zwave-config", "").
		SetLogger(logger.NewZwaveLogger()).
		SetDeviceFactory(zwaveDeviceFactory).
		SetEventsCallback(eventsCallback(d.config.DeviceMap)).
		SetNotificationCallback(notificationCallback).
		AddIntOption("SaveLogLevel", LOG_LEVEL.NONE).
		AddIntOption("QueueLogLevel", LOG_LEVEL.NONE).
		AddIntOption("DumpTrigger", LOG_LEVEL.NONE).
		AddIntOption("PollInterval", 500).
		AddBoolOption("IntervalBetweenPolls", true).
		AddBoolOption("ValidateValueChanges", true)

	if d.config.Debug {
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

// Stop func
func (d *ZDriver) Stop() error {
	d.zwaveAPI.QuitSignal() <- 0
	d.zwaveAPI.Shutdown(0)
	return nil
}

func (d *ZDriver) wait() int {
	return <-d.exit
}
