package main

import (
	"errors"
	"reflect"
)

type content struct {
	Key         string
	ContentType string
	Content     []byte
}

type dataBase struct {
	storage map[string]content
}

func (db *dataBase) write(key string, value interface{}) error {
	c, err := serialize(value)
	if err != nil {
		return err
	}

	contentType := reflect.TypeOf(value).String()

	newContent := content{
		Key:         key,
		ContentType: contentType,
		Content:     c.Bytes(),
	}

	db.storage[key] = newContent
	return nil
}

func (db *dataBase) read(key string) (interface{}, error) {
	state, ok := db.storage[key]
	if !ok {
		return nil, errors.New("no such entry in database")
	}

	instance, err := deserialize(state.Content, state.ContentType)
	if err != nil {
		return nil, err
	}

	return instance, nil
}

func create() (dataBase, error) {
	db := dataBase{
		storage: make(map[string]content),
	}
	return db, nil
}
