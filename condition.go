package grule

import (
	"fmt"
	"strings"
)

// Condition is the interface that all conditions need to implement
type Condition interface {
	Satisfied(Event) bool  // Determines if our condition is satisfied or not
	Pack() (string, error) // Pack takes this condition, and writes it out to a txt file
	Unpack(string) error   // Unpack takes the string input, and coverts it to go objects
}

// ConditionEQ implements the 'equals' condition
type ConditionEQ struct {
	fact     Fact          // The fact within the system that we are comparing to (the left hand side)
	property EventProperty // The property that we are extracting from the event
}

func (c *ConditionEQ) Satisfied(event Event) bool {
	return c.fact.Equals(c.property, event)
}

func (c *ConditionEQ) Pack() (string, error) {
	var packed []string
	packed = append(packed, "IF", c.fact.Property(), "EQUALS", c.property.ToString())
	return strings.Join(packed, " "), nil
}

func (c *ConditionEQ) Unpack(input string) error {
	unpacked := strings.Fields(input)
	if len(unpacked) != 4 {
		return fmt.Errorf("invalid number of args for ConditionEQ")
	}

	// Deserialize now

}
