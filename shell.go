package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/abiosoft/ishell"
)

type shell struct {
	cli    *ishell.Shell
	driver *zDriver
}

func newShell(driver *zDriver) *shell {
	shell := &shell{
		cli:    ishell.New(),
		driver: driver,
	}

	shell.register()

	return shell
}

func (s *shell) register() {
	s.cli.Register("exit", func(args ...string) (string, error) {
		s.driver.stop()
		s.cli.Stop()
		return "Finished", nil
	})

	s.cli.Register("dimmers", func(args ...string) (string, error) {
		out := ""
		for key, val := range deviceMap {
			out += fmt.Sprintf("%s: %#v\n", key, val)
		}
		return out, nil
	})

	s.cli.Register("set-dimmers", func(args ...string) (string, error) {
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
}
