package main

import (
	"fmt"
	"log"
)

type TestActorState struct {
	A int
	B string
}

func main() {
	fmt.Println("Starting up Poorleans")

	// Register all state types
	registerStateType(TestActorState{})

	db, err := createDatabase()
	if err != nil {
		log.Fatal(err)
	}

	state := TestActorState{
		A: 100,
		B: "David",
	}
	err = db.addEntry("myKey", state)
	if err != nil {
		log.Fatal(err)
	}

	db.print()
}
