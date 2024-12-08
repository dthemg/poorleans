package main

import (
	"fmt"
	"log"
	"time"

	"github.com/davecgh/go-spew/spew"
)

type TestActorState2 struct {
	C string
	D int
}

type MessageType1 struct {
	MyMessage string
}

func registerContentTypes() error {
	fmt.Println("Registering serializers")
	if err := registerContentType(TestActorState{}); err != nil {
		return err
	}
	if err := registerContentType(MessageType1{}); err != nil {
		return err
	}

	return registerContentType(TestActorState2{})
}

func main() {
	fmt.Println("Starting up Poorleans")

	db, err := create()
	if err != nil {
		log.Fatal(err)
	}

	err = registerContentTypes()
	if err != nil {
		log.Fatal(err)
	}

	// Run as separate process (?)
	go messageReaderLoop(&db)

	testActorState := TestActorState{
		A: 100,
		B: "David",
	}
	testKey := "myKey"
	// Move into BaseActor
	err = db.writeGrainState(testKey, testActorState)
	if err != nil {
		log.Fatal(err)
	}

	db.print()

	readTestState, err := db.readGrainState(testKey)
	if err != nil {
		log.Fatal(err)
	}
	s1, ok := readTestState.(*TestActorState)
	if !ok {
		log.Fatal("could not read state")
	}

	// Write messages

	messageKey := "my-message-key-1"
	err = db.appendMessage(
		messageKey,
		"my-operation",
		MessageType1{MyMessage: "Top secret message"})
	if err != nil {
		log.Fatal("could not write message")
	}
	msg, err := db.popOldestMessage(messageKey)
	if err != nil {
		log.Fatalf("could not read message: %s", err.Error())
	}

	fmt.Println("Grain states:")
	spew.Dump(s1)

	fmt.Println("Messages:")
	spew.Dump(msg)

	// New beginnings
	intActor := NewIntActor(38)
	intActor.Process(message{})
	intState1 := intActor.GetState()
	spew.Dump(intState1)
	intActor.Process(message{})
	intState2 := intActor.GetState()
	spew.Dump(intState2)

	for range 10 {
		fmt.Println("preventing program exit")
		time.Sleep(time.Second * 10)
	}
}
