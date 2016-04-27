package main

import (
	"flag"
	"os"

	"github.com/shayne/zwave/devices"
)

var deviceMap = map[string]*devices.DimmerDevice{}

func main() {
	var debug bool

	flagset := flag.NewFlagSet("zwave", flag.ContinueOnError)
	flagset.BoolVar(&debug, "debug", false, "Enable debugging")
	flagset.Parse(os.Args[1:])

	zwaveDriver, err := newZwaveDriver(debug)
	zwaveDriver.setReadyCallback(func() {
		go func() {
		}()
	})

	zwaveDriver.start()

	if err != nil {
		os.Exit(1)
	}

	// os.Exit(zwaveDriver.wait())

	shell := newShell(zwaveDriver)
	shell.cli.Println("zwave shell")
	shell.cli.Start()
}
