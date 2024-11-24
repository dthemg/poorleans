package main

type ActorRegistry struct {
	// Backed by in-memory-db
}

// Probably add some prefix to the keys or something

func (db *dataBase) RegisterActor(key string, actor Actor) error {
	state := actor.DefaultState()
	err := db.write(key, state)
	if err != nil {
		return err
	}
	return nil
}

func (db *dataBase) GetActorState(key string) (interface{}, error) {
	state, err := db.read(key)
	if err != nil {
		return nil, err
	}
	return state, nil
}

func (db *dataBase) WriteActorState(key string, newState interface{}) {
	db.write(key, newState)
}

// Some infinite loop somewhere picking up on which messages are ready to be read? Using channels maybe?
// Needs some notion of which entries are messages? Probably separate from state table

// TODO
func (db *dataBase) SendMessage(key string, msg string) {
}
