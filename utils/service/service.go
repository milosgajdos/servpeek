// Package service implements utlity functions to interact with various service managers
package service

const (
	// Running service status
	Running Status = iota + 1
	// Stopped service status
	Stopped
	// Unknown service status
	Unknown
)

// Status defines service status
type Status int

// String implements stringer interface for Status
func (s Status) String() string {
	switch s {
	case Running:
		return "running"
	case Stopped:
		return "stopped"
	default:
		return "unknown"
	}
}
