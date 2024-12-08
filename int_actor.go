package main

// Add an actual test implementation of this
type IntActor struct {
	BaseActor
}

// Initialization
func NewIntActor(initialState int) *IntActor {
	actor := &IntActor{
		BaseActor: BaseActor{
			inbox: make(chan message, 10),
			state: initialState,
		},
	}
	return actor
}

// Process
func (a *IntActor) Process(msg message) {
	a.state = a.state.(int) + 1
}
