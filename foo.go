package grule

// Foo is the object that you implement in order to make it accessible within this code.
// NOTE: structure of this might change till I arrive at a nice place with it...
type Foo interface {
	EQ(FooProp, Foo) bool // Tests equality on the FooProp passed in
	Value() string
}

// FooProp is a type that is information contained by the Foo struct
type FooProp string

// FooSTRING is an example implementation that just has the value of a string
type FooSTRING struct {
	value string
}

func (f1 *FooSTRING) EQ(f2prop FooProp, f2 Foo) bool {
	// Cast to FooSTRING
	candidate, ok := f2.(*FooSTRING)
	if !ok {
		return false
	}
	return candidate.value == f1.value
}

func (f1 *FooSTRING) Value() string {
	return f1.value
}
