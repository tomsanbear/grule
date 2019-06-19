package grule

// Fact represents some immutable property(s) within the system
type Fact interface {
	Equals(EventProperty, Event) bool // Takes an event, and an event property and determines equivalence
	Property() string                 // Returns the property this fact represents
}
