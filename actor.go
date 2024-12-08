package main

import "sync"

/*
Allowed operations for an actor:
	* Receive a message
	* Create another actor
	* Send a message
	* Designate how to handle next message (through state)
*/

// Unclear future of this one...
type Actor interface {
	Process(msg message)
	GetState() (interface{}, error)
	DefaultState() interface{}
	Start() // Start processing messages
	Stop()  // Stop processing messages
}

type BaseActor struct {
	inbox   chan message
	state   interface{} // Something more refined
	running bool
	wg      sync.WaitGroup
}

func (a *BaseActor) Start(processor func(msg message)) {
	a.running = true
	a.wg.Add(1)

	go func() {
		defer a.wg.Done()
		for msg := range a.inbox {
			processor(msg)
		}
	}()
}

func (a *BaseActor) Stop() {
	if a.running {
		close(a.inbox)
		a.running = false
		a.wg.Wait()
	}
}

func (a *BaseActor) GetState() interface{} {
	return a.state
}

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
