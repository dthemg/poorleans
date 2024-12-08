package main

import (
	"errors"
	"reflect"
)

type database struct {
	grainStates map[string]entry               // key is grain-type/id (eventually)
	messages    map[string][]serializedMessage // key is message-type/id (eventually)
}

// Message to a grain
// Contains instructions for a grain to perform a specific operation
// on a specific set of data
type serializedMessage struct {
	GrainType string
	Operation string
	Content   entry
}

type message struct {
	GrainType string
	Operation string
	Content   interface{}
}

type entry struct {
	ContentType string
	Content     []byte
}

func newEntry(value interface{}) (entry, error) {
	c, err := serialize(value)
	if err != nil {
		return entry{}, err
	}

	contentType := reflect.TypeOf(value).String()
	newEntry := entry{
		ContentType: contentType,
		Content:     c.Bytes(),
	}

	return newEntry, nil
}

func (db *database) writeGrainState(key string, value interface{}) error {
	entry, err := newEntry(value)
	if err != nil {
		return err
	}

	db.grainStates[key] = entry
	return nil
}

// Not quite the same though, storage has to be an enumerable type
func (db *database) appendMessage(key string, op string, value interface{}) error {
	entry, err := newEntry(value)
	if err != nil {
		return err
	}
	msg := serializedMessage{
		Operation: op,
		Content:   entry,
	}
	db.messages[key] = append(db.messages[key], msg)
	return nil
}

func (db *database) readGrainState(key string) (interface{}, error) {
	state, ok := db.grainStates[key]
	if !ok {
		return nil, errors.New("no such entry in database")
	}

	instance, err := deserialize(state.Content, state.ContentType)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func (db *database) popOldestMessage(key string) (message, error) {
	messages, ok := db.messages[key]
	if !ok {
		return message{}, errors.New("no message list for key in database")
	}

	if len(messages) == 0 {
		return message{}, errors.New("no messages to read")
	}

	firstMessage := messages[0]

	instance, err := deserialize(firstMessage.Content.Content, firstMessage.Content.ContentType)
	if err != nil {
		return message{}, err
	}

	db.messages[key] = messages[1:] // Remove the first message from the queue
	deserialized := message{
		Operation: firstMessage.Operation,
		GrainType: firstMessage.GrainType,
		Content:   instance,
	}
	return deserialized, nil
}

func create() (database, error) {
	db := database{
		grainStates: make(map[string]entry),
		messages:    make(map[string][]serializedMessage),
	}
	return db, nil
}
