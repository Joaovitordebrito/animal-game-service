package service

import (
	"animal-game/data"
	"log"

	"github.com/google/uuid"
)

func CreateUser(userName string) error {

	id := uuid.New().String()

	err := data.CreateUser(id, userName)

	if err != nil {
		log.Fatal("Failed to insert user")
	}

	return nil
}
