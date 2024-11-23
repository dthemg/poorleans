package main

import (
	"fmt"
	"log"
)

type TestActorState1 struct {
	A int
	B string
}

type TestActorState2 struct {
	C string
	D int
}

func registerAllStateTypes() error {
	fmt.Println("Registering serializers")
	if err := registerStateType(TestActorState1{}); err != nil {
		return err
	}
	return registerStateType(TestActorState2{})
}

func main() {
	fmt.Println("Starting up Poorleans")

	db, err := createDatabase()
	if err != nil {
		log.Fatal(err)
	}

	err = registerAllStateTypes()
	if err != nil {
		log.Fatal(err)
	}

	state1 := TestActorState1{
		A: 100,
		B: "David",
	}
	err = db.updateState("myKey", state1)
	if err != nil {
		log.Fatal(err)
	}

	state2 := TestActorState2{
		C: "hello",
		D: 32,
	}
	err = db.updateState("myKey2", state2)
	if err != nil {
		log.Fatal(err)
	}

	db.print()

	readState1, err := db.getState("myKey")
	if err != nil {
		log.Fatal(err)
	}
	s1, ok := readState1.(*TestActorState1)
	if !ok {
		log.Fatal("could not read state")
	}

	readState2, err := db.getState("myKey2")
	if err != nil {
		log.Fatal(err)
	}
	s2, ok := readState2.(*TestActorState2)
	if !ok {
		log.Fatal("could not read state")
	}

	fmt.Println("Reading state back:")
	fmt.Printf("\nTestActorState1: { A: %v, B: %s }\n", s1.A, s1.B)
	fmt.Printf("\nTestActorState2: { C: %s, D: %v }\n", s2.C, s2.D)
}
