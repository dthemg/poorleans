package main

/*
Allowed operations for an actor:
	* Receive a message
	* Create another actor
	* Send a message
	* Designate how to handle next message (through state)
*/

// Also need to serialize which methods the actors have

type Actor interface {
	Process()
}

type CounterActor struct {
}

type CounterState struct {
	counter int
}
