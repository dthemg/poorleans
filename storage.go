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
	storage map[string]actorState
}

var typeRegistry = map[string]reflect.Type{}

// Keep a registry of existing types
func registerStateType(state interface{}) error {
	stateType := reflect.TypeOf(state)
	typeName := stateType.String()

	_, ok := typeRegistry[typeName]
	if ok {
		return errors.New("multiple serializer registrations")
	}

	fmt.Printf("\tRegistered serializer for %s\n", typeName)
	typeRegistry[typeName] = stateType
	gob.Register(state)
	return nil
}

func (db *dataBase) updateState(key string, value interface{}) error {
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

	db.storage[key] = newState
	return nil
}

func (db *dataBase) getState(key string) (interface{}, error) {
	state, ok := db.storage[key]
	if !ok {
		return nil, errors.New("no such entry in database")
	}

	stateType, ok := typeRegistry[state.StateType]
	if !ok {
		log.Fatalf("Missing type for %s", state.StateType)
	}
	instance := reflect.New(stateType).Interface()

	buf := bytes.NewBuffer(state.State)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(instance); err != nil {
		log.Fatal("Error when decoding state")
	}

	return instance, nil
}

func (db *dataBase) print() {
	fmt.Println("DB contents:")
	for key, state := range db.storage {
		stateType, ok := typeRegistry[state.StateType]
		if !ok {
			log.Fatalf("Missing type for %s", state.StateType)
		}
		instance := reflect.New(stateType).Interface()

		buf := bytes.NewBuffer(state.State)
		dec := gob.NewDecoder(buf)
		if err := dec.Decode(instance); err != nil {
			log.Fatal("Error when decoding")
		}

		fmt.Printf("\t%s: %s\n", key, stateType)
	}
}

func createDatabase() (dataBase, error) {
	db := dataBase{
		storage: make(map[string]actorState),
	}
	return db, nil
}
