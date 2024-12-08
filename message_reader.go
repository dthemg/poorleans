package main

import (
	"fmt"
	"time"
)

func messageReaderLoop(db *database) {
	for i := range 100 {
		fmt.Printf("%v\n", i)
		time.Sleep(time.Second)

		// Loop through messages
		for key := range db.messages {

			// Try to pop message
			msg, err := db.popOldestMessage(key)
			if err != nil {
				fmt.Println(err.Error()) // Normal
			} else {
				fmt.Println("Read the thing")
				fmt.Println(msg.Content)
				fmt.Println(msg.GrainType)
				fmt.Println(msg.Operation)
			}

			// Wtf am I supposed to do now
			// Tell grain X do handle message Y?
			// In-memory lookup of grains with associated functions
		}
	}
}
