package main

import (
	"flag"
	"os"

	"github.com/shayne/zwave/devices"
)

var deviceMap = map[interface{}]interface{}{}

func main() {
	var debug bool

	flagset := flag.NewFlagSet("zwave", flag.ContinueOnError)
	flagset.BoolVar(&debug, "debug", false, "Enable debugging")
	flagset.Parse(os.Args[1:])

	zwaveDriver, err := devices.NewZwaveDriver(&devices.ZDriverCfg{
		DeviceMap: deviceMap,
		Debug:     debug,
	})

	zwaveDriver.Start()

	if err != nil {
		os.Exit(1)
	}

	// os.Exit(zwaveDriver.wait())

	shell := newShell(zwaveDriver)
	shell.cli.Println("zwave shell")
	shell.cli.Start()
}
