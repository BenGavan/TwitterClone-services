package main

type database struct {

}


func newDatabase() database {
	return database{}
}

func (db *database) savePasswordHash(uuid string, passwordHash string) error {
	return nil
}