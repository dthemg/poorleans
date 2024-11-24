package main

/*
Allowed operations for an actor:
	* Receive a message
	* Create another actor
	* Send a message
	* Designate how to handle next message (through state)
*/

type Actor interface {
	Process()
	GetState() (interface{}, error)
	DefaultState() interface{}
}

type ActorImpl struct {
	Actor
}

func (a ActorImpl) ProcessActor() {
	a.Actor.Process()
}
