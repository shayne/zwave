package utils

// support for uint8

// Equatable interface
type Equatable interface {
	Equals(w Equatable) bool
}

// WrappedUint8 interface
type WrappedUint8 struct {
	val uint8
}

// WrapUint8 func
func WrapUint8(v uint8) *WrappedUint8 {
	return &WrappedUint8{v}
}

// Equals func
func (w *WrappedUint8) Equals(other Equatable) bool {
	switch other.(type) {
	case *WrappedUint8:
		return w.val == other.(*WrappedUint8).val
	default:
		return false
	}
}

// Unwrap func
func (w *WrappedUint8) Unwrap() uint8 {
	return w.val
}

// support for bool

// WrappedBool struct
type WrappedBool struct {
	val bool
}

// WrapBool func
func WrapBool(v bool) *WrappedBool {
	return &WrappedBool{v}
}

// Equals func
func (w *WrappedBool) Equals(other Equatable) bool {
	switch other.(type) {
	case *WrappedBool:
		return w.val == other.(*WrappedBool).val
	default:
		return false
	}
}

// Unwrap func
func (w *WrappedBool) Unwrap() bool {
	return w.val
}
