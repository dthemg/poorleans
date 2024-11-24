package main

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"log"
	"reflect"
)

func (db *dataBase) print() {
	fmt.Println("DB contents:")
	for key, state := range db.storage {
		stateType, ok := typeRegistry[state.ContentType]
		if !ok {
			log.Fatalf("Missing type for %s", state.ContentType)
		}
		instance := reflect.New(stateType).Interface()

		buf := bytes.NewBuffer(state.Content)
		dec := gob.NewDecoder(buf)
		if err := dec.Decode(instance); err != nil {
			log.Fatal("Error when decoding")
		}

		fmt.Printf("\t%s: %s\n", key, stateType)
	}
}
