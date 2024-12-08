package main

type TestActor struct {
	BaseActor
}

type TestActorState struct {
	A int
	B string
}

func CreateTestActor(st TestActorState) *TestActor {
	actor := &TestActor{
		BaseActor: BaseActor{
			inbox: make(chan message, 10),
			state: st,
		},
	}
	return actor
}

// Process
func (a *TestActor) Process(msg message) {
	asState := a.state.(TestActorState)
	asState.A = 42
	asState.B = "Q"
	// Persist somehow
	// maybe by implementing the following?
	// readStateAsync method
	// writeStateAsync method
}
