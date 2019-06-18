package grule

// Condition is the interface that all conditions need to implement
type Condition interface {
	Satisfies(Foo) bool    // Determines if our condition is satisfied or not
	Pack() (string, error) // Serializes the logic to a human readable format
	Unpack(string) error   // Deserializes the logic to machine readable format
}

type ConditionEQ struct {
	LHS Foo     // Left hand side is always the 'immutable' property being compared, a list, a number, etc...
	RHS BarProp // RHS is a property of the Bar that is looking to be compared
}

// Satisfies compares the LHS to the RHS
func (c *ConditionEQ) Satisfies(foo Foo) bool {
	return c.LHS.EQ(foo)
}

// Pack our shit up into one line
func (c *ConditionEQ) Pack() (string, error) {
	var packed []string

	// Always an IF
	packed = append(packed, "IF")

	// Pack up the Foo
	packedbar, err := c.RHS
}
