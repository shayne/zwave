package types

// Device type
type Device interface {
}

// Driver type
type Driver interface {
	Start() error
	Stop() error

	Connection() Connection
}

// Connection type
type Connection interface {
	ExportChannel(Device, Channel, string) error
}

// Channel type
type Channel interface {
	SendState(float64) error
}

// DimmerDevice type
type DimmerDevice interface {
	SetBrightness(float64) error
}
