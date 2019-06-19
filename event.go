package grule

// Event is the interface for supplying input events to the system
type Event interface {
}

// EventProperty is the namespaced property that we are using to compare into
type EventProperty interface {
	ToString() string
}
