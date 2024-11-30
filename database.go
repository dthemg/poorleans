package main

import (
	"errors"
	"reflect"
)

type database struct {
	grainStates map[string]entry   // key is grain-type/id (eventually)
	messages    map[string][]entry // key is message-type/id (eventually)
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
func (db *database) appendMessage(key string, value interface{}) error {
	entry, err := newEntry(value)
	if err != nil {
		return err
	}

	db.messages[key] = append(db.messages[key], entry)
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

func (db *database) popOldestMessage(key string) (interface{}, error) {
	messages, ok := db.messages[key]
	if !ok {
		return nil, errors.New("no message list for key in database")
	}

	if len(messages) == 0 {
		return nil, errors.New("no messages to read")
	}

	message := messages[0]

	instance, err := deserialize(message.Content, message.ContentType)
	if err != nil {
		return nil, err
	}

	db.messages[key] = messages[1:] // Remove the message from the queue

	return instance, nil
}

func create() (database, error) {
	db := database{
		grainStates: make(map[string]entry),
		messages:    make(map[string][]entry),
	}
	return db, nil
}
