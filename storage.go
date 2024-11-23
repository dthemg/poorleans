package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"log"
	"reflect"
)

type actorState struct {
	Key       string
	StateType string
	State     []byte
}

type dataBase struct {
	states []actorState
}

var typeRegistry = map[string]reflect.Type{}

// Keep a registry of existing types
func registerStateType(state interface{}) {
	stateType := reflect.TypeOf(state)
	typeName := stateType.String()

	// Only register if it does not already exist
	_, ok := typeRegistry[typeName]
	if ok {
		return
	}

	fmt.Printf("Adding serializer for %s\n", typeName)
	typeRegistry[typeName] = stateType
	gob.Register(state)
}

func (db *dataBase) addEntry(key string, value interface{}) error {
	for i := 0; i < len(db.states); i++ {
		if db.states[i].Key == key {
			return errors.New("entry already exists")
		}
	}

	// Register state type so it later can be deserialized

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(value); err != nil {
		return err
	}

	stateType := reflect.TypeOf(value).String()

	newState := actorState{
		Key:       key,
		StateType: stateType,
		State:     buf.Bytes(),
	}

	db.states = append(db.states, newState)
	return nil
}

func (db *dataBase) print() {
	fmt.Println("DB contents:")
	for i := 0; i < len(db.states); i++ {
		var e = db.states[i]

		stateType, ok := typeRegistry[e.StateType]
		if !ok {
			log.Fatal("Missing type")
		}
		instance := reflect.New(stateType).Interface()

		buf := bytes.NewBuffer(e.State)
		dec := gob.NewDecoder(buf)
		if err := dec.Decode(instance); err != nil {
			log.Fatal("Error when decoding")
		}

		fmt.Printf("%s: %s", e.Key, e.StateType)
	}
}

func createDatabase() (dataBase, error) {
	db := dataBase{}
	return db, nil
}
