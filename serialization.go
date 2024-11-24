package main

import (
	"bytes"
	"encoding/gob"
	"errors"
	"fmt"
	"reflect"
)

var typeRegistry = map[string]reflect.Type{}

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

func serialize(value interface{}) (*bytes.Buffer, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(value)
	if err != nil {
		return nil, err
	}
	return &buf, nil
}

func deserialize(content []byte, contentType string) (interface{}, error) {
	stateType, ok := typeRegistry[contentType]
	if !ok {
		return nil, fmt.Errorf("missing type for %s", stateType)
	}
	instance := reflect.New(stateType).Interface()

	buf := bytes.NewBuffer(content)
	dec := gob.NewDecoder(buf)
	if err := dec.Decode(instance); err != nil {
		return nil, fmt.Errorf("error when decoding type %s", stateType)
	}

	return instance, nil
}
