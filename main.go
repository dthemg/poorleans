package main

import (
	"fmt"
	"log"

	"github.com/davecgh/go-spew/spew"
)

type TestActorState1 struct {
	A int
	B string
}

type TestActorState2 struct {
	C string
	D int
}

func registerContentTypes() error {
	fmt.Println("Registering serializers")
	if err := registerContentType(TestActorState1{}); err != nil {
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

	state1 := TestActorState1{
		A: 100,
		B: "David",
	}
	err = db.write("myKey", state1)
	if err != nil {
		log.Fatal(err)
	}

	state2 := TestActorState2{
		C: "hello",
		D: 32,
	}
	err = db.write("myKey2", state2)
	if err != nil {
		log.Fatal(err)
	}

	db.print()

	readState1, err := db.read("myKey")
	if err != nil {
		log.Fatal(err)
	}
	s1, ok := readState1.(*TestActorState1)
	if !ok {
		log.Fatal("could not read state")
	}

	readState2, err := db.read("myKey2")
	if err != nil {
		log.Fatal(err)
	}
	s2, ok := readState2.(*TestActorState2)
	if !ok {
		log.Fatal("could not read state")
	}

	fmt.Println("Reading state back:")
	spew.Dump(s1)
	spew.Dump(s2)
}
