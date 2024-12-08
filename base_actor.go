package main

import "sync"

/*
Allowed operations for an actor:
	* Receive a message
	* Create another actor
	* Send a message
	* Designate how to handle next message (through state)
*/

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
