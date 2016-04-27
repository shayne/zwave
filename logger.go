package main

import "github.com/shayne/zwave/go-openzwave"

type zLogger struct {
}

func newZwaveLogger() openzwave.Logger {
	return &zLogger{}
}

// Infof func
func (*zLogger) Infof(message string, args ...interface{}) {
}

// Warningf func
func (*zLogger) Warningf(message string, args ...interface{}) {
}

// Errorf func
func (*zLogger) Errorf(message string, args ...interface{}) {
}

// Debugf func
func (*zLogger) Debugf(message string, args ...interface{}) {
}

// Tracef func
func (*zLogger) Tracef(message string, args ...interface{}) {
}
