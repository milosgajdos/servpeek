package service

const (
	Running Status = iota + 1
	Stopped
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
